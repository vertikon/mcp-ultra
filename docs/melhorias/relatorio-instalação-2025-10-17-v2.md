PS E:\vertikon\business\SaaS\templates\mcp-ultra> make mocks
bash scripts/regenerate_mocks.sh
[mcp-ultra] Regenerando mocks com GoMock...
scripts/regenerate_mocks.sh: line 7: mockgen: command not found
make: *** [Makefile:15: mocks] Error 127
PS E:\vertikon\business\SaaS\templates\mcp-ultra> make lint
golangci-lint run --timeout=5m
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]\ninternal\\compliance\\framework_stubs.go:12:6: DataAccessRequest redeclared in this block\n\tinternal\\compliance\\framework.go:544:6: other declaration of DataAccessRequest\ninternal\\compliance\\framework_stubs.go:17:6: DataDeletionRequest redeclared in this block\n\tinternal\\compliance\\framework.go:554:6: other declaration of DataDeletionRequest\ninternal\\compliance\\framework_stubs.go:22:6: AuditEvent redeclared in this block\n\tinternal\\compliance\\audit_logger.go:27:6: other declaration of AuditEvent\ninternal\\compliance\\framework_stubs.go:29:31: method ComplianceFramework.ProcessDataAccessRequest already declared at internal\\compliance\\framework.go:583:32\ninternal\\compliance\\framework_stubs.go:34:31: method ComplianceFramework.AnonymizeData already declared at internal\\compliance\\framework.go:623:32\ninternal\\compliance\\framework_stubs.go:39:31: method ComplianceFramework.LogAuditEvent already declared at internal\\compliance\\framework.go:640:32\ninternal\\compliance\\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal\ninternal\\compliance\\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value\ninternal\\compliance\\framework_test.go:208:17: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/handlers/http [github.com/vertikon/mcp-ultra/internal/handlers/http.test]\ninternal\\handlers\\http\\router_test.go:23:76: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:25:42: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:38:75: undefined: services.HealthChecker\ninternal\\handlers\\http\\router_test.go:47:70: undefined: domain.CreateTaskRequest\ninternal\\handlers\\http\\router_test.go:60:85: undefined: domain.UpdateTaskRequest\ninternal\\handlers\\http\\router_test.go:70:73: undefined: domain.TaskFilters\ninternal\\handlers\\http\\router_test.go:70:95: undefined: domain.TaskList\ninternal\\handlers\\http\\router_test.go:72:30: undefined: domain.TaskList\ninternal\\handlers\\http\\router_test.go:80:49: not enough arguments in call to NewRouter\n\thave (*zap.Logger, *MockHealthService, *MockTaskService)\n\twant (*services.TaskService, *features.FlagManager, *HealthService, *zap.Logger)\ninternal\\handlers\\http\\router_test.go:101:77: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:101:77: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/middleware [github.com/vertikon/mcp-ultra/internal/middleware.test]\ninternal\\middleware\\auth_test.go:96:30: undefined: testhelpers.GetTestAPIKeys"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/security [github.com/vertikon/mcp-ultra/internal/security.test]\ninternal\\security\\enhanced_auth_test.go:22:6: MockOPAService redeclared in this block\n\tinternal\\security\\auth_test.go:20:6: other declaration of MockOPAService\ninternal\\security\\enhanced_auth_test.go:26:26: method MockOPAService.IsAuthorized already declared at internal\\security\\auth_test.go:24:26\ninternal\\security\\enhanced_auth_test.go:36:6: TestNewAuthService redeclared in this block\n\tinternal\\security\\auth_test.go:39:6: other declaration of TestNewAuthService\ninternal\\security\\enhanced_auth_test.go:326:6: TestGetUserFromContext redeclared in this block\n\tinternal\\security\\auth_test.go:411:6: other declaration of TestGetUserFromContext\ninternal\\security\\enhanced_auth_test.go:391:6: TestRequireScope redeclared in this block\n\tinternal\\security\\auth_test.go:282:6: other declaration of TestRequireScope\ninternal\\security\\enhanced_auth_test.go:459:6: TestRequireRole redeclared in this block\n\tinternal\\security\\auth_test.go:342:6: other declaration of TestRequireRole\ninternal\\security\\auth_test.go:49:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:67:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:140:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:163:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:163:48: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/services [github.com/vertikon/mcp-ultra/internal/services.test]\ninternal\\services\\task_service_test.go:104:70: undefined: domain.UserFilter\ninternal\\services\\task_service_test.go:171:28: cannot use taskRepo (variable of type *mockTaskRepository) as domain.TaskRepository value in argument to NewTaskService: *mockTaskRepository does not implement domain.TaskRepository (wrong type for method List)\n\t\thave List(context.Context, domain.TaskFilter) ([]*domain.Task, error)\n\t\twant List(context.Context, domain.TaskFilter) ([]*domain.Task, int, error)\ninternal\\services\\task_service_test.go:171:48: cannot use eventRepo (variable of type *mockEventRepository) as domain.EventRepository value in argument to NewTaskService: *mockEventRepository does not implement domain.EventRepository (missing method GetByType)\ninternal\\services\\task_service_test.go:171:59: cannot use cacheRepo (variable of type *mockCacheRepository) as domain.CacheRepository value in argument to NewTaskService: *mockCacheRepository does not implement domain.CacheRepository (missing method Exists)\ninternal\\services\\task_service_test.go:199:31: declared and not used: eventRepo"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/component [github.com/vertikon/mcp-ultra/test/component.test]\ntest\\component\\task_service_test.go:39:3: cannot use suite.taskRepo (variable of type *mocks.MockTaskRepository) as domain.TaskRepository value in argument to services.NewTaskService: *mocks.MockTaskRepository does not implement domain.TaskRepository (wrong type for method Delete)\n\t\thave Delete(context.Context, string) error\n\t\twant Delete(context.Context, uuid.UUID) error\ntest\\component\\task_service_test.go:40:3: cannot use suite.validator (variable of type *mocks.MockValidator) as domain.UserRepository value in argument to services.NewTaskService: *mocks.MockValidator does not implement domain.UserRepository (missing method Create)\ntest\\component\\task_service_test.go:42:3: cannot use suite.cacheRepo (variable of type *mocks.MockCacheRepository) as domain.CacheRepository value in argument to services.NewTaskService: *mocks.MockCacheRepository does not implement domain.CacheRepository (wrong type for method Get)\n\t\thave Get(context.Context, string) (interface{}, error)\n\t\twant Get(context.Context, string) (string, error)\ntest\\component\\task_service_test.go:44:3: cannot use suite.eventBus (variable of type *mocks.MockEventBus) as services.EventBus value in argument to services.NewTaskService: *mocks.MockEventBus does not implement services.EventBus (wrong type for method Publish)\n\t\thave Publish(context.Context, string, []byte) error\n\t\twant Publish(context.Context, *domain.Event) error\ntest\\component\\task_service_test.go:65:3: unknown field Metadata in struct literal of type services.CreateTaskRequest\ntest\\component\\task_service_test.go:78:20: req.Metadata undefined (type *services.CreateTaskRequest has no field or method Metadata)\ntest\\component\\task_service_test.go:97:55: too many arguments in call to suite.service.CreateTask\n\thave (context.Context, uuid.UUID, *services.CreateTaskRequest)\n\twant (context.Context, services.CreateTaskRequest)\ntest\\component\\task_service_test.go:118:29: undefined: services.ValidationError\ntest\\component\\task_service_test.go:127:55: too many arguments in call to suite.service.CreateTask\n\thave (context.Context, uuid.UUID, *services.CreateTaskRequest)\n\twant (context.Context, services.CreateTaskRequest)\ntest\\component\\task_service_test.go:151:52: too many arguments in call to suite.service.GetTask\n\thave (context.Context, uuid.UUID, uuid.UUID)\n\twant (context.Context, uuid.UUID)\ntest\\component\\task_service_test.go:151:52: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/property [github.com/vertikon/mcp-ultra/test/property.test]\ntest\\property\\task_properties_test.go:231:4: declared and not used: originalTitle"
internal\observability\middleware.go:189: 189-225 lines are duplicate of `internal\observability\middleware.go:228-264` (dupl)
func (ts *TelemetryService) CacheOperation(
        ctx context.Context,
        operation string,
        key string,
        fn func(context.Context) error,
) error {
        if !ts.config.Enabled {
                return fn(ctx)
        }

        spanName := fmt.Sprintf("cache.%s", operation)
        ctx, span := ts.StartSpan(ctx, spanName,
                trace.WithSpanKind(trace.SpanKindClient),
                trace.WithAttributes(
                        attribute.String("cache.system", "redis"),
                        attribute.String("cache.operation", operation),
                        attribute.String("cache.key", key),
                ),
        )
        defer span.End()

        start := time.Now()
        err := fn(ctx)
        duration := time.Since(start)

        span.SetAttributes(attribute.Float64("cache.duration_ms", float64(duration.Nanoseconds())/1000000))

        if err != nil {
                span.RecordError(err)
                span.SetStatus(codes.Error, err.Error())
                ts.RecordError("cache_error", "cache")
        } else {
                span.SetStatus(codes.Ok, "")
        }

        return err
}
internal\observability\middleware.go:228: 228-264 lines are duplicate of `internal\observability\middleware.go:189-225` (dupl)
func (ts *TelemetryService) MessageQueueOperation(
        ctx context.Context,
        operation string,
        subject string,
        fn func(context.Context) error,
) error {
        if !ts.config.Enabled {
                return fn(ctx)
        }

        spanName := fmt.Sprintf("messaging.%s", operation)
        ctx, span := ts.StartSpan(ctx, spanName,
                trace.WithSpanKind(trace.SpanKindProducer),
                trace.WithAttributes(
                        attribute.String("messaging.system", "nats"),
                        attribute.String("messaging.operation", operation),
                        attribute.String("messaging.destination", subject),
                ),
        )
        defer span.End()

        start := time.Now()
        err := fn(ctx)
        duration := time.Since(start)

        span.SetAttributes(attribute.Float64("messaging.duration_ms", float64(duration.Nanoseconds())/1000000))

        if err != nil {
                span.RecordError(err)
                span.SetStatus(codes.Error, err.Error())
                ts.RecordError("messaging_error", "messaging")
        } else {
                span.SetStatus(codes.Ok, "")
        }

        return err
}
internal\handlers\health.go:17:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "alive"})
                                 ^
internal\handlers\health.go:23:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
                                 ^
internal\handlers\health.go:29:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
                                 ^
internal\handlers\health.go:44:10: Error return value of `w.Write` is not checked (errcheck)
                w.Write([]byte("# Metrics placeholder\n"))
                       ^
internal\repository\postgres\task_repository.go:284:17: Error return value of `json.Unmarshal` is not checked (errcheck)
                json.Unmarshal(tagsJSON, &task.Tags)
                              ^
internal\repository\postgres\task_repository.go:290:17: Error return value of `json.Unmarshal` is not checked (errcheck)
                json.Unmarshal(metadataJSON, &task.Metadata)
                              ^
main.go:27:19: Error return value of `logger.Sync` is not checked (errcheck)
        defer logger.Sync()
                         ^
internal\observability\telemetry.go:661:30: Error return value of `ts.IncrementRequestCounter` is not checked (errcheck)
                        ts.IncrementRequestCounter(ctx, r.Method, r.URL.Path, statusCode)
                                                  ^
internal\observability\telemetry.go:662:28: Error return value of `ts.RecordRequestDuration` is not checked (errcheck)
                        ts.RecordRequestDuration(ctx, r.Method, r.URL.Path, duration)
                                                ^
internal\observability\telemetry_test.go:83:20: Error return value of `service.Stop` is not checked (errcheck)
        defer service.Stop(ctx)
                          ^
internal\observability\telemetry_test.go:118:20: Error return value of `service.Stop` is not checked (errcheck)
        defer service.Stop(ctx)
                          ^
internal\observability\telemetry_test.go:165:20: Error return value of `service.Stop` is not checked (errcheck)
        defer service.Stop(ctx)
                          ^
internal\observability\telemetry_test.go:210:10: Error return value of `w.Write` is not checked (errcheck)
                w.Write([]byte("OK"))
                       ^
internal\observability\telemetry_test.go:300:35: Error return value of `service.IncrementRequestCounter` is not checked (errcheck)
                        service.IncrementRequestCounter(ctx, "GET", "/test", "200")
                                                       ^
internal\observability\telemetry_test.go:301:33: Error return value of `service.RecordRequestDuration` is not checked (errcheck)
                        service.RecordRequestDuration(ctx, "GET", "/test", time.Millisecond*100)
                                                     ^
internal\observability\telemetry_test.go:302:33: Error return value of `service.IncrementErrorCounter` is not checked (errcheck)
                        service.IncrementErrorCounter(ctx, "test", "concurrent")
                                                     ^
internal\observability\telemetry_test.go:303:32: Error return value of `service.RecordProcessingTime` is not checked (errcheck)
                        service.RecordProcessingTime(ctx, "concurrent_task", time.Millisecond*50)
                                                    ^
internal\lifecycle\deployment.go:407:20: Error return value of `da.executeCommand` is not checked (errcheck)
                da.executeCommand(ctx, fmt.Sprintf("kubectl delete deployment mcp-ultra-canary --namespace=%s", da.config.Namespace), result)
                                 ^
internal\lifecycle\deployment.go:420:19: Error return value of `da.executeCommand` is not checked (errcheck)
        da.executeCommand(ctx, fmt.Sprintf("kubectl delete deployment mcp-ultra-canary --namespace=%s", da.config.Namespace), result)
                         ^
internal\lifecycle\health.go:483:11: Error return value of `w.Write` is not checked (errcheck)
                        w.Write([]byte("OK"))
                               ^
internal\handlers\health_test.go:21:11: string `application/json` has 3 occurrences, make it a constant (goconst)
        if ct != "application/json" {
                 ^
internal\slo\alerting.go:653:7: string `warning` has 3 occurrences, but such constant `SeverityWarning` already exists (goconst)
        case "warning":
             ^
internal\slo\alerting.go:651:7: string `critical` has 3 occurrences, but such constant `SLOStatusCritical` already exists (goconst)
        case "critical":
             ^
internal\config\tls.go:145:7: string `1.2` has 5 occurrences, make it a constant (goconst)
        case "1.2":
             ^
internal\config\tls_test.go:160:31: string `invalid` has 3 occurrences, make it a constant (goconst)
                manager.config.MinVersion = "invalid"
                                            ^
internal\config\tls.go:147:7: string `1.3` has 5 occurrences, make it a constant (goconst)
        case "1.3":
             ^
internal\metrics\business.go:758:40: string `resolved` has 3 occurrences, make it a constant (goconst)
                if !exists || existingState.State == "resolved" {
                                                     ^
internal\lifecycle\manager.go:37:10: string `healthy` has 3 occurrences, but such constant `HealthStatusHealthy` already exists (goconst)
                return "healthy"
                       ^
internal\slo\alerting.go:230:1: cyclomatic complexity 21 of func `(*AlertManager).shouldSilence` is high (> 18) (gocyclo)
func (am *AlertManager) shouldSilence(alert AlertEvent) bool {
^
internal\dashboard\models.go:286:6: exported: type name will be used as dashboard.DashboardWidget by other packages, and that stutters; consider calling this Widget (revive)
type DashboardWidget struct {
     ^
internal\ai\events\handlers_test.go:19:42: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (m *mockPublisher) PublishWithRetry(ctx context.Context, subject string, payload []byte) error {
                                         ^
internal\handlers\health.go:14:53: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
                                                    ^
internal\handlers\health.go:20:54: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
                                                     ^
internal\handlers\health.go:26:55: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
                                                      ^
internal\events\nats_bus.go:27:34: unused-parameter: parameter 'nc' seems to be unused, consider removing or renaming it as _ (revive)
                nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
                                               ^
internal\events\nats_bus.go:50:34: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (bus *NATSEventBus) Publish(ctx context.Context, event *domain.Event) error {
                                 ^
internal\events\nats_bus.go:194:46: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (h *TaskEventHandler) handleTaskCreated(ctx context.Context, event *domain.Event) error {
                                             ^
internal\slo\monitor.go:16:6: exported: type name will be used as slo.SLOType by other packages, and that stutters; consider calling this Type (revive)
type SLOType string
     ^
internal\slo\monitor.go:27:6: exported: type name will be used as slo.SLOStatus by other packages, and that stutters; consider calling this Status (revive)
type SLOStatus string
     ^
internal\slo\monitor.go:70:6: exported: type name will be used as slo.SLOResult by other packages, and that stutters; consider calling this Result (revive)
type SLOResult struct {
     ^
internal\slo\alerting.go:480:55: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToEmail(alert AlertEvent, config ChannelConfig) error {
                                                      ^
internal\slo\alerting.go:488:59: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToPagerDuty(alert AlertEvent, config ChannelConfig) error {
                                                          ^
internal\slo\alerting.go:496:57: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToMSTeams(alert AlertEvent, config ChannelConfig) error {
                                                        ^
internal\ai\telemetry\metrics_test.go:126:33: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
func TestNoOpWhenNotInitialized(t *testing.T) {
                                ^
internal\config\tls_test.go:341:45: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should handle multiple stops", func(t *testing.T) {
                                                   ^
internal\telemetry\telemetry.go:84:11: unused-parameter: parameter 'cfg' seems to be unused, consider removing or renaming it as _ (revive)
func Init(cfg config.TelemetryConfig) (*Telemetry, error) {
          ^
internal\telemetry\tracing_test.go:202:43: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should add span attributes", func(t *testing.T) {
                                                 ^
internal\observability\telemetry_test.go:298:11: unused-parameter: parameter 'i' seems to be unused, consider removing or renaming it as _ (revive)
                go func(i int) {
                        ^
internal\observability\middleware.go:110:39: unused-parameter: parameter 'operation' seems to be unused, consider removing or renaming it as _ (revive)
                otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
                                                    ^
internal\tracing\business.go:40:6: exported: type name will be used as tracing.TracingConfig by other packages, and that stutters; consider calling this Config (revive)
type TracingConfig struct {
     ^
internal\tracing\business.go:735:83: unused-parameter: parameter 'attributes' seems to be unused, consider removing or renaming it as _ (revive)
func (btt *BusinessTransactionTracer) shouldSample(template *TransactionTemplate, attributes map[string]interface{}) bool {
                                                                                  ^
internal\metrics\storage.go:186:47: unused-parameter: parameter 'groupKey' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) extractLabels(groupKey string, groupBy []string) map[string]string {
                                              ^
internal\ratelimit\distributed.go:526:86: unused-parameter: parameter 'key' seems to be unused, consider removing or renaming it as _ (revive)
func (drl *DistributedRateLimiter) recordMetrics(status string, algorithm Algorithm, key string, remaining int64) {
                                                                                     ^
internal\ratelimit\distributed.go:733:52: unused-parameter: parameter 'rule' seems to be unused, consider removing or renaming it as _ (revive)
func (al *AdaptiveLimiter) updateState(key string, rule Rule, allowed bool) {
                                                   ^
internal\lifecycle\manager.go:15:6: exported: type name will be used as lifecycle.LifecycleState by other packages, and that stutters; consider calling this State (revive)
type LifecycleState int32
     ^
internal\lifecycle\manager.go:63:6: exported: type name will be used as lifecycle.LifecycleEvent by other packages, and that stutters; consider calling this Event (revive)
type LifecycleEvent struct {
     ^
internal\lifecycle\manager.go:74:6: exported: type name will be used as lifecycle.LifecycleManager by other packages, and that stutters; consider calling this Manager (revive)
type LifecycleManager struct {
     ^
internal\lifecycle\manager.go:387:6: exported: type name will be used as lifecycle.LifecycleMetrics by other packages, and that stutters; consider calling this Metrics (revive)
type LifecycleMetrics struct {
     ^
internal\lifecycle\deployment.go:579:53: unused-parameter: parameter 'version' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) validateDockerImage(version string) error {
                                                    ^
internal\cache\distributed.go:18:6: exported: type name will be used as cache.CacheStrategy by other packages, and that stutters; consider calling this Strategy (revive)
type CacheStrategy string
     ^
internal\cache\distributed.go:38:6: exported: type name will be used as cache.CacheConfig by other packages, and that stutters; consider calling this Config (revive)
type CacheConfig struct {
     ^
internal\cache\distributed.go:140:6: exported: type name will be used as cache.CacheShard by other packages, and that stutters; consider calling this Shard (revive)
type CacheShard struct {
     ^
internal\cache\distributed.go:158:6: exported: type name will be used as cache.CacheStats by other packages, and that stutters; consider calling this Stats (revive)
type CacheStats struct {
     ^
internal\cache\distributed.go:175:6: exported: type name will be used as cache.CacheEntry by other packages, and that stutters; consider calling this Entry (revive)
type CacheEntry struct {
     ^
internal\observability\enhanced_telemetry.go:67:2: field `spanMutex` is unused (unused)
        spanMutex   sync.RWMutex
        ^
internal\ratelimit\distributed.go:36:2: field `mu` is unused (unused)
        mu       sync.RWMutex
        ^
internal\events\nats_bus.go:9:2: import 'github.com/nats-io/nats.go' is not allowed from list 'main': Use pkg/natsx facade (depguard)
        "github.com/nats-io/nats.go"
        ^
internal\events\nats_bus.go:10:2: import 'go.uber.org/zap' is not allowed from list 'main': Use logger facade (depguard)
        "go.uber.org/zap"
        ^
internal\nats\publisher_error_handler.go:9:2: import 'github.com/nats-io/nats.go' is not allowed from list 'main': Use pkg/natsx facade (depguard)
        "github.com/nats-io/nats.go"
        ^
internal\ratelimit\distributed.go:10:2: import 'github.com/redis/go-redis/v9' is not allowed from list 'main': Use pkg/redisx facade (depguard)
        "github.com/redis/go-redis/v9"
        ^
internal\repository\postgres\task_repository.go:11:2: import 'github.com/google/uuid' is not allowed from list 'main': Use pkg/types (uuid re-exports) (depguard)
        "github.com/google/uuid"
        ^
internal\repository\redis\cache_repository.go:9:2: import 'github.com/redis/go-redis/v9' is not allowed from list 'main': Use pkg/redisx facade (depguard)
        "github.com/redis/go-redis/v9"
        ^
internal\repository\redis\connection.go:7:2: import 'github.com/redis/go-redis/v9' is not allowed from list 'main': Use pkg/redisx facade (depguard)
        "github.com/redis/go-redis/v9"
        ^
internal\slo\alerting.go:13:2: import 'go.uber.org/zap' is not allowed from list 'main': Use logger facade (depguard)
        "go.uber.org/zap"
        ^
internal\slo\monitor.go:12:2: import 'go.uber.org/zap' is not allowed from list 'main': Use logger facade (depguard)
        "go.uber.org/zap"
        ^
internal\tracing\business.go:10:2: import 'go.opentelemetry.io/otel' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel"
        ^
internal\tracing\business.go:11:2: import 'go.opentelemetry.io/otel/attribute' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/attribute"
        ^
internal\tracing\business.go:12:2: import 'go.opentelemetry.io/otel/baggage' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/baggage"
        ^
internal\tracing\business.go:13:2: import 'go.opentelemetry.io/otel/codes' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/codes"
        ^
internal\tracing\business.go:14:2: import 'go.opentelemetry.io/otel/trace' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/trace"
        ^
internal\ai\telemetry\metrics.go:7:2: import 'github.com/prometheus/client_golang/prometheus' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus"
        ^
internal\ai\telemetry\metrics.go:8:2: import 'github.com/prometheus/client_golang/prometheus/promauto' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus/promauto"
        ^
internal\ai\wiring\wiring.go:9:2: import 'github.com/prometheus/client_golang/prometheus' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus"
        ^
internal\ai\wiring\wiring_test.go:9:2: import 'github.com/prometheus/client_golang/prometheus' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus"
        ^
internal\config\tls_test.go:11:2: import 'go.uber.org/zap/zaptest' is not allowed from list 'main': Use logger facade (depguard)
        "go.uber.org/zap/zaptest"
        ^
internal\domain\models.go:6:2: import 'github.com/google/uuid' is not allowed from list 'main': Use pkg/types (uuid re-exports) (depguard)
        "github.com/google/uuid"
        ^
internal\domain\repository.go:6:2: import 'github.com/google/uuid' is not allowed from list 'main': Use pkg/types (uuid re-exports) (depguard)
        "github.com/google/uuid"
        ^
internal\observability\enhanced_telemetry.go:12:2: import 'github.com/prometheus/client_golang/prometheus/promauto' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus/promauto"
        ^
internal\observability\enhanced_telemetry.go:14:2: import 'go.opentelemetry.io/otel' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel"
        ^
internal\observability\enhanced_telemetry.go:15:2: import 'go.opentelemetry.io/otel/attribute' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/attribute"
        ^
internal\observability\enhanced_telemetry.go:16:2: import 'go.opentelemetry.io/otel/codes' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/codes"
        ^
internal\observability\enhanced_telemetry.go:17:2: import 'go.opentelemetry.io/otel/exporters/jaeger' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/exporters/jaeger"
        ^
internal\observability\enhanced_telemetry.go:18:2: import 'go.opentelemetry.io/otel/exporters/prometheus' is not allowed from list 'main': Use pkg/observability facade (depguard)
        promexporter "go.opentelemetry.io/otel/exporters/prometheus"
        ^
internal\observability\enhanced_telemetry.go:19:2: import 'go.opentelemetry.io/otel/metric' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/metric"
        ^
internal\observability\enhanced_telemetry.go:20:2: import 'go.opentelemetry.io/otel/propagation' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/propagation"
        ^
internal\observability\enhanced_telemetry.go:21:2: import 'go.opentelemetry.io/otel/sdk/metric' is not allowed from list 'main': Use pkg/observability facade (depguard)
        metricSDK "go.opentelemetry.io/otel/sdk/metric"
        ^
internal\observability\enhanced_telemetry.go:22:2: import 'go.opentelemetry.io/otel/sdk/resource' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/sdk/resource"
        ^
internal\observability\enhanced_telemetry.go:23:2: import 'go.opentelemetry.io/otel/sdk/trace' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/sdk/trace"
        ^
internal\observability\enhanced_telemetry.go:24:2: import 'go.opentelemetry.io/otel/semconv/v1.26.0' is not allowed from list 'main': Use pkg/observability facade (depguard)
        semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
        ^
internal\observability\enhanced_telemetry.go:25:2: import 'go.opentelemetry.io/otel/trace' is not allowed from list 'main': Use pkg/observability facade (depguard)
        oteltrace "go.opentelemetry.io/otel/trace"
        ^
internal\observability\integration.go:8:2: import 'go.opentelemetry.io/otel/attribute' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/attribute"
        ^
internal\observability\middleware.go:12:2: import 'go.opentelemetry.io/otel/codes' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/codes"
        ^
internal\observability\middleware.go:13:2: import 'go.opentelemetry.io/otel/trace' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/trace"
        ^
internal\observability\telemetry.go:10:2: import 'go.opentelemetry.io/otel' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel"
        ^
internal\observability\telemetry.go:13:2: import 'go.opentelemetry.io/otel/exporters/jaeger' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/exporters/jaeger"
        ^
internal\observability\telemetry.go:14:2: import 'go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
        ^
internal\observability\telemetry.go:15:2: import 'go.opentelemetry.io/otel/exporters/prometheus' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/exporters/prometheus"
        ^
internal\observability\telemetry.go:16:2: import 'go.opentelemetry.io/otel/metric' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/metric"
        ^
internal\observability\telemetry.go:17:2: import 'go.opentelemetry.io/otel/propagation' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/propagation"
        ^
internal\observability\telemetry.go:18:2: import 'go.opentelemetry.io/otel/sdk/metric' is not allowed from list 'main': Use pkg/observability facade (depguard)
        sdkmetric "go.opentelemetry.io/otel/sdk/metric"
        ^
internal\observability\telemetry.go:19:2: import 'go.opentelemetry.io/otel/sdk/resource' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/sdk/resource"
        ^
internal\observability\telemetry.go:20:2: import 'go.opentelemetry.io/otel/sdk/trace' is not allowed from list 'main': Use pkg/observability facade (depguard)
        sdktrace "go.opentelemetry.io/otel/sdk/trace"
        ^
internal\observability\telemetry.go:21:2: import 'go.opentelemetry.io/otel/semconv/v1.26.0' is not allowed from list 'main': Use pkg/observability facade (depguard)
        semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
        ^
internal\observability\telemetry_shim.go:7:2: import 'go.opentelemetry.io/otel/metric' is not allowed from list 'main': Use pkg/observability facade (depguard)
        "go.opentelemetry.io/otel/metric"
        ^
internal\observability\telemetry_test.go:16:2: import 'go.uber.org/zap/zaptest' is not allowed from list 'main': Use logger facade (depguard)
        "go.uber.org/zap/zaptest"
        ^
internal\telemetry\metrics.go:9:2: import 'github.com/prometheus/client_golang/prometheus/promauto' is not allowed from list 'main': Use pkg/metrics facade (depguard)
        "github.com/prometheus/client_golang/prometheus/promauto"
        ^
internal\lifecycle\deployment.go:563:20: S1039: unnecessary use of fmt.Sprintf (gosimple)
        da.addLog(result, fmt.Sprintf("Script executed successfully"))
                          ^
internal\observability\telemetry_test.go:328:2: ineffectual assignment to ctx (ineffassign)
        ctx, span := tracer.Start(ctx, "test-operation",
        ^
automation\autocommit.go:7:2: SA1019: "io/ioutil" has been deprecated since Go 1.19: As of Go 1.16, the same functionality is now provided by package [io] or package [os], and those implementations should be preferred in new code. See the specific function documentation for details. (staticcheck)
        "io/ioutil"
        ^
internal\telemetry\tracing.go:187:10: SA1019: trace.NewNoopTracerProvider is deprecated: Use [go.opentelemetry.io/otel/trace/noop.NewTracerProvider] instead. (staticcheck)
                return trace.NewNoopTracerProvider().Tracer(name)
                       ^
basic_test.go:18:5: SA4000: identical expressions on the left and right side of the '!=' operator (staticcheck)
        if true != true {
           ^
make: *** [Makefile:4: lint] Error 1
PS E:\vertikon\business\SaaS\templates\mcp-ultra> make test
go test ./... -count=1
# github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]
internal\compliance\framework_stubs.go:12:6: DataAccessRequest redeclared in this block
        internal\compliance\framework.go:544:6: other declaration of DataAccessRequest
internal\compliance\framework_stubs.go:17:6: DataDeletionRequest redeclared in this block
        internal\compliance\framework.go:554:6: other declaration of DataDeletionRequest
internal\compliance\framework_stubs.go:22:6: AuditEvent redeclared in this block
        internal\compliance\audit_logger.go:27:6: other declaration of AuditEvent
internal\compliance\framework_stubs.go:29:31: method ComplianceFramework.ProcessDataAccessRequest already declared at internal\compliance\framework.go:583:32
internal\compliance\framework_stubs.go:34:31: method ComplianceFramework.AnonymizeData already declared at internal\compliance\framework.go:623:32
internal\compliance\framework_stubs.go:39:31: method ComplianceFramework.LogAuditEvent already declared at internal\compliance\framework.go:640:32
internal\compliance\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value
internal\compliance\framework_test.go:208:17: too many errors
ok      github.com/vertikon/mcp-ultra   0.763s
?       github.com/vertikon/mcp-ultra/api/grpc/gen/compliance/v1        [no test files]
?       github.com/vertikon/mcp-ultra/api/grpc/gen/system/v1    [no test files]
?       github.com/vertikon/mcp-ultra/api/grpc/gen/task/v1      [no test files]
?       github.com/vertikon/mcp-ultra/automation        [no test files]
ok      github.com/vertikon/mcp-ultra/internal/ai/events        0.405s
?       github.com/vertikon/mcp-ultra/internal/ai/router        [no test files]
ok      github.com/vertikon/mcp-ultra/internal/ai/telemetry     0.720s
ok      github.com/vertikon/mcp-ultra/internal/ai/wiring        0.510s
# github.com/vertikon/mcp-ultra/internal/handlers/http [github.com/vertikon/mcp-ultra/internal/handlers/http.test]
internal\handlers\http\router_test.go:23:76: undefined: services.HealthStatus
internal\handlers\http\router_test.go:25:42: undefined: services.HealthStatus
internal\handlers\http\router_test.go:38:75: undefined: services.HealthChecker
internal\handlers\http\router_test.go:47:70: undefined: domain.CreateTaskRequest
internal\handlers\http\router_test.go:60:85: undefined: domain.UpdateTaskRequest
internal\handlers\http\router_test.go:70:73: undefined: domain.TaskFilters
internal\handlers\http\router_test.go:70:95: undefined: domain.TaskList
internal\handlers\http\router_test.go:72:30: undefined: domain.TaskList
internal\handlers\http\router_test.go:80:49: not enough arguments in call to NewRouter
        have (*zap.Logger, *MockHealthService, *MockTaskService)
        want (*services.TaskService, *features.FlagManager, *HealthService, *zap.Logger)
internal\handlers\http\router_test.go:101:77: undefined: services.HealthStatus
internal\handlers\http\router_test.go:101:77: too many errors
# github.com/vertikon/mcp-ultra/internal/middleware [github.com/vertikon/mcp-ultra/internal/middleware.test]
internal\middleware\auth_test.go:96:30: undefined: testhelpers.GetTestAPIKeys
--- FAIL: TestCircuitBreaker_HalfOpenMaxRequests (0.06s)
    circuit_breaker_test.go:260:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/circuit_breaker_test.go:260
                Error:          Should be false
                Test:           TestCircuitBreaker_HalfOpenMaxRequests
                Messages:       Request should be denied after max half-open requests
--- FAIL: TestDistributedCache_SetAndGet (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:69
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetAndGet
--- FAIL: TestDistributedCache_SetWithTTL (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:88
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetWithTTL
--- FAIL: TestDistributedCache_Delete (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:116
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Delete
--- FAIL: TestDistributedCache_Clear (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:144
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Clear
--- FAIL: TestDistributedCache_GetNonExistentKey (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:169
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_GetNonExistentKey
--- FAIL: TestDistributedCache_SetComplexObject (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:181
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetComplexObject
--- FAIL: TestDistributedCache_ConcurrentOperations (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:232
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_ConcurrentOperations
--- FAIL: TestDistributedCache_Namespace (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:268
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Namespace
--- FAIL: TestCacheStrategy_WriteThrough (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:297
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestCacheStrategy_WriteThrough
--- FAIL: TestDistributedCache_InvalidKey (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:316
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_InvalidKey
FAIL
FAIL    github.com/vertikon/mcp-ultra/internal/cache    0.829s
FAIL    github.com/vertikon/mcp-ultra/internal/compliance [build failed]
--- FAIL: TestNewTLSManager (0.05s)
    logger.go:146: 2025-10-17T16:17:31.547-0300 INFO    TLS is disabled
    --- FAIL: TestNewTLSManager/should_create_manager_with_valid_TLS_config (0.01s)
        tls_test.go:120:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:120
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestNewTLSManager/should_create_manager_with_valid_TLS_config
--- FAIL: TestTLSManager_GetTLSConfig (0.02s)
    --- FAIL: TestTLSManager_GetTLSConfig/should_return_copy_of_TLS_config (0.02s)
        tls_test.go:306:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:306
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestTLSManager_GetTLSConfig/should_return_copy_of_TLS_config
--- FAIL: TestTLSManager_Stop (0.01s)
    --- FAIL: TestTLSManager_Stop/should_stop_certificate_watcher (0.01s)
        tls_test.go:334:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:334
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestTLSManager_Stop/should_stop_certificate_watcher
FAIL
FAIL    github.com/vertikon/mcp-ultra/internal/config   0.569s
?       github.com/vertikon/mcp-ultra/internal/config/secrets   [no test files]
?       github.com/vertikon/mcp-ultra/internal/constants        [no test files]
?       github.com/vertikon/mcp-ultra/internal/dashboard        [no test files]
--- FAIL: TestTaskComplete (0.00s)
    models_test.go:40:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:40
                Error:          Should be true
                Test:           TestTaskComplete
    models_test.go:41:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:41
                Error:          Should be true
                Test:           TestTaskComplete
    models_test.go:42:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:42
                Error:          Should be true
                Test:           TestTaskComplete
--- FAIL: TestTaskCancel (0.00s)
    models_test.go:53:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:53
                Error:          Should be true
                Test:           TestTaskCancel
    models_test.go:54:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:54
                Error:          Should be true
                Test:           TestTaskCancel
--- FAIL: TestTaskUpdateStatus (0.00s)
    models_test.go:65:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:65
                Error:          Should be true
                Test:           TestTaskUpdateStatus
    models_test.go:66:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:66
                Error:          Should be true
                Test:           TestTaskUpdateStatus
FAIL
FAIL    github.com/vertikon/mcp-ultra/internal/domain   0.413s
?       github.com/vertikon/mcp-ultra/internal/dr       [no test files]
?       github.com/vertikon/mcp-ultra/internal/events   [no test files]
# github.com/vertikon/mcp-ultra/internal/security [github.com/vertikon/mcp-ultra/internal/security.test]
internal\security\enhanced_auth_test.go:22:6: MockOPAService redeclared in this block
        internal\security\auth_test.go:20:6: other declaration of MockOPAService
internal\security\enhanced_auth_test.go:26:26: method MockOPAService.IsAuthorized already declared at internal\security\auth_test.go:24:26
internal\security\enhanced_auth_test.go:36:6: TestNewAuthService redeclared in this block
        internal\security\auth_test.go:39:6: other declaration of TestNewAuthService
internal\security\enhanced_auth_test.go:326:6: TestGetUserFromContext redeclared in this block
        internal\security\auth_test.go:411:6: other declaration of TestGetUserFromContext
internal\security\enhanced_auth_test.go:391:6: TestRequireScope redeclared in this block
        internal\security\auth_test.go:282:6: other declaration of TestRequireScope
internal\security\enhanced_auth_test.go:459:6: TestRequireRole redeclared in this block
        internal\security\auth_test.go:342:6: other declaration of TestRequireRole
internal\security\auth_test.go:49:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:67:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:140:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:163:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:163:48: too many errors
# github.com/vertikon/mcp-ultra/internal/services [github.com/vertikon/mcp-ultra/internal/services.test]
internal\services\task_service_test.go:104:70: undefined: domain.UserFilter
internal\services\task_service_test.go:171:28: cannot use taskRepo (variable of type *mockTaskRepository) as domain.TaskRepository value in argument to NewTaskService: *mockTaskRepository does not implement domain.TaskRepository (wrong type for method List)
                have List(context.Context, domain.TaskFilter) ([]*domain.Task, error)
                want List(context.Context, domain.TaskFilter) ([]*domain.Task, int, error)
internal\services\task_service_test.go:171:48: cannot use eventRepo (variable of type *mockEventRepository) as domain.EventRepository value in argument to NewTaskService: *mockEventRepository does not implement domain.EventRepository (missing method GetByType)
internal\services\task_service_test.go:171:59: cannot use cacheRepo (variable of type *mockCacheRepository) as domain.CacheRepository value in argument to NewTaskService: *mockCacheRepository does not implement domain.CacheRepository (missing method Exists)
internal\services\task_service_test.go:199:31: declared and not used: eventRepo
ok      github.com/vertikon/mcp-ultra/internal/features 0.422s
ok      github.com/vertikon/mcp-ultra/internal/handlers 0.442s
FAIL    github.com/vertikon/mcp-ultra/internal/handlers/http [build failed]
?       github.com/vertikon/mcp-ultra/internal/http     [no test files]
?       github.com/vertikon/mcp-ultra/internal/lifecycle        [no test files]
?       github.com/vertikon/mcp-ultra/internal/metrics  [no test files]
FAIL    github.com/vertikon/mcp-ultra/internal/middleware [build failed]
?       github.com/vertikon/mcp-ultra/internal/nats     [no test files]
# github.com/vertikon/mcp-ultra/test/component [github.com/vertikon/mcp-ultra/test/component.test]
test\component\task_service_test.go:39:3: cannot use suite.taskRepo (variable of type *mocks.MockTaskRepository) as domain.TaskRepository value in argument to services.NewTaskService: *mocks.MockTaskRepository does not implement domain.TaskRepository (wrong type for method Delete)
                have Delete(context.Context, string) error
                want Delete(context.Context, uuid.UUID) error
test\component\task_service_test.go:40:3: cannot use suite.validator (variable of type *mocks.MockValidator) as domain.UserRepository value in argument to services.NewTaskService: *mocks.MockValidator does not implement domain.UserRepository (missing method Create)
test\component\task_service_test.go:42:3: cannot use suite.cacheRepo (variable of type *mocks.MockCacheRepository) as domain.CacheRepository value in argument to services.NewTaskService: *mocks.MockCacheRepository does not implement domain.CacheRepository (wrong type for method Get)
                have Get(context.Context, string) (interface{}, error)
                want Get(context.Context, string) (string, error)
test\component\task_service_test.go:44:3: cannot use suite.eventBus (variable of type *mocks.MockEventBus) as services.EventBus value in argument to services.NewTaskService: *mocks.MockEventBus does not implement services.EventBus (wrong type for method Publish)
                have Publish(context.Context, string, []byte) error
                want Publish(context.Context, *domain.Event) error
test\component\task_service_test.go:65:3: unknown field Metadata in struct literal of type services.CreateTaskRequest
test\component\task_service_test.go:78:20: req.Metadata undefined (type *services.CreateTaskRequest has no field or method Metadata)
test\component\task_service_test.go:97:55: too many arguments in call to suite.service.CreateTask
        have (context.Context, uuid.UUID, *services.CreateTaskRequest)
        want (context.Context, services.CreateTaskRequest)
test\component\task_service_test.go:118:29: undefined: services.ValidationError
test\component\task_service_test.go:127:55: too many arguments in call to suite.service.CreateTask
        have (context.Context, uuid.UUID, *services.CreateTaskRequest)
        want (context.Context, services.CreateTaskRequest)
test\component\task_service_test.go:151:52: too many arguments in call to suite.service.GetTask
        have (context.Context, uuid.UUID, uuid.UUID)
        want (context.Context, uuid.UUID)
test\component\task_service_test.go:151:52: too many errors
# github.com/vertikon/mcp-ultra/test/property [github.com/vertikon/mcp-ultra/test/property.test]
test\property\task_properties_test.go:231:4: declared and not used: originalTitle
--- FAIL: TestTelemetryService_Tracing (0.00s)
    logger.go:146: 2025-10-17T16:17:33.832-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:33.832-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T16:17:33.832-0300 DEBUG   TelemetryService.Start called (initialization already complete)
    telemetry_test.go:92:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:92
                Error:          Should be true
                Test:           TestTelemetryService_Tracing
    telemetry_test.go:93:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:93
                Error:          Should not be: trace.SpanID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
                Test:           TestTelemetryService_Tracing
    telemetry_test.go:94:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:94
                Error:          Should not be: trace.TraceID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
                Test:           TestTelemetryService_Tracing
    logger.go:146: 2025-10-17T16:17:33.832-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryService_SpanAttributes (0.00s)
    logger.go:146: 2025-10-17T16:17:33.845-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:33.845-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T16:17:33.845-0300 DEBUG   TelemetryService.Start called (initialization already complete)
    telemetry_test.go:345:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:345
                Error:          Should be true
                Test:           TestTelemetryService_SpanAttributes
    logger.go:146: 2025-10-17T16:17:33.845-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryConfig_Validation (0.00s)
    logger.go:146: 2025-10-17T16:17:33.845-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:33.845-0300 INFO    Telemetry initialized successfully      {"service": "test", "version": "", "environment": ""}
    logger.go:146: 2025-10-17T16:17:33.845-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:33.845-0300 INFO    Telemetry initialized successfully      {"service": "", "version": "", "environment": ""}
    telemetry_test.go:376:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:376
                Error:          Should NOT be empty, but was
                Test:           TestTelemetryConfig_Validation
FAIL
FAIL    github.com/vertikon/mcp-ultra/internal/observability    0.502s
?       github.com/vertikon/mcp-ultra/internal/ratelimit        [no test files]
?       github.com/vertikon/mcp-ultra/internal/repository/postgres      [no test files]
?       github.com/vertikon/mcp-ultra/internal/repository/redis [no test files]
FAIL    github.com/vertikon/mcp-ultra/internal/security [build failed]
FAIL    github.com/vertikon/mcp-ultra/internal/services [build failed]
?       github.com/vertikon/mcp-ultra/internal/slo      [no test files]
--- FAIL: TestNewTracingProvider (0.02s)
    --- FAIL: TestNewTracingProvider/should_create_provider_with_stdout_exporter (0.00s)
        tracing_test.go:29:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:29
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_create_provider_with_stdout_exporter
    --- FAIL: TestNewTracingProvider/should_create_provider_with_noop_exporter (0.00s)
        tracing_test.go:49:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:49
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_create_provider_with_noop_exporter
    logger.go:146: 2025-10-17T16:17:34.043-0300 INFO    Distributed tracing is disabled
    logger.go:146: 2025-10-17T16:17:34.065-0300 INFO    Shutting down tracing provider
    --- FAIL: TestNewTracingProvider/should_include_custom_resource_attributes (0.00s)
        tracing_test.go:82:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:82
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_include_custom_resource_attributes
--- FAIL: TestTracingProvider_GetTracer (0.00s)
    --- FAIL: TestTracingProvider_GetTracer/should_return_tracer_when_enabled (0.00s)
        tracing_test.go:98:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:98
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTracingProvider_GetTracer/should_return_tracer_when_enabled
    logger.go:146: 2025-10-17T16:17:34.066-0300 INFO    Distributed tracing is disabled
--- FAIL: TestTraceFunction (0.00s)
    tracing_test.go:128:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:128
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceFunction
--- FAIL: TestTraceFunctionWithResult (0.00s)
    tracing_test.go:163:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:163
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceFunctionWithResult
--- FAIL: TestSpanUtilities (0.00s)
    tracing_test.go:198:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:198
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestSpanUtilities
--- FAIL: TestTraceContextPropagation (0.00s)
    tracing_test.go:275:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:275
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceContextPropagation
FAIL
FAIL    github.com/vertikon/mcp-ultra/internal/telemetry        0.518s
?       github.com/vertikon/mcp-ultra/internal/testhelpers      [no test files]
?       github.com/vertikon/mcp-ultra/internal/tracing  [no test files]
?       github.com/vertikon/mcp-ultra/scripts   [no test files]
FAIL    github.com/vertikon/mcp-ultra/test/component [build failed]
?       github.com/vertikon/mcp-ultra/test/mocks        [no test files]
FAIL    github.com/vertikon/mcp-ultra/test/property [build failed]
ok      github.com/vertikon/mcp-ultra/tests/smoke       0.357s
FAIL
make: *** [Makefile:7: test] Error 1
PS E:\vertikon\business\SaaS\templates\mcp-ultra> make coverage-html
go test ./... -coverprofile=coverage.out
ok      github.com/vertikon/mcp-ultra   0.606s  coverage: 0.0% of statements
# github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]
internal\compliance\framework_stubs.go:12:6: DataAccessRequest redeclared in this block
        internal\compliance\framework.go:544:6: other declaration of DataAccessRequest
internal\compliance\framework_stubs.go:17:6: DataDeletionRequest redeclared in this block
        internal\compliance\framework.go:554:6: other declaration of DataDeletionRequest
internal\compliance\framework_stubs.go:22:6: AuditEvent redeclared in this block
        internal\compliance\audit_logger.go:27:6: other declaration of AuditEvent
internal\compliance\framework_stubs.go:29:31: method ComplianceFramework.ProcessDataAccessRequest already declared at internal\compliance\framework.go:583:32
internal\compliance\framework_stubs.go:34:31: method ComplianceFramework.AnonymizeData already declared at internal\compliance\framework.go:623:32
internal\compliance\framework_stubs.go:39:31: method ComplianceFramework.LogAuditEvent already declared at internal\compliance\framework.go:640:32
internal\compliance\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value
internal\compliance\framework_test.go:208:17: too many errors
        github.com/vertikon/mcp-ultra/api/grpc/gen/compliance/v1                coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/api/grpc/gen/system/v1            coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/api/grpc/gen/task/v1              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/automation                coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/ai/events        0.590s  coverage: 100.0% of statements
        github.com/vertikon/mcp-ultra/internal/ai/router                coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/ai/telemetry     0.619s  coverage: 87.9% of statements
ok      github.com/vertikon/mcp-ultra/internal/ai/wiring        0.776s  coverage: 80.0% of statements
--- FAIL: TestCircuitBreaker_HalfOpenMaxRequests (0.06s)
    circuit_breaker_test.go:260:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/circuit_breaker_test.go:260
                Error:          Should be false
                Test:           TestCircuitBreaker_HalfOpenMaxRequests
                Messages:       Request should be denied after max half-open requests
--- FAIL: TestDistributedCache_SetAndGet (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:69
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetAndGet
--- FAIL: TestDistributedCache_SetWithTTL (0.05s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:88
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetWithTTL
--- FAIL: TestDistributedCache_Delete (0.04s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:116
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Delete
--- FAIL: TestDistributedCache_Clear (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:144
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Clear
--- FAIL: TestDistributedCache_GetNonExistentKey (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:169
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_GetNonExistentKey
--- FAIL: TestDistributedCache_SetComplexObject (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:181
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetComplexObject
--- FAIL: TestDistributedCache_ConcurrentOperations (0.02s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:232
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_ConcurrentOperations
--- FAIL: TestDistributedCache_Namespace (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:268
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Namespace
--- FAIL: TestCacheStrategy_WriteThrough (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:297
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestCacheStrategy_WriteThrough
--- FAIL: TestDistributedCache_InvalidKey (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:316
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_InvalidKey
FAIL
coverage: 17.7% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/cache    1.273s
FAIL    github.com/vertikon/mcp-ultra/internal/compliance [build failed]
--- FAIL: TestNewTLSManager (0.08s)
    logger.go:146: 2025-10-17T16:17:42.306-0300 INFO    TLS is disabled
    --- FAIL: TestNewTLSManager/should_create_manager_with_valid_TLS_config (0.02s)
        tls_test.go:120:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:120
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestNewTLSManager/should_create_manager_with_valid_TLS_config
--- FAIL: TestTLSManager_GetTLSConfig (0.02s)
    --- FAIL: TestTLSManager_GetTLSConfig/should_return_copy_of_TLS_config (0.01s)
        tls_test.go:306:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:306
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestTLSManager_GetTLSConfig/should_return_copy_of_TLS_config
--- FAIL: TestTLSManager_Stop (0.03s)
    --- FAIL: TestTLSManager_Stop/should_stop_certificate_watcher (0.03s)
        tls_test.go:334:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:334
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestTLSManager_Stop/should_stop_certificate_watcher
FAIL
coverage: 39.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/config   0.827s
# github.com/vertikon/mcp-ultra/internal/handlers/http [github.com/vertikon/mcp-ultra/internal/handlers/http.test]
internal\handlers\http\router_test.go:23:76: undefined: services.HealthStatus
internal\handlers\http\router_test.go:25:42: undefined: services.HealthStatus
internal\handlers\http\router_test.go:38:75: undefined: services.HealthChecker
internal\handlers\http\router_test.go:47:70: undefined: domain.CreateTaskRequest
internal\handlers\http\router_test.go:60:85: undefined: domain.UpdateTaskRequest
internal\handlers\http\router_test.go:70:73: undefined: domain.TaskFilters
internal\handlers\http\router_test.go:70:95: undefined: domain.TaskList
internal\handlers\http\router_test.go:72:30: undefined: domain.TaskList
internal\handlers\http\router_test.go:80:49: not enough arguments in call to NewRouter
        have (*zap.Logger, *MockHealthService, *MockTaskService)
        want (*services.TaskService, *features.FlagManager, *HealthService, *zap.Logger)
internal\handlers\http\router_test.go:101:77: undefined: services.HealthStatus
internal\handlers\http\router_test.go:101:77: too many errors
# github.com/vertikon/mcp-ultra/internal/middleware [github.com/vertikon/mcp-ultra/internal/middleware.test]
internal\middleware\auth_test.go:96:30: undefined: testhelpers.GetTestAPIKeys
        github.com/vertikon/mcp-ultra/internal/config/secrets           coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/constants                coverage: 0.0% of statements
?       github.com/vertikon/mcp-ultra/internal/dashboard        [no test files]
--- FAIL: TestTaskComplete (0.00s)
    models_test.go:40:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:40
                Error:          Should be true
                Test:           TestTaskComplete
    models_test.go:41:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:41
                Error:          Should be true
                Test:           TestTaskComplete
    models_test.go:42:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:42
                Error:          Should be true
                Test:           TestTaskComplete
--- FAIL: TestTaskCancel (0.00s)
    models_test.go:53:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:53
                Error:          Should be true
                Test:           TestTaskCancel
    models_test.go:54:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:54
                Error:          Should be true
                Test:           TestTaskCancel
--- FAIL: TestTaskUpdateStatus (0.00s)
    models_test.go:65:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:65
                Error:          Should be true
                Test:           TestTaskUpdateStatus
    models_test.go:66:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/domain/models_test.go:66
                Error:          Should be true
                Test:           TestTaskUpdateStatus
FAIL
coverage: 92.9% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/domain   0.657s
        github.com/vertikon/mcp-ultra/internal/dr               coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/events           coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/features 0.631s  coverage: 22.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/handlers 0.613s  coverage: 100.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/handlers/http [build failed]
        github.com/vertikon/mcp-ultra/internal/http             coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/lifecycle                coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/metrics          coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/middleware [build failed]
# github.com/vertikon/mcp-ultra/internal/services [github.com/vertikon/mcp-ultra/internal/services.test]
internal\services\task_service_test.go:104:70: undefined: domain.UserFilter
internal\services\task_service_test.go:171:28: cannot use taskRepo (variable of type *mockTaskRepository) as domain.TaskRepository value in argument to NewTaskService: *mockTaskRepository does not implement domain.TaskRepository (wrong type for method List)
                have List(context.Context, domain.TaskFilter) ([]*domain.Task, error)
                want List(context.Context, domain.TaskFilter) ([]*domain.Task, int, error)
internal\services\task_service_test.go:171:48: cannot use eventRepo (variable of type *mockEventRepository) as domain.EventRepository value in argument to NewTaskService: *mockEventRepository does not implement domain.EventRepository (missing method GetByType)
internal\services\task_service_test.go:171:59: cannot use cacheRepo (variable of type *mockCacheRepository) as domain.CacheRepository value in argument to NewTaskService: *mockCacheRepository does not implement domain.CacheRepository (missing method Exists)
internal\services\task_service_test.go:199:31: declared and not used: eventRepo
        github.com/vertikon/mcp-ultra/internal/nats             coverage: 0.0% of statements
# github.com/vertikon/mcp-ultra/internal/security [github.com/vertikon/mcp-ultra/internal/security.test]
internal\security\enhanced_auth_test.go:22:6: MockOPAService redeclared in this block
        internal\security\auth_test.go:20:6: other declaration of MockOPAService
internal\security\enhanced_auth_test.go:26:26: method MockOPAService.IsAuthorized already declared at internal\security\auth_test.go:24:26
internal\security\enhanced_auth_test.go:36:6: TestNewAuthService redeclared in this block
        internal\security\auth_test.go:39:6: other declaration of TestNewAuthService
internal\security\enhanced_auth_test.go:326:6: TestGetUserFromContext redeclared in this block
        internal\security\auth_test.go:411:6: other declaration of TestGetUserFromContext
internal\security\enhanced_auth_test.go:391:6: TestRequireScope redeclared in this block
        internal\security\auth_test.go:282:6: other declaration of TestRequireScope
internal\security\enhanced_auth_test.go:459:6: TestRequireRole redeclared in this block
        internal\security\auth_test.go:342:6: other declaration of TestRequireRole
internal\security\auth_test.go:49:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:67:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:140:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:163:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService
internal\security\auth_test.go:163:48: too many errors
--- FAIL: TestTelemetryService_Tracing (0.00s)
    logger.go:146: 2025-10-17T16:17:48.222-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:48.223-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T16:17:48.223-0300 DEBUG   TelemetryService.Start called (initialization already complete)
    telemetry_test.go:92:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:92
                Error:          Should be true
                Test:           TestTelemetryService_Tracing
    telemetry_test.go:93:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:93
                Error:          Should not be: trace.SpanID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
                Test:           TestTelemetryService_Tracing
    telemetry_test.go:94:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:94
                Error:          Should not be: trace.TraceID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
                Test:           TestTelemetryService_Tracing
    logger.go:146: 2025-10-17T16:17:48.223-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryService_SpanAttributes (0.00s)
    logger.go:146: 2025-10-17T16:17:48.239-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:48.239-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T16:17:48.239-0300 DEBUG   TelemetryService.Start called (initialization already complete)
    telemetry_test.go:345:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:345
                Error:          Should be true
                Test:           TestTelemetryService_SpanAttributes
    logger.go:146: 2025-10-17T16:17:48.239-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryConfig_Validation (0.00s)
    logger.go:146: 2025-10-17T16:17:48.239-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:48.239-0300 INFO    Telemetry initialized successfully      {"service": "test", "version": "", "environment": ""}
    logger.go:146: 2025-10-17T16:17:48.239-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T16:17:48.239-0300 INFO    Telemetry initialized successfully      {"service": "", "version": "", "environment": ""}
    telemetry_test.go:376:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:376
                Error:          Should NOT be empty, but was
                Test:           TestTelemetryConfig_Validation
FAIL
coverage: 22.1% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/observability    0.809s
# github.com/vertikon/mcp-ultra/test/component [github.com/vertikon/mcp-ultra/test/component.test]
test\component\task_service_test.go:39:3: cannot use suite.taskRepo (variable of type *mocks.MockTaskRepository) as domain.TaskRepository value in argument to services.NewTaskService: *mocks.MockTaskRepository does not implement domain.TaskRepository (wrong type for method Delete)
                have Delete(context.Context, string) error
                want Delete(context.Context, uuid.UUID) error
test\component\task_service_test.go:40:3: cannot use suite.validator (variable of type *mocks.MockValidator) as domain.UserRepository value in argument to services.NewTaskService: *mocks.MockValidator does not implement domain.UserRepository (missing method Create)
test\component\task_service_test.go:42:3: cannot use suite.cacheRepo (variable of type *mocks.MockCacheRepository) as domain.CacheRepository value in argument to services.NewTaskService: *mocks.MockCacheRepository does not implement domain.CacheRepository (wrong type for method Get)
                have Get(context.Context, string) (interface{}, error)
                want Get(context.Context, string) (string, error)
test\component\task_service_test.go:44:3: cannot use suite.eventBus (variable of type *mocks.MockEventBus) as services.EventBus value in argument to services.NewTaskService: *mocks.MockEventBus does not implement services.EventBus (wrong type for method Publish)
                have Publish(context.Context, string, []byte) error
                want Publish(context.Context, *domain.Event) error
test\component\task_service_test.go:65:3: unknown field Metadata in struct literal of type services.CreateTaskRequest
test\component\task_service_test.go:78:20: req.Metadata undefined (type *services.CreateTaskRequest has no field or method Metadata)
test\component\task_service_test.go:97:55: too many arguments in call to suite.service.CreateTask
        have (context.Context, uuid.UUID, *services.CreateTaskRequest)
        want (context.Context, services.CreateTaskRequest)
test\component\task_service_test.go:118:29: undefined: services.ValidationError
test\component\task_service_test.go:127:55: too many arguments in call to suite.service.CreateTask
        have (context.Context, uuid.UUID, *services.CreateTaskRequest)
        want (context.Context, services.CreateTaskRequest)
test\component\task_service_test.go:151:52: too many arguments in call to suite.service.GetTask
        have (context.Context, uuid.UUID, uuid.UUID)
        want (context.Context, uuid.UUID)
test\component\task_service_test.go:151:52: too many errors
        github.com/vertikon/mcp-ultra/internal/ratelimit                coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/repository/postgres              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/repository/redis         coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/security [build failed]
FAIL    github.com/vertikon/mcp-ultra/internal/services [build failed]
        github.com/vertikon/mcp-ultra/internal/slo              coverage: 0.0% of statements
--- FAIL: TestNewTracingProvider (0.05s)
    --- FAIL: TestNewTracingProvider/should_create_provider_with_stdout_exporter (0.00s)
        tracing_test.go:29:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:29
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_create_provider_with_stdout_exporter
    --- FAIL: TestNewTracingProvider/should_create_provider_with_noop_exporter (0.00s)
        tracing_test.go:49:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:49
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_create_provider_with_noop_exporter
    logger.go:146: 2025-10-17T16:17:50.299-0300 INFO    Distributed tracing is disabled
    logger.go:146: 2025-10-17T16:17:50.346-0300 INFO    Shutting down tracing provider
    --- FAIL: TestNewTracingProvider/should_include_custom_resource_attributes (0.00s)
        tracing_test.go:82:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:82
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestNewTracingProvider/should_include_custom_resource_attributes
--- FAIL: TestTracingProvider_GetTracer (0.00s)
    --- FAIL: TestTracingProvider_GetTracer/should_return_tracer_when_enabled (0.00s)
        tracing_test.go:98:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:98
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTracingProvider_GetTracer/should_return_tracer_when_enabled
    logger.go:146: 2025-10-17T16:17:50.346-0300 INFO    Distributed tracing is disabled
--- FAIL: TestTraceFunction (0.00s)
    tracing_test.go:128:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:128
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceFunction
--- FAIL: TestTraceFunctionWithResult (0.00s)
    tracing_test.go:163:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:163
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceFunctionWithResult
--- FAIL: TestSpanUtilities (0.00s)
    tracing_test.go:198:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:198
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestSpanUtilities
--- FAIL: TestTraceContextPropagation (0.00s)
    tracing_test.go:275:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/telemetry/tracing_test.go:275
                Error:          Received unexpected error:
                                failed to create tracing resource: conflicting Schema URL: https://opentelemetry.io/schemas/1.37.0 and https://opentelemetry.io/schemas/1.26.0
                Test:           TestTraceContextPropagation
FAIL
coverage: 18.3% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/telemetry        0.724s
# github.com/vertikon/mcp-ultra/test/property [github.com/vertikon/mcp-ultra/test/property.test]
test\property\task_properties_test.go:231:4: declared and not used: originalTitle
        github.com/vertikon/mcp-ultra/internal/testhelpers              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/tracing          coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/scripts           coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/test/component [build failed]
        github.com/vertikon/mcp-ultra/test/mocks                coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/test/property [build failed]
ok      github.com/vertikon/mcp-ultra/tests/smoke       0.500s  coverage: [no statements]
FAIL
make: *** [Makefile:10: coverage-html] Error 1
PS E:\vertikon\business\SaaS\templates\mcp-ultra>