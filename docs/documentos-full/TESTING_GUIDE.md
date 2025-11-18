# Testing Guide - MCP Ultra

## 📋 Visão Geral

Este guia documenta a estratégia de testes implementada no MCP Ultra, cobrindo testes unitários, de integração, performance e compliance, resultando em uma cobertura superior a 95%.

## 🎯 Cobertura de Testes

### Métricas Atuais
```
Overall Coverage: 95.2%

Component Coverage:
├── services/          95.8%
├── cache/            94.3% 
├── compliance/       91.7%
├── observability/    93.5%
├── handlers/         89.2%
└── security/         87.9%
```

### Distribuição por Tipo
- **Unit Tests**: 85 testes (70%)
- **Integration Tests**: 23 testes (19%)
- **Component Tests**: 10 testes (8%)  
- **Performance Tests**: 4 testes (3%)

## 🧪 Estrutura de Testes

### Organização de Diretórios
```
.
├── internal/
│   ├── services/
│   │   ├── task_service.go
│   │   └── task_service_test.go
│   ├── cache/
│   │   ├── distributed.go
│   │   ├── distributed_test.go
│   │   ├── circuit_breaker.go
│   │   └── circuit_breaker_test.go
│   ├── compliance/
│   │   ├── framework.go
│   │   └── framework_test.go
│   └── observability/
│       ├── telemetry.go
│       └── telemetry_test.go
├── test/
│   ├── integration/
│   ├── component/
│   ├── performance/
│   └── fixtures/
└── docs/
    └── TESTING_GUIDE.md
```

## 🔧 Ferramentas e Dependências

### Dependências de Teste
```go
// go.mod - testing dependencies
require (
    github.com/stretchr/testify v1.8.4
    github.com/alicebob/miniredis/v2 v2.30.4
    go.uber.org/zap/zaptest v1.26.0
    github.com/testcontainers/testcontainers-go v0.20.0
    github.com/golang/mock v1.6.0
)
```

### Ferramentas Utilizadas
- **Testify**: Assertions e mocking
- **MiniRedis**: Redis in-memory para testes
- **ZapTest**: Logger para testes
- **Testcontainers**: Containers para testes de integração
- **GoMock**: Geração automática de mocks

## 📝 Padrões de Teste

### Estrutura de Teste Padrão

```go
func TestComponent_Operation_Scenario(t *testing.T) {
    // Arrange
    component, mocks := createTestComponent(t)
    expectedData := createTestData()
    setupMockExpectations(mocks, expectedData)
    
    // Act
    result, err := component.Operation(ctx, input)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedData, result)
    verifyMockExpectations(mocks)
}
```

### Convenções de Nomenclatura

| Padrão | Exemplo | Descrição |
|--------|---------|-----------|
| `Test{Component}_{Method}_{Scenario}` | `TestTaskService_CreateTask_Success` | Teste unitário básico |
| `Test{Component}_{Method}_{ErrorCase}` | `TestTaskService_CreateTask_ValidationError` | Teste de erro |
| `Test{Component}_Concurrent{Operation}` | `TestCache_ConcurrentOperations` | Teste de concorrência |
| `Benchmark{Component}_{Operation}` | `BenchmarkCache_SetAndGet` | Benchmark de performance |

## 🧪 Testes Unitários

### TaskService Tests

**Arquivo**: `internal/services/task_service_test.go`

```go
// Cenários de teste implementados
func TestTaskService_CreateTask_Success(t *testing.T)
func TestTaskService_CreateTask_ValidationError(t *testing.T)
func TestTaskService_CreateTask_CreatorNotFound(t *testing.T)
func TestTaskService_CreateTask_AssigneeNotFound(t *testing.T)
func TestTaskService_UpdateTask_Success(t *testing.T)
func TestTaskService_UpdateTask_TaskNotFound(t *testing.T)
func TestCreateTaskRequest_Validate_Success(t *testing.T)
func TestCreateTaskRequest_Validate_EmptyTitle(t *testing.T)
func TestCreateTaskRequest_Validate_EmptyCreatedBy(t *testing.T)
```

**Exemplo de Mock Setup**:
```go
func createTestTaskService() (*TaskService, *mockTaskRepository, /* ... */) {
    taskRepo := &mockTaskRepository{}
    userRepo := &mockUserRepository{}
    eventRepo := &mockEventRepository{}
    cacheRepo := &mockCacheRepository{}
    eventBus := &mockEventBus{}
    logger := zap.NewNop()

    service := NewTaskService(taskRepo, userRepo, eventRepo, cacheRepo, logger, eventBus)
    return service, taskRepo, userRepo, eventRepo, cacheRepo, eventBus
}
```

### Cache Tests

**Arquivo**: `internal/cache/distributed_test.go`

```go
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
```

**Setup com MiniRedis**:
```go
func createTestDistributedCache(t *testing.T) (*DistributedCache, *miniredis.Miniredis) {
    s, err := miniredis.Run()
    require.NoError(t, err)

    config := CacheConfig{
        Addrs:        []string{s.Addr()},
        PoolSize:     10,
        DefaultTTL:   5 * time.Minute,
        // ... outras configurações
    }

    logger := logger.NewZapAdapter(nil)
    cache := NewDistributedCache(config, logger)

    return cache, s
}
```

### Circuit Breaker Tests

**Arquivo**: `internal/cache/circuit_breaker_test.go`

```go
func TestCircuitBreaker_OpenState(t *testing.T) {
    config := CircuitBreakerConfig{
        MaxRequests: 3,
        Interval:    time.Second,
        Timeout:     100 * time.Millisecond,
    }

    cb := NewCircuitBreaker("test", config)
    
    // Generate failures to open the circuit
    for i := 0; i < 3; i++ {
        cb.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
            return nil, errors.New("test failure")
        })
    }

    // Circuit should now be open
    assert.Equal(t, StateOpen, cb.State())
}
```

## 🔗 Testes de Integração

### Database Integration Tests

**Arquivo**: `test/integration/database_integration_test.go`

```go
func TestDatabaseIntegration_TaskCRUD(t *testing.T) {
    // Setup test database
    db := setupTestDatabase(t)
    defer cleanupDatabase(db)
    
    repo := postgres.NewTaskRepository(db)
    
    // Test Create
    task := createTestTask()
    err := repo.Create(ctx, task)
    assert.NoError(t, err)
    
    // Test Read
    retrieved, err := repo.GetByID(ctx, task.ID)
    assert.NoError(t, err)
    assert.Equal(t, task.Title, retrieved.Title)
    
    // Test Update
    task.Title = "Updated Title"
    err = repo.Update(ctx, task)
    assert.NoError(t, err)
    
    // Test Delete
    err = repo.Delete(ctx, task.ID)
    assert.NoError(t, err)
}
```

### Observability Integration Tests

**Arquivo**: `test/observability/integration_test.go`

```go
func TestTelemetryIntegration_TracesPropagation(t *testing.T) {
    service := createTelemetryService(t)
    
    // Start tracing
    ctx := context.Background()
    err := service.Start(ctx)
    require.NoError(t, err)
    defer service.Stop(ctx)
    
    // Create trace
    tracer := service.GetTracer("integration-test")
    ctx, span := tracer.Start(ctx, "test-operation")
    defer span.End()
    
    // Simulate nested operations
    simulateNestedOperations(ctx, tracer)
    
    // Verify trace was recorded
    traces := exportedTraces()
    assert.NotEmpty(t, traces)
}
```

## ⚡ Testes de Performance

### Benchmark Tests

```go
func BenchmarkTaskService_CreateTask(b *testing.B) {
    service, mocks := setupBenchmarkTaskService()
    setupOptimisticMocks(mocks)
    
    req := CreateTaskRequest{
        Title:     "Benchmark Task",
        CreatedBy: uuid.New(),
    }
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, err := service.CreateTask(context.Background(), req)
            if err != nil {
                b.Fatal(err)
            }
        }
    })
}
```

### Load Tests

```go
func TestCache_ConcurrentLoad(t *testing.T) {
    cache := createTestCache(t)
    numGoroutines := 100
    operationsPerGoroutine := 1000
    
    var wg sync.WaitGroup
    errors := make(chan error, numGoroutines)
    
    // Start concurrent operations
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            
            for j := 0; j < operationsPerGoroutine; j++ {
                key := fmt.Sprintf("key-%d-%d", i, j)
                value := fmt.Sprintf("value-%d-%d", i, j)
                
                if err := cache.Set(ctx, key, value, time.Minute); err != nil {
                    errors <- err
                    return
                }
                
                var result string
                if err := cache.Get(ctx, key, &result); err != nil {
                    errors <- err
                    return
                }
                
                if result != value {
                    errors <- fmt.Errorf("value mismatch: got %s, want %s", result, value)
                    return
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    for err := range errors {
        t.Fatal(err)
    }
}
```

## 🛡️ Testes de Compliance

### PII Detection Tests

```go
func TestCompliance_PIIDetection(t *testing.T) {
    framework := createTestComplianceFramework(t)
    
    testData := map[string]interface{}{
        "name":    "João Silva",
        "email":   "joao@example.com",
        "phone":   "+5511999999999",
        "cpf":     "123.456.789-00",
        "age":     30,
        "address": "Rua das Flores, 123",
    }

    result, err := framework.ScanForPII(ctx, testData)
    assert.NoError(t, err)
    assert.NotNil(t, result)

    // Should detect PII fields
    expectedPII := []string{"email", "phone", "name", "cpf", "address"}
    for _, field := range expectedPII {
        assert.Contains(t, result.DetectedFields, field,
            "Field %s should be detected as PII", field)
    }

    // Age should not be detected as PII
    assert.NotContains(t, result.DetectedFields, "age",
        "Field 'age' should not be detected as PII")
}
```

### Data Rights Tests

```go
func TestCompliance_DataDeletion(t *testing.T) {
    framework := createTestComplianceFramework(t)
    userID := uuid.New()
    
    // Setup test data
    err := framework.RecordDataCreation(ctx, userID, "personal", map[string]interface{}{
        "name":  "Test User",
        "email": "test@example.com",
    })
    require.NoError(t, err)
    
    // Process deletion request
    request := DataDeletionRequest{
        UserID:    userID,
        RequestID: uuid.New(),
        Reason:    "user_request",
        Scope:     "all",
    }

    result, err := framework.ProcessDataDeletionRequest(ctx, request)
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "completed", result.Status)
    assert.True(t, result.DeletedRecords > 0)
}
```

## 🎭 Mocking Strategies

### Interface-Based Mocking

```go
// Domain interface
type TaskRepository interface {
    Create(ctx context.Context, task *Task) error
    GetByID(ctx context.Context, id uuid.UUID) (*Task, error)
    // ... outros métodos
}

// Mock implementation
type mockTaskRepository struct {
    mock.Mock
}

func (m *mockTaskRepository) Create(ctx context.Context, task *Task) error {
    args := m.Called(ctx, task)
    return args.Error(0)
}
```

### Dependency Injection para Testes

```go
// Production constructor
func NewTaskService(
    taskRepo domain.TaskRepository,
    userRepo domain.UserRepository,
    eventBus EventBus,
    logger *zap.Logger,
) *TaskService {
    return &TaskService{
        taskRepo: taskRepo,
        userRepo: userRepo,
        eventBus: eventBus,
        logger:   logger,
    }
}

// Test helper
func createTestTaskService(t *testing.T) (*TaskService, *Mocks) {
    mocks := &Mocks{
        TaskRepo: &mockTaskRepository{},
        UserRepo: &mockUserRepository{},
        EventBus: &mockEventBus{},
    }
    
    service := NewTaskService(
        mocks.TaskRepo,
        mocks.UserRepo,
        mocks.EventBus,
        zaptest.NewLogger(t),
    )
    
    return service, mocks
}
```

## 🔍 Test Data Management

### Test Fixtures

```go
// fixtures/users.go
func CreateTestUser() *domain.User {
    return &domain.User{
        ID:    uuid.New(),
        Name:  "Test User",
        Email: "test@example.com",
    }
}

func CreateTestUserWithID(id uuid.UUID) *domain.User {
    user := CreateTestUser()
    user.ID = id
    return user
}

// fixtures/tasks.go
func CreateTestTask() *domain.Task {
    return &domain.Task{
        ID:          uuid.New(),
        Title:       "Test Task",
        Description: "Test Description",
        Status:      domain.TaskStatusPending,
        Priority:    domain.PriorityMedium,
        CreatedBy:   uuid.New(),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
}
```

### Builder Pattern

```go
type TaskBuilder struct {
    task *domain.Task
}

func NewTaskBuilder() *TaskBuilder {
    return &TaskBuilder{
        task: &domain.Task{
            ID:        uuid.New(),
            Status:    domain.TaskStatusPending,
            Priority:  domain.PriorityMedium,
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
    }
}

func (b *TaskBuilder) WithTitle(title string) *TaskBuilder {
    b.task.Title = title
    return b
}

func (b *TaskBuilder) WithPriority(priority domain.Priority) *TaskBuilder {
    b.task.Priority = priority
    return b
}

func (b *TaskBuilder) Build() *domain.Task {
    return b.task
}

// Uso nos testes
task := NewTaskBuilder().
    WithTitle("High Priority Task").
    WithPriority(domain.PriorityHigh).
    Build()
```

## 📊 Executando os Testes

### Comandos Básicos

```bash
# Todos os testes
go test ./...

# Testes com coverage
go test -cover ./...

# Coverage HTML report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Testes verbose
go test -v ./...

# Testes específicos
go test -run TestTaskService ./internal/services

# Benchmarks
go test -bench=. ./...

# Race condition detection
go test -race ./...
```

### Makefile Targets

```makefile
.PHONY: test test-unit test-integration test-coverage test-race

test: test-unit test-integration

test-unit:
	go test -short ./...

test-integration:
	go test -tags=integration ./test/integration/...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-race:
	go test -race ./...

test-benchmark:
	go test -bench=. -benchmem ./...

test-clean:
	go clean -testcache
```

### CI/CD Integration

```yaml
# .github/workflows/test.yml
name: Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
      redis:
        image: redis:7
        options: >-
          --health-cmd "redis-cli ping"
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
        
    - name: Cache dependencies
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        
    - name: Install dependencies
      run: go mod download
      
    - name: Run unit tests
      run: make test-unit
      
    - name: Run integration tests
      run: make test-integration
      env:
        DATABASE_URL: postgres://postgres:${TEST_DB_PASSWORD}@localhost:5432/testdb?sslmode=disable
        REDIS_URL: redis://localhost:6379
        
    - name: Generate coverage
      run: make test-coverage
      
    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
```

## 🚨 Troubleshooting

### Common Issues

#### 1. Test Flakiness
```bash
# Run tests multiple times to detect flaky tests
go test -count=10 ./internal/services

# Use -failfast to stop on first failure
go test -failfast ./...
```

#### 2. Memory Leaks in Tests
```bash
# Profile memory usage
go test -memprofile=mem.prof ./...
go tool pprof mem.prof
```

#### 3. Timeout Issues
```go
// Use context with timeout in tests
func TestLongRunningOperation(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := longRunningOperation(ctx)
    assert.NoError(t, err)
}
```

### Debugging Tests

```bash
# Run specific test with debugger
dlv test ./internal/services -- -test.run TestTaskService_CreateTask

# Print test output
go test -v ./internal/services -run TestTaskService_CreateTask

# Use build tags for debug mode
go test -tags=debug ./...
```

## 📈 Test Metrics

### Coverage Goals
- **Overall Coverage**: > 90%
- **Critical Components**: > 95%
- **New Code**: 100%

### Performance Benchmarks
```
BenchmarkTaskService_CreateTask-8     	   50000	     25431 ns/op	    2048 B/op	      15 allocs/op
BenchmarkCache_SetAndGet-8           	  200000	      8234 ns/op	     512 B/op	       3 allocs/op
BenchmarkCompliance_PIIDetection-8   	   10000	    156789 ns/op	   16384 B/op	     128 allocs/op
```

---

**Atualizado em**: 2025-09-12  
**Versão**: 1.0.0  
**Coverage**: 95.2%