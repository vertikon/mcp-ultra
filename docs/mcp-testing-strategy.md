# 📊 MCP Testing Strategy Enhancement

## Test Pyramid Completa

### 1. Unit Tests (70% cobertura)
```go
// test/unit/domain/service_test.go
func TestTaskService_CreateTask(t *testing.T) {
    tests := []struct {
        name    string
        input   domain.Task
        mockErr error
        wantErr bool
    }{
        {
            name:    "success with valid input",
            input:   domain.Task{Title: "Valid", Owner: "user1"},
            mockErr: nil,
            wantErr: false,
        },
        {
            name:    "fail with empty title",
            input:   domain.Task{Title: "", Owner: "user1"},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Arrange
            mockRepo := &mocks.Repository{}
            mockAudit := &mocks.Auditor{}
            
            if tt.mockErr == nil && tt.input.Title != "" {
                mockRepo.On("CreateTask", mock.Anything, mock.Anything).
                    Return(tt.input, nil)
                mockAudit.On("Emit", mock.Anything, "task.created", mock.Anything).
                    Return(nil)
            }
            
            // Act
            result, err := domain.CreateTask(context.TODO(), mockRepo, mockAudit, tt.input)
            
            // Assert
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.input.Title, result.Title)
                mockRepo.AssertExpectations(t)
                mockAudit.AssertExpectations(t)
            }
        })
    }
}
```

### 2. Integration Tests (com containers)
```go
// test/integration/repository_test.go
func TestPostgresRepository_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    // Setup testcontainers
    ctx := context.Background()
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:16"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
        testcontainers.WithWaitStrategy(
            wait.ForLog("database system is ready to accept connections").
                WithOccurrence(2).
                WithStartupTimeout(5 * time.Second),
        ),
    )
    require.NoError(t, err)
    defer pgContainer.Terminate(ctx)
    
    // Get connection string
    connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
    require.NoError(t, err)
    
    // Run migrations
    m, err := migrate.New(
        "file://../../internal/repository/pg/migrations",
        connStr,
    )
    require.NoError(t, err)
    require.NoError(t, m.Up())
    
    // Test repository
    repo, err := pg.New(connStr)
    require.NoError(t, err)
    
    task := domain.Task{
        ID:    uuid.New().String(),
        Title: "Integration Test",
        Owner: "tester",
    }
    
    created, err := repo.CreateTask(ctx, task)
    assert.NoError(t, err)
    assert.Equal(t, task.Title, created.Title)
}
```

### 3. Contract Tests (Schema validation)
```go
// test/contract/events_test.go
func TestEventContract_NATS(t *testing.T) {
    // Load Avro schema
    schema, err := avro.ParseFiles("../../api/schemas/event.avsc")
    require.NoError(t, err)
    
    // Test event serialization
    event := domain.Envelope{
        ID:      uuid.New().String(),
        TS:      time.Now(),
        Type:    "task.created",
        TraceID: "trace-123",
        Payload: map[string]interface{}{
            "id":    "task-123",
            "title": "Test Task",
        },
    }
    
    // Validate against schema
    native, err := avro.Generic(schema, event)
    require.NoError(t, err)
    
    // Serialize
    binary, err := avro.Marshal(schema, native)
    require.NoError(t, err)
    assert.NotEmpty(t, binary)
    
    // Deserialize and verify
    var decoded domain.Envelope
    err = avro.Unmarshal(schema, binary, &decoded)
    require.NoError(t, err)
    assert.Equal(t, event.Type, decoded.Type)
}
```

### 4. E2E Tests (API completa)
```go
// test/e2e/api_test.go
func TestAPI_E2E_TaskFlow(t *testing.T) {
    // Start services
    pool, err := dockertest.NewPool("")
    require.NoError(t, err)
    
    // Start dependencies
    postgres := startPostgres(t, pool)
    nats := startNATS(t, pool)
    redis := startRedis(t, pool)
    
    // Start application
    os.Setenv("POSTGRES_DSN", postgres.DSN)
    os.Setenv("NATS_URL", nats.URL)
    os.Setenv("REDIS_ADDR", redis.Addr)
    
    go main() // Start the actual service
    waitForHealthy(t, "http://localhost:9655/healthz")
    
    // Test flow
    client := &http.Client{Timeout: 10 * time.Second}
    
    // 1. Create task
    payload := `{"title":"E2E Task","owner":"e2e-user"}`
    resp, err := client.Post(
        "http://localhost:9655/api/v1/tasks",
        "application/json",
        strings.NewReader(payload),
    )
    require.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    
    var created map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&created)
    taskID := created["id"].(string)
    
    // 2. Verify event was published
    verifyEventPublished(t, nats, "task.created", taskID)
    
    // 3. Get task
    resp, err = client.Get(fmt.Sprintf("http://localhost:9655/api/v1/tasks/%s", taskID))
    require.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    
    // 4. Verify metrics
    metrics := getMetrics(t, "http://localhost:9655/metrics")
    assert.Contains(t, metrics, "http_request_duration_seconds")
}
```

### 5. Mutation Tests (com go-mutesting)
```bash
# Makefile addition
test-mutation:
	go get -u github.com/zimmski/go-mutesting/...
	go-mutesting ./internal/domain/... --verbose
```

### 6. Property-Based Tests (com gopter)
```go
// test/property/domain_test.go
func TestTaskProperties(t *testing.T) {
    properties := gopter.NewProperties(nil)
    
    properties.Property("task with non-empty title always succeeds", prop.ForAll(
        func(title string, owner string) bool {
            if title == "" {
                return true // Skip empty titles
            }
            
            task := domain.Task{Title: title, Owner: owner}
            result, err := domain.ValidateTask(task)
            
            return err == nil && result.Title == title
        },
        gen.AlphaString(),
        gen.AlphaString(),
    ))
    
    properties.TestingRun(t)
}
```

### 7. Performance Tests
```go
// test/benchmark/service_bench_test.go
func BenchmarkCreateTask(b *testing.B) {
    repo := setupMockRepo()
    audit := setupMockAudit()
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            task := domain.Task{
                Title: "Benchmark Task",
                Owner: "bench-user",
            }
            _, _ = domain.CreateTask(context.Background(), repo, audit, task)
        }
    })
    
    b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "tasks/sec")
}
```

### 8. Chaos Tests (com Litmus)
```yaml
# test/chaos/network-delay.yaml
apiVersion: litmuschaos.io/v1alpha1
kind: ChaosEngine
metadata:
  name: mcp-chaos
spec:
  appinfo:
    appns: default
    applabel: app=mcp-ultra-reference
  chaosServiceAccount: litmus-admin
  experiments:
    - name: pod-network-latency
      spec:
        components:
          env:
            - name: NETWORK_LATENCY
              value: '2000' # 2s latency
            - name: DURATION
              value: '60'
```

### 9. Security Tests
```go
// test/security/api_security_test.go
func TestAPI_Security(t *testing.T) {
    tests := []struct {
        name       string
        endpoint   string
        method     string
        headers    map[string]string
        wantStatus int
    }{
        {
            name:       "SQL injection attempt",
            endpoint:   "/api/v1/tasks?filter='; DROP TABLE tasks;--",
            method:     "GET",
            wantStatus: http.StatusBadRequest,
        },
        {
            name:       "XSS attempt",
            endpoint:   "/api/v1/tasks",
            method:     "POST",
            headers:    map[string]string{"Content-Type": "application/json"},
            wantStatus: http.StatusBadRequest,
        },
        {
            name:       "Missing auth token",
            endpoint:   "/api/v1/tasks",
            method:     "GET",
            wantStatus: http.StatusUnauthorized,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Test Automation Pipeline

```yaml
# .github/workflows/test-pipeline.yml
name: Complete Test Pipeline
on: [push, pull_request]

jobs:
  unit-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
      - run: go test ./... -coverprofile=coverage.out -short
      - uses: codecov/codecov-action@v3
  
  integration-tests:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
      nats:
        image: nats:2.10
        ports: ["4222:4222"]
    steps:
      - uses: actions/checkout@v4
      - run: go test ./test/integration/... -tags=integration
  
  contract-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: go test ./test/contract/... -tags=contract
  
  e2e-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: docker-compose -f docker-compose.test.yml up -d
      - run: go test ./test/e2e/... -tags=e2e
      - run: docker-compose -f docker-compose.test.yml down
  
  performance-tests:
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v4
      - run: go test -bench=. ./test/benchmark/...
      - uses: benchmark-action/github-action-benchmark@v1
  
  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: aquasecurity/trivy-action@master
      - run: gosec ./...
      - run: nancy sleuth # Check dependencies
```

## Coverage Goals

```yaml
coverage:
  unit: 80%
  integration: 70%
  e2e: 60%
  mutation: 50%
  
quality-gates:
  - coverage >= 75%
  - no critical security issues
  - performance regression < 10%
  - all contract tests passing
```