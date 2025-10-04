package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vertikon/mcp-ultra/pkg/logger"
)

func createTestDistributedCache(t *testing.T) (*DistributedCache, *miniredis.Miniredis) {
	s, err := miniredis.Run()
	require.NoError(t, err)

	config := CacheConfig{
		Addrs:                []string{s.Addr()},
		Password:             "",
		DB:                   0,
		PoolSize:             10,
		MinIdleConns:         5,
		MaxConnAge:           30 * time.Minute,
		PoolTimeout:          5 * time.Second,
		IdleTimeout:          10 * time.Minute,
		IdleCheckFrequency:   time.Minute,
		DefaultTTL:           5 * time.Minute,
		MaxKeySize:           1024,
		MaxValueSize:         1024 * 1024,
		Strategy:             StrategyWriteThrough,
		EvictionPolicy:       EvictionLRU,
		EnableMetrics:        true,
		EnableTracing:        false,
		PrefixNamespace:      "test",
		CompressionEnabled:   false,
		CompressionLevel:     6,
		Partitions:           4,
		ReplicationFactor:    1,
		ConsistentHashing:    true,
		EnableCircuitBreaker: false,
		CircuitBreakerConfig: CircuitBreakerConfig{
			MaxRequests: 100,
			Interval:    time.Minute,
			Timeout:     30 * time.Second,
		},
	}

	logger := logger.NewZapAdapter(nil)
	cache := NewDistributedCache(config, logger)

	return cache, s
}

func TestDistributedCache_SetAndGet(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "test_key"
	value := "test_value"

	// Test Set
	err := cache.Set(ctx, key, value, time.Minute)
	assert.NoError(t, err)

	// Test Get
	var result string
	err = cache.Get(ctx, key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestDistributedCache_SetWithTTL(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "ttl_key"
	value := "ttl_value"
	ttl := 100 * time.Millisecond

	// Set value with short TTL
	err := cache.Set(ctx, key, value, ttl)
	assert.NoError(t, err)

	// Verify value exists initially
	var result string
	err = cache.Get(ctx, key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)

	// Wait for TTL to expire
	time.Sleep(150 * time.Millisecond)

	// Verify value no longer exists
	err = cache.Get(ctx, key, &result)
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err)
}

func TestDistributedCache_Delete(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "delete_key"
	value := "delete_value"

	// Set value
	err := cache.Set(ctx, key, value, time.Minute)
	assert.NoError(t, err)

	// Verify value exists
	var result string
	err = cache.Get(ctx, key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)

	// Delete value
	err = cache.Delete(ctx, key)
	assert.NoError(t, err)

	// Verify value no longer exists
	err = cache.Get(ctx, key, &result)
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err)
}

func TestDistributedCache_Clear(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()

	// Set multiple values
	keys := []string{"clear_key1", "clear_key2", "clear_key3"}
	for _, key := range keys {
		err := cache.Set(ctx, key, "value", time.Minute)
		assert.NoError(t, err)
	}

	// Clear all keys matching pattern
	err := cache.Clear(ctx, "clear_*")
	assert.NoError(t, err)

	// Verify all keys are deleted
	for _, key := range keys {
		var result string
		err = cache.Get(ctx, key, &result)
		assert.Error(t, err)
		assert.Equal(t, redis.Nil, err)
	}
}

func TestDistributedCache_GetNonExistentKey(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "non_existent_key"

	var result string
	err := cache.Get(ctx, key, &result)
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err)
}

func TestDistributedCache_SetComplexObject(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "complex_object"

	type ComplexObject struct {
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Tags   []string `json:"tags"`
		Active bool     `json:"active"`
	}

	originalObject := ComplexObject{
		ID:     123,
		Name:   "Test Object",
		Tags:   []string{"tag1", "tag2", "tag3"},
		Active: true,
	}

	// Set complex object
	err := cache.Set(ctx, key, originalObject, time.Minute)
	assert.NoError(t, err)

	// Get complex object
	var retrievedObject ComplexObject
	err = cache.Get(ctx, key, &retrievedObject)
	assert.NoError(t, err)
	assert.Equal(t, originalObject, retrievedObject)
}

func TestDistributedCache_ConcurrentOperations(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	numOperations := 100

	// Run concurrent set operations
	done := make(chan bool, numOperations)
	for i := 0; i < numOperations; i++ {
		go func(i int) {
			key := fmt.Sprintf("concurrent_key_%d", i)
			value := fmt.Sprintf("concurrent_value_%d", i)
			err := cache.Set(ctx, key, value, time.Minute)
			assert.NoError(t, err)
			done <- true
		}(i)
	}

	// Wait for all operations to complete
	for i := 0; i < numOperations; i++ {
		<-done
	}

	// Verify all values were set correctly
	for i := 0; i < numOperations; i++ {
		key := fmt.Sprintf("concurrent_key_%d", i)
		expectedValue := fmt.Sprintf("concurrent_value_%d", i)

		var actualValue string
		err := cache.Get(ctx, key, &actualValue)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, actualValue)
	}
}

func TestDistributedCache_Namespace(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()
	key := "namespace_key"
	value := "namespace_value"

	// Set value (should be prefixed with namespace)
	err := cache.Set(ctx, key, value, time.Minute)
	assert.NoError(t, err)

	// Verify the key exists with the namespace prefix in Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: miniredis.Addr(),
	})
	defer redisClient.Close()

	namespacedKey := "test:" + key
	exists := redisClient.Exists(ctx, namespacedKey)
	assert.Equal(t, int64(1), exists.Val())

	// Get value through cache (should handle namespace automatically)
	var result string
	err = cache.Get(ctx, key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestCacheStrategy_WriteThrough(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	// WriteThrough strategy should write to both cache and backing store
	// For this test, we'll just verify the cache behavior
	ctx := context.Background()
	key := "write_through_key"
	value := "write_through_value"

	err := cache.Set(ctx, key, value, time.Minute)
	assert.NoError(t, err)

	var result string
	err = cache.Get(ctx, key, &result)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestDistributedCache_InvalidKey(t *testing.T) {
	cache, miniredis := createTestDistributedCache(t)
	defer miniredis.Close()

	ctx := context.Background()

	// Test with empty key
	err := cache.Set(ctx, "", "value", time.Minute)
	assert.Error(t, err)

	// Test with very long key (exceeding MaxKeySize)
	longKey := string(make([]byte, 2000))
	err = cache.Set(ctx, longKey, "value", time.Minute)
	assert.Error(t, err)
}
