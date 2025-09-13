package cache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCircuitBreaker_ClosedState(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 3,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Circuit breaker should start in closed state
	assert.Equal(t, StateClosed, cb.State())

	// Successful calls should keep circuit closed
	for i := 0; i < 5; i++ {
		result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return "success", nil
		})
		
		assert.NoError(t, err)
		assert.Equal(t, "success", result)
		assert.Equal(t, StateClosed, cb.State())
	}
}

func TestCircuitBreaker_OpenState(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 3,
		Interval:    time.Second,
		Timeout:     100 * time.Millisecond,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Generate failures to open the circuit
	for i := 0; i < 3; i++ {
		result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("test failure")
		})
		
		assert.Error(t, err)
		assert.Nil(t, result)
	}

	// Circuit should now be open
	assert.Equal(t, StateOpen, cb.State())

	// Calls should fail immediately without executing the function
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		t.Fatal("Function should not be called when circuit is open")
		return nil, nil
	})
	
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "circuit breaker is open")
}

func TestCircuitBreaker_HalfOpenState(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 2,
		Interval:    time.Second,
		Timeout:     50 * time.Millisecond,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Open the circuit
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failure")
		})
	}
	
	assert.Equal(t, StateOpen, cb.State())

	// Wait for timeout to transition to half-open
	time.Sleep(60 * time.Millisecond)

	// Next call should transition to half-open
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		return "success", nil
	})
	
	assert.NoError(t, err)
	assert.Equal(t, "success", result)
	assert.Equal(t, StateClosed, cb.State()) // Should transition back to closed on success
}

func TestCircuitBreaker_HalfOpenToOpen(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 2,
		Interval:    time.Second,
		Timeout:     50 * time.Millisecond,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Open the circuit
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failure")
		})
	}
	
	// Wait for timeout
	time.Sleep(60 * time.Millisecond)

	// Fail again in half-open state
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		return nil, errors.New("still failing")
	})
	
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, StateOpen, cb.State()) // Should transition back to open
}

func TestCircuitBreaker_Metrics(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 5,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Execute some successful calls
	for i := 0; i < 3; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return "success", nil
		})
	}
	
	// Execute some failed calls
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failure")
		})
	}
	
	metrics := cb.Metrics()
	assert.Equal(t, uint64(5), metrics.TotalRequests)
	assert.Equal(t, uint64(3), metrics.TotalSuccesses)
	assert.Equal(t, uint64(2), metrics.TotalFailures)
	assert.Equal(t, StateClosed, metrics.State)
}

func TestCircuitBreaker_ConcurrentExecution(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 10,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	cb := NewCircuitBreaker("test", config)
	
	numGoroutines := 50
	results := make(chan error, numGoroutines)
	
	// Execute concurrent calls
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			_, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
				time.Sleep(10 * time.Millisecond) // Simulate some work
				if i%10 == 0 {
					return nil, errors.New("simulated failure")
				}
				return "success", nil
			})
			results <- err
		}(i)
	}
	
	// Collect results
	successCount := 0
	failureCount := 0
	for i := 0; i < numGoroutines; i++ {
		err := <-results
		if err != nil {
			failureCount++
		} else {
			successCount++
		}
	}
	
	// Verify we got the expected mix of successes and failures
	assert.Equal(t, 5, failureCount) // Every 10th call fails
	assert.Equal(t, 45, successCount)
	assert.Equal(t, StateClosed, cb.State()) // Should still be closed
}

func TestCircuitBreaker_ContextCancellation(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 5,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	cb := NewCircuitBreaker("test", config)
	
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately
	
	result, err := cb.Execute(ctx, func(ctx context.Context) (interface{}, error) {
		return "should not execute", nil
	})
	
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, context.Canceled, err)
}

func TestCircuitBreaker_Reset(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 2,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	cb := NewCircuitBreaker("test", config)
	
	// Open the circuit
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failure")
		})
	}
	
	assert.Equal(t, StateOpen, cb.State())
	
	// Reset the circuit breaker
	cb.Reset()
	
	// Should be back to closed state
	assert.Equal(t, StateClosed, cb.State())
	
	// Should accept calls normally
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		return "success after reset", nil
	})
	
	assert.NoError(t, err)
	assert.Equal(t, "success after reset", result)
}

func TestCircuitBreaker_Name(t *testing.T) {
	config := CircuitBreakerConfig{
		MaxRequests: 5,
		Interval:    time.Second,
		Timeout:     time.Second,
	}

	name := "test-circuit-breaker"
	cb := NewCircuitBreaker(name, config)
	
	assert.Equal(t, name, cb.Name())
}

func TestCircuitBreakerConfig_Validation(t *testing.T) {
	// Test valid config
	validConfig := CircuitBreakerConfig{
		MaxRequests: 5,
		Interval:    time.Second,
		Timeout:     time.Second,
	}
	
	cb := NewCircuitBreaker("test", validConfig)
	assert.NotNil(t, cb)
	
	// Test with zero MaxRequests (should use default)
	zeroMaxConfig := CircuitBreakerConfig{
		MaxRequests: 0,
		Interval:    time.Second,
		Timeout:     time.Second,
	}
	
	cb = NewCircuitBreaker("test", zeroMaxConfig)
	assert.NotNil(t, cb)
	
	// The circuit breaker should still function with default values
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		return "test", nil
	})
	
	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}