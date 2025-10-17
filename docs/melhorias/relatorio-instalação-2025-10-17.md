Instale o PowerShell mais recente para obter novos recursos e aprimoramentos! https://aka.ms/PSWindows

✅ GPT5 Integration carregado
🚀 Carregando profile Vertikon...
  ✓ Go bin adicionado ao PATH
✅ Profile Vertikon carregado!
   Root: E:\vertikon
   Digite 'aliases' para ver comandos disponíveis
   Digite 'Check-GoTools' para verificar ferramentas

PS E:\vertikon\business\SaaS\templates\mcp-ultra> go test ./internal/compliance/... -count=1
# github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]
internal\compliance\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value
internal\compliance\framework_test.go:230:3: unknown field UserID in struct literal of type DataDeletionRequest
internal\compliance\framework_test.go:231:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:236:17: assignment mismatch: 2 variables but framework.ProcessDataDeletionRequest returns 1 value
internal\compliance\framework_test.go:254:25: assignment mismatch: 2 variables but framework.AnonymizeData returns 1 value
internal\compliance\framework_test.go:254:68: too many arguments in call to framework.AnonymizeData
        have (context.Context, map[string]interface{}, string)
        want (context.Context, string)
internal\compliance\framework_test.go:279:46: too many arguments in call to framework.LogAuditEvent
        have (context.Context, uuid.UUID, string, map[string]interface{})
        want (context.Context, AuditEvent)
internal\compliance\framework_test.go:279:46: too many errors
FAIL    github.com/vertikon/mcp-ultra/internal/compliance [build failed]
FAIL
PS E:\vertikon\business\SaaS\templates\mcp-ultra> golangci-lint run internal/compliance/...
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]\ninternal\\compliance\\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal\ninternal\\compliance\\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value\ninternal\\compliance\\framework_test.go:230:3: unknown field UserID in struct literal of type DataDeletionRequest\ninternal\\compliance\\framework_test.go:231:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal\ninternal\\compliance\\framework_test.go:236:17: assignment mismatch: 2 variables but framework.ProcessDataDeletionRequest returns 1 value\ninternal\\compliance\\framework_test.go:254:25: assignment mismatch: 2 variables but framework.AnonymizeData returns 1 value\ninternal\\compliance\\framework_test.go:254:68: too many arguments in call to framework.AnonymizeData\n\thave (context.Context, map[string]interface{}, string)\n\twant (context.Context, string)\ninternal\\compliance\\framework_test.go:279:46: too many arguments in call to framework.LogAuditEvent\n\thave (context.Context, uuid.UUID, string, map[string]interface{})\n\twant (context.Context, AuditEvent)\ninternal\\compliance\\framework_test.go:279:46: too many errors"
PS E:\vertikon\business\SaaS\templates\mcp-ultra> gofmt -w internal/compliance/framework.go
PS E:\vertikon\business\SaaS\templates\mcp-ultra> goimports -w internal/compliance/framework.go
PS E:\vertikon\business\SaaS\templates\mcp-ultra> gofmt -s -l .
PS E:\vertikon\business\SaaS\templates\mcp-ultra> # deve retornar vazio → [11/20] Formatação ✅ PASS
PS E:\vertikon\business\SaaS\templates\mcp-ultra> # 1. Corrigir erros não tratados (manual)        S
PS E:\vertikon\business\SaaS\templates\mcp-ultra> # 2. Aplicar formatação segura
PS E:\vertikon\business\SaaS\templates\mcp-ultra> gofmt -w .
PS E:\vertikon\business\SaaS\templates\mcp-ultra> goimports -w .
PS E:\vertikon\business\SaaS\templates\mcp-ultra>
PS E:\vertikon\business\SaaS\templates\mcp-ultra> # 3. Rodar linter e testes
PS E:\vertikon\business\SaaS\templates\mcp-ultra> golangci-lint run
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]\ninternal\\compliance\\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal\ninternal\\compliance\\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest\ninternal\\compliance\\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value\ninternal\\compliance\\framework_test.go:230:3: unknown field UserID in struct literal of type DataDeletionRequest\ninternal\\compliance\\framework_test.go:231:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal\ninternal\\compliance\\framework_test.go:236:17: assignment mismatch: 2 variables but framework.ProcessDataDeletionRequest returns 1 value\ninternal\\compliance\\framework_test.go:254:25: assignment mismatch: 2 variables but framework.AnonymizeData returns 1 value\ninternal\\compliance\\framework_test.go:254:68: too many arguments in call to framework.AnonymizeData\n\thave (context.Context, map[string]interface{}, string)\n\twant (context.Context, string)\ninternal\\compliance\\framework_test.go:279:46: too many arguments in call to framework.LogAuditEvent\n\thave (context.Context, uuid.UUID, string, map[string]interface{})\n\twant (context.Context, AuditEvent)\ninternal\\compliance\\framework_test.go:279:46: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/handlers/http [github.com/vertikon/mcp-ultra/internal/handlers/http.test]\ninternal\\handlers\\http\\router_test.go:23:76: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:25:42: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:38:75: undefined: services.HealthChecker\ninternal\\handlers\\http\\router_test.go:47:70: undefined: domain.CreateTaskRequest\ninternal\\handlers\\http\\router_test.go:60:85: undefined: domain.UpdateTaskRequest\ninternal\\handlers\\http\\router_test.go:70:73: undefined: domain.TaskFilters\ninternal\\handlers\\http\\router_test.go:70:95: undefined: domain.TaskList\ninternal\\handlers\\http\\router_test.go:72:30: undefined: domain.TaskList\ninternal\\handlers\\http\\router_test.go:80:49: not enough arguments in call to NewRouter\n\thave (*zap.Logger, *MockHealthService, *MockTaskService)\n\twant (*services.TaskService, *features.FlagManager, *HealthService, *zap.Logger)\ninternal\\handlers\\http\\router_test.go:101:77: undefined: services.HealthStatus\ninternal\\handlers\\http\\router_test.go:101:77: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/middleware [github.com/vertikon/mcp-ultra/internal/middleware.test]\ninternal\\middleware\\auth_test.go:96:30: undefined: testhelpers.GetTestAPIKeys"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/security [github.com/vertikon/mcp-ultra/internal/security.test]\ninternal\\security\\enhanced_auth_test.go:22:6: MockOPAService redeclared in this block\n\tinternal\\security\\auth_test.go:20:6: other declaration of MockOPAService\ninternal\\security\\enhanced_auth_test.go:26:26: method MockOPAService.IsAuthorized already declared at internal\\security\\auth_test.go:24:26\ninternal\\security\\enhanced_auth_test.go:36:6: TestNewAuthService redeclared in this block\n\tinternal\\security\\auth_test.go:39:6: other declaration of TestNewAuthService\ninternal\\security\\enhanced_auth_test.go:326:6: TestGetUserFromContext redeclared in this block\n\tinternal\\security\\auth_test.go:411:6: other declaration of TestGetUserFromContext\ninternal\\security\\enhanced_auth_test.go:391:6: TestRequireScope redeclared in this block\n\tinternal\\security\\auth_test.go:282:6: other declaration of TestRequireScope\ninternal\\security\\enhanced_auth_test.go:459:6: TestRequireRole redeclared in this block\n\tinternal\\security\\auth_test.go:342:6: other declaration of TestRequireRole\ninternal\\security\\auth_test.go:49:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:67:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:140:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:163:48: cannot use opa (variable of type *MockOPAService) as *OPAService value in argument to NewAuthService\ninternal\\security\\auth_test.go:163:48: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/internal/services [github.com/vertikon/mcp-ultra/internal/services.test]\ninternal\\services\\task_service_test.go:104:70: undefined: domain.UserFilter\ninternal\\services\\task_service_test.go:171:28: cannot use taskRepo (variable of type *mockTaskRepository) as domain.TaskRepository value in argument to NewTaskService: *mockTaskRepository does not implement domain.TaskRepository (wrong type for method List)\n\t\thave List(context.Context, domain.TaskFilter) ([]*domain.Task, error)\n\t\twant List(context.Context, domain.TaskFilter) ([]*domain.Task, int, error)\ninternal\\services\\task_service_test.go:171:48: cannot use eventRepo (variable of type *mockEventRepository) as domain.EventRepository value in argument to NewTaskService: *mockEventRepository does not implement domain.EventRepository (missing method GetByType)\ninternal\\services\\task_service_test.go:171:59: cannot use cacheRepo (variable of type *mockCacheRepository) as domain.CacheRepository value in argument to NewTaskService: *mockCacheRepository does not implement domain.CacheRepository (missing method Exists)\ninternal\\services\\task_service_test.go:199:31: declared and not used: eventRepo"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/compliance_test [github.com/vertikon/mcp-ultra/test/compliance.test]\ntest\\compliance\\compliance_integration_test.go:369:3: declared and not used: result"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/component [github.com/vertikon/mcp-ultra/test/component.test]\ntest\\component\\task_service_test.go:39:3: cannot use suite.taskRepo (variable of type *mocks.MockTaskRepository) as domain.TaskRepository value in argument to services.NewTaskService: *mocks.MockTaskRepository does not implement domain.TaskRepository (wrong type for method Delete)\n\t\thave Delete(context.Context, string) error\n\t\twant Delete(context.Context, uuid.UUID) error\ntest\\component\\task_service_test.go:40:3: cannot use suite.validator (variable of type *mocks.MockValidator) as domain.UserRepository value in argument to services.NewTaskService: *mocks.MockValidator does not implement domain.UserRepository (missing method Create)\ntest\\component\\task_service_test.go:42:3: cannot use suite.cacheRepo (variable of type *mocks.MockCacheRepository) as domain.CacheRepository value in argument to services.NewTaskService: *mocks.MockCacheRepository does not implement domain.CacheRepository (wrong type for method Get)\n\t\thave Get(context.Context, string) (interface{}, error)\n\t\twant Get(context.Context, string) (string, error)\ntest\\component\\task_service_test.go:44:3: cannot use suite.eventBus (variable of type *mocks.MockEventBus) as services.EventBus value in argument to services.NewTaskService: *mocks.MockEventBus does not implement services.EventBus (wrong type for method Publish)\n\t\thave Publish(context.Context, string, []byte) error\n\t\twant Publish(context.Context, *domain.Event) error\ntest\\component\\task_service_test.go:65:3: unknown field Metadata in struct literal of type services.CreateTaskRequest\ntest\\component\\task_service_test.go:78:20: req.Metadata undefined (type *services.CreateTaskRequest has no field or method Metadata)\ntest\\component\\task_service_test.go:97:55: too many arguments in call to suite.service.CreateTask\n\thave (context.Context, uuid.UUID, *services.CreateTaskRequest)\n\twant (context.Context, services.CreateTaskRequest)\ntest\\component\\task_service_test.go:118:29: undefined: services.ValidationError\ntest\\component\\task_service_test.go:127:55: too many arguments in call to suite.service.CreateTask\n\thave (context.Context, uuid.UUID, *services.CreateTaskRequest)\n\twant (context.Context, services.CreateTaskRequest)\ntest\\component\\task_service_test.go:151:52: too many arguments in call to suite.service.GetTask\n\thave (context.Context, uuid.UUID, uuid.UUID)\n\twant (context.Context, uuid.UUID)\ntest\\component\\task_service_test.go:151:52: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/integration [github.com/vertikon/mcp-ultra/test/integration.test]\ntest\\integration\\database_integration_test.go:71:19: undefined: testcontainers.NewLogWaitStrategy\ntest\\integration\\database_integration_test.go:121:21: undefined: postgresRepo.RunMigrations\ntest\\integration\\database_integration_test.go:141:23: suite.taskRepo.DB undefined (type *\"github.com/vertikon/mcp-ultra/internal/repository/postgres\".TaskRepository has no field or method DB)\ntest\\integration\\database_integration_test.go:146:28: suite.cacheRepo.Client undefined (type *\"github.com/vertikon/mcp-ultra/internal/repository/redis\".CacheRepository has no field or method Client, but does have unexported field client)\ntest\\integration\\database_integration_test.go:170:22: assignment mismatch: 2 variables but suite.taskRepo.Create returns 1 value\ntest\\integration\\database_integration_test.go:188:22: assignment mismatch: 2 variables but suite.taskRepo.Update returns 1 value\ntest\\integration\\database_integration_test.go:195:24: assignment mismatch: 2 variables but suite.taskRepo.Update returns 1 value\ntest\\integration\\database_integration_test.go:202:3: unknown field UserID in struct literal of type domain.TaskFilter\ntest\\integration\\database_integration_test.go:203:11: cannot use domain.TaskStatusCompleted (constant \"completed\" of string type domain.TaskStatus) as []domain.TaskStatus value in struct literal\ntest\\integration\\database_integration_test.go:208:48: cannot use filter (variable of type *domain.TaskFilter) as domain.TaskFilter value in argument to suite.taskRepo.List\ntest\\integration\\database_integration_test.go:208:48: too many errors"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/observability_test [github.com/vertikon/mcp-ultra/test/observability.test]\ntest\\observability\\integration_test.go:101:21: telemetryService.CreateAttribute undefined (type *observability.TelemetryService has no field or method CreateAttribute)\ntest\\observability\\integration_test.go:102:21: telemetryService.CreateAttribute undefined (type *observability.TelemetryService has no field or method CreateAttribute)\ntest\\observability\\integration_test.go:110:21: telemetryService.CreateAttribute undefined (type *observability.TelemetryService has no field or method CreateAttribute)\ntest\\observability\\integration_test.go:128:20: telemetryService.IncrementCounter undefined (type *observability.TelemetryService has no field or method IncrementCounter)"
level=error msg="[linters_context] typechecking error: : # github.com/vertikon/mcp-ultra/test/property [github.com/vertikon/mcp-ultra/test/property.test]\ntest\\property\\task_properties_test.go:231:4: declared and not used: originalTitle"
automation\autocommit.go:1:1: package-comments: should have a package comment (revive)
package main
^
automation\autocommit.go:16:1: Comment should end in a period (godot)
// Config represents the configuration for the auto-commit tool
^
automation\autocommit.go:17:13: fieldalignment: struct with 144 pointer bytes could be 136 (govet)
type Config struct {
            ^
automation\autocommit.go:31:1: Comment should end in a period (godot)
// DefaultConfig returns a default configuration
^
automation\autocommit.go:46:1: Comment should end in a period (godot)
// ensureDirectory creates directory structure if it doesn't exist
^
automation\autocommit.go:50:10: G301: Expect directory permissions to be 0750 or less (gosec)
                return os.MkdirAll(path, 0755)
                       ^
automation\autocommit.go:50:28: octalLiteral: use new octal literal style, 0o755 (gocritic)
                return os.MkdirAll(path, 0755)
                                         ^
automation\autocommit.go:55:1: Comment should end in a period (godot)
// runCommand executes a shell command and returns output
^
automation\autocommit.go:56:22: `runCommand` - `command` always receives `"git"` (unparam)
func runCommand(dir, command string, args ...string) (string, error) {
                     ^
automation\autocommit.go:72:1: Comment should end in a period (godot)
// initializeGitRepo initializes a git repository if it doesn't exist
^
automation\autocommit.go:84:1: `if os.IsNotExist(err)` has complex nested blocks (complexity: 6) (nestif)
        if _, err := os.Stat(filepath.Join(repoPath, ".git")); os.IsNotExist(err) {
^
automation\autocommit.go:103:13: G306: Expect WriteFile permissions to be 0600 or less (gosec)
                if err := ioutil.WriteFile(gitignorePath, []byte(config.GitIgnore), 0644); err != nil {
                          ^
automation\autocommit.go:103:71: octalLiteral: use new octal literal style, 0o644 (gocritic)
                if err := ioutil.WriteFile(gitignorePath, []byte(config.GitIgnore), 0644); err != nil {
                                                                                    ^
automation\autocommit.go:109:1: The line is 842 characters long, which exceeds the maximum of 140 characters. (lll)
                readmeContent := fmt.Sprintf("# %s\n\n✨ Repositório criado automaticamente via **MCP Ultra** by Vertikon.\n\n🤖 **MCP Ultra Features:**\n- ✅ Criação automática de repositórios GitHub\n- ✅ Automação completa de commits e push\n- ✅ Integração MCP Server <-> GitHub API\n- ✅ Gerenciamento de diretórios locais\n- ✅ Scripts de setup automático\n- ✅ Pipeline de testes end-to-end\n\n⏰ **Criado em:** %s\n🏢 **Organização:** %s\n🔧 **Template:** [MCP Ultra](https://github.com/vertikon/mcp-ultra)\n\n---\n\n🚀 **Próximos passos:**\n1. Clone o repositório: `git clone %s`\n2. Adicione seus arquivos e código\n3. Use `autocommit commit %s` para commits automáticos\n4. Explore as ferramentas MCP Ultra disponíveis\n\n💡 **Dica:** Este repositório foi criado com MCP Ultra, um template completo para automação GitHub desenvolvido pela Vertikon.\n",
^
automation\autocommit.go:117:13: G306: Expect WriteFile permissions to be 0600 or less (gosec)
                if err := ioutil.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
                          ^
automation\autocommit.go:117:65: octalLiteral: use new octal literal style, 0o644 (gocritic)
                if err := ioutil.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
                                                                              ^
automation\autocommit.go:133:1: Comment should end in a period (godot)
// commitAndPush commits changes and pushes to GitHub
^
automation\autocommit.go:190:1: Comment should end in a period (godot)
// loadConfigFromFile loads configuration from JSON file
^
automation\autocommit.go:199:15: G304: Potential file inclusion via variable (gosec)
        data, err := ioutil.ReadFile(filename)
                     ^
automation\autocommit.go:212:1: Comment should end in a period (godot)
// saveConfigToFile saves configuration to JSON file
^
automation\autocommit.go:219:12: G306: Expect WriteFile permissions to be 0600 or less (gosec)
        if err := ioutil.WriteFile(filename, data, 0644); err != nil {
                  ^
automation\autocommit.go:219:45: octalLiteral: use new octal literal style, 0o644 (gocritic)
        if err := ioutil.WriteFile(filename, data, 0644); err != nil {
                                                   ^
automation\autocommit.go:227:1: Comment should end in a period (godot)
// interactiveConfig allows user to input configuration interactively
^
automation\autocommit.go:236:12: Error return value of `reader.ReadString` is not checked (errcheck)
        if token, _ := reader.ReadString('\n'); strings.TrimSpace(token) != "" {
                  ^
automation\autocommit.go:241:10: Error return value of `reader.ReadString` is not checked (errcheck)
        if org, _ := reader.ReadString('\n'); strings.TrimSpace(org) != "" {
                ^
automation\autocommit.go:246:11: Error return value of `reader.ReadString` is not checked (errcheck)
        if repo, _ := reader.ReadString('\n'); strings.TrimSpace(repo) != "" {
                 ^
automation\autocommit.go:251:11: Error return value of `reader.ReadString` is not checked (errcheck)
        if path, _ := reader.ReadString('\n'); strings.TrimSpace(path) != "" {
                 ^
automation\autocommit.go:256:10: Error return value of `reader.ReadString` is not checked (errcheck)
        if msg, _ := reader.ReadString('\n'); strings.TrimSpace(msg) != "" {
                ^
basic_test.go:1:1: package-comments: should have a package comment (revive)
package main
^
basic_test.go:7:1: Comment should end in a period (godot)
// TestBasic is a basic test to ensure the test runner works
^
basic_test.go:14:1: Comment should end in a period (godot)
// TestVersion tests that version constants are not empty
^
basic_test.go:18:5: SA4000: identical expressions on the left and right side of the '!=' operator (staticcheck)
        if true != true {
           ^
basic_test.go:18:5: dupSubExpr: suspicious identical LHS and RHS for `!=` operator (gocritic)
        if true != true {
           ^
internal\ai\events\handlers.go:1:1: package-comments: should have a package comment (revive)
package events
^
internal\ai\events\handlers.go:15:6: exported: exported type Base should have comment or be unexported (revive)
type Base struct {
     ^
internal\ai\events\handlers.go:22:6: exported: exported type RouterDecision should have comment or be unexported (revive)
type RouterDecision struct {
     ^
internal\ai\events\handlers.go:30:6: exported: exported type PolicyBlock should have comment or be unexported (revive)
type PolicyBlock struct {
     ^
internal\ai\events\handlers.go:37:6: exported: exported type InferenceError should have comment or be unexported (revive)
type InferenceError struct {
     ^
internal\ai\events\handlers.go:45:6: exported: exported type InferenceSummary should have comment or be unexported (revive)
type InferenceSummary struct {
     ^
internal\ai\events\handlers.go:57:1: exported: exported function PublishRouterDecision should have comment or be unexported (revive)
func PublishRouterDecision(ctx context.Context, pub EventPublisher, subject string, e RouterDecision) error {
^
internal\ai\events\handlers.go:59:5: Error return value of `json.Marshal` is not checked (errcheck)
        b, _ := json.Marshal(e)
           ^
internal\ai\events\handlers.go:63:1: exported: exported function PublishPolicyBlock should have comment or be unexported (revive)
func PublishPolicyBlock(ctx context.Context, pub EventPublisher, subject string, e PolicyBlock) error {
^
internal\ai\events\handlers.go:65:5: Error return value of `json.Marshal` is not checked (errcheck)
        b, _ := json.Marshal(e)
           ^
internal\ai\events\handlers.go:69:1: exported: exported function PublishInferenceError should have comment or be unexported (revive)
func PublishInferenceError(ctx context.Context, pub EventPublisher, subject string, e InferenceError) error {
^
internal\ai\events\handlers.go:71:5: Error return value of `json.Marshal` is not checked (errcheck)
        b, _ := json.Marshal(e)
           ^
internal\ai\events\handlers.go:75:1: exported: exported function PublishInferenceSummary should have comment or be unexported (revive)
func PublishInferenceSummary(ctx context.Context, pub EventPublisher, subject string, e InferenceSummary) error {
^
internal\ai\events\handlers.go:77:5: Error return value of `json.Marshal` is not checked (errcheck)
        b, _ := json.Marshal(e)
           ^
internal\ai\events\handlers_test.go:9:1: Comment should end in a period (godot)
// Mock publisher for testing
^
internal\ai\events\handlers_test.go:19:42: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (m *mockPublisher) PublishWithRetry(ctx context.Context, subject string, payload []byte) error {
                                         ^
internal\ai\router\router.go:1:1: package-comments: should have a package comment (revive)
package router
^
internal\ai\router\router.go:11:6: exported: exported type Flags should have comment or be unexported (revive)
type Flags struct {
     ^
internal\ai\router\router.go:12:5: fieldalignment: struct with 40 pointer bytes could be 24 (govet)
        AI struct {
           ^
internal\ai\router\router.go:20:6: exported: exported type Rule should have comment or be unexported (revive)
type Rule struct {
     ^
internal\ai\router\router.go:25:6: exported: exported type Rules should have comment or be unexported (revive)
type Rules struct {
     ^
internal\ai\router\router.go:35:6: exported: exported type Decision should have comment or be unexported (revive)
type Decision struct {
     ^
internal\ai\router\router.go:41:6: exported: exported type Router should have comment or be unexported (revive)
type Router struct {
     ^
internal\ai\router\router.go:47:1: exported: exported function Load should have comment or be unexported (revive)
func Load(basePath string) (*Router, error) {
^
internal\ai\router\router.go:52:15: G304: Potential file inclusion via variable (gosec)
        if b, err := os.ReadFile(ff); err == nil {
                     ^
internal\ai\router\router.go:53:3: Error return value of `json.Unmarshal` is not checked (errcheck)
                _ = json.Unmarshal(b, &r.flags)
                ^
internal\ai\router\router.go:55:15: G304: Potential file inclusion via variable (gosec)
        if b, err := os.ReadFile(rules); err == nil {
                     ^
internal\ai\router\router.go:56:3: Error return value of `json.Unmarshal` is not checked (errcheck)
                _ = json.Unmarshal(b, &r.rules)
                ^
internal\ai\router\router.go:61:1: exported: exported method Router.Enabled should have comment or be unexported (revive)
func (r *Router) Enabled() bool {
^
internal\ai\router\router.go:67:1: exported: exported method Router.Decide should have comment or be unexported (revive)
func (r *Router) Decide(useCase string) (Decision, error) {
^
internal\ai\telemetry\metrics.go:1:1: package-comments: should have a package comment (revive)
package telemetry
^
internal\ai\telemetry\metrics.go:25:6: exported: exported type Labels should have comment or be unexported (revive)
type Labels struct {
     ^
internal\ai\telemetry\metrics.go:38:1: exported: exported function Init should have comment or be unexported (revive)
func Init(reg prometheus.Registerer) {
^
internal\ai\telemetry\metrics.go:79:6: exported: exported type InferenceMeta should have comment or be unexported (revive)
type InferenceMeta struct {
     ^
internal\ai\telemetry\metrics.go:79:20: fieldalignment: struct with 232 pointer bytes could be 200 (govet)
type InferenceMeta struct {
                   ^
internal\ai\telemetry\metrics.go:88:1: exported: exported function ObserveStart should have comment or be unexported (revive)
func ObserveStart() time.Time { return time.Now() }
^
internal\ai\telemetry\metrics.go:90:1: exported: exported function ObserveInference should have comment or be unexported (revive)
func ObserveInference(meta InferenceMeta) {
^
internal\ai\telemetry\metrics.go:112:1: exported: exported function IncPolicyBlock should have comment or be unexported (revive)
func IncPolicyBlock(l Labels) {
^
internal\ai\telemetry\metrics.go:119:1: exported: exported function IncRouterDecision should have comment or be unexported (revive)
func IncRouterDecision(l Labels) {
^
internal\ai\telemetry\metrics.go:126:1: exported: exported function IncBudgetBreach should have comment or be unexported (revive)
func IncBudgetBreach(scope string) {
^
internal\ai\telemetry\metrics_test.go:126:33: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
func TestNoOpWhenNotInitialized(t *testing.T) {
                                ^
internal\ai\wiring\wiring.go:1:1: package-comments: should have a package comment (revive)
package wiring
^
internal\ai\wiring\wiring.go:15:6: exported: exported type Config should have comment or be unexported (revive)
type Config struct {
     ^
internal\ai\wiring\wiring.go:15:13: fieldalignment: struct with 32 pointer bytes could be 24 (govet)
type Config struct {
            ^
internal\ai\wiring\wiring.go:26:1: exported: exported function Init should have comment or be unexported (revive)
func Init(ctx context.Context, cfg Config) (*Service, error) {
^
internal\ai\wiring\wiring.go:30:8: Error return value of `os.Getwd` is not checked (errcheck)
                cwd, _ := os.Getwd()
                     ^
internal\ai\wiring\wiring.go:34:5: Error return value of `router.Load` is not checked (errcheck)
        r, _ := router.Load(base)
           ^
internal\ai\wiring\wiring_test.go:16:31: octalLiteral: use new octal literal style, 0o755 (gocritic)
        if err := os.MkdirAll(aiDir, 0755); err != nil {
                                     ^
internal\ai\wiring\wiring_test.go:22:91: octalLiteral: use new octal literal style, 0o644 (gocritic)
        if err := os.WriteFile(filepath.Join(aiDir, "feature_flags.json"), []byte(flagsContent), 0644); err != nil {
                                                                                                 ^
internal\ai\wiring\wiring_test.go:55:35: octalLiteral: use new octal literal style, 0o755 (gocritic)
        if err := os.MkdirAll(configDir, 0755); err != nil {
                                         ^
internal\ai\wiring\wiring_test.go:61:91: octalLiteral: use new octal literal style, 0o644 (gocritic)
        if err := os.WriteFile(filepath.Join(aiDir, "feature_flags.json"), []byte(flagsContent), 0644); err != nil {
                                                                                                 ^
internal\ai\wiring\wiring_test.go:75:97: octalLiteral: use new octal literal style, 0o644 (gocritic)
        if err := os.WriteFile(filepath.Join(configDir, "ai-router.rules.json"), []byte(rulesContent), 0644); err != nil {
                                                                                                       ^
internal\cache\circuit_breaker.go:1:1: package-comments: should have a package comment (revive)
package cache
^
internal\cache\circuit_breaker.go:8:1: Comment should end in a period (godot)
// CircuitBreakerState represents the state of a circuit breaker
^
internal\cache\circuit_breaker.go:12:2: exported: exported const CircuitBreakerClosed should have comment (or a comment on this block) or be unexported (revive)
        CircuitBreakerClosed CircuitBreakerState = iota
        ^
internal\cache\circuit_breaker.go:17:1: Comment should end in a period (godot)
// String returns string representation of circuit breaker state
^
internal\cache\circuit_breaker.go:31:1: Comment should end in a period (godot)
// CircuitBreakerConfig configures circuit breaker behavior
^
internal\cache\circuit_breaker.go:39:1: Comment should end in a period (godot)
// CircuitBreaker implements the circuit breaker pattern
^
internal\cache\circuit_breaker.go:40:21: fieldalignment: struct with 120 pointer bytes could be 32 (govet)
type CircuitBreaker struct {
                    ^
internal\cache\circuit_breaker.go:56:1: Comment should end in a period (godot)
// NewCircuitBreaker creates a new circuit breaker
^
internal\cache\circuit_breaker.go:67:1: Comment should end in a period (godot)
// Allow checks if the request should be allowed through
^
internal\cache\circuit_breaker.go:99:1: Comment should end in a period (godot)
// RecordSuccess records a successful operation
^
internal\cache\circuit_breaker.go:104:2: missing cases in switch of type cache.CircuitBreakerState: cache.CircuitBreakerOpen (exhaustive)
        switch cb.state {
        ^
internal\cache\circuit_breaker.go:119:1: Comment should end in a period (godot)
// RecordFailure records a failed operation
^
internal\cache\circuit_breaker.go:126:2: missing cases in switch of type cache.CircuitBreakerState: cache.CircuitBreakerOpen (exhaustive)
        switch cb.state {
        ^
internal\cache\circuit_breaker.go:140:1: Comment should end in a period (godot)
// State returns the current state of the circuit breaker
^
internal\cache\circuit_breaker.go:148:1: Comment should end in a period (godot)
// Stats returns circuit breaker statistics
^
internal\cache\circuit_breaker.go:166:1: Comment should end in a period (godot)
// CircuitBreakerStats contains circuit breaker statistics
^
internal\cache\circuit_breaker.go:167:26: fieldalignment: struct with 56 pointer bytes could be 24 (govet)
type CircuitBreakerStats struct {
                         ^
internal\cache\circuit_breaker.go:179:1: Comment should end in a period (godot)
// OnStateChange sets a callback for state changes
^
internal\cache\circuit_breaker.go:187:1: Comment should end in a period (godot)
// Reset resets the circuit breaker to closed state
^
internal\cache\circuit_breaker.go:204:1: Comment should end in a period (godot)
// ForceOpen forces the circuit breaker to open state
^
internal\cache\circuit_breaker.go:218:1: Comment should end in a period (godot)
// setState sets the state and triggers callback if registered
^
internal\cache\circuit_breaker.go:228:1: Comment should end in a period (godot)
// AdaptiveCircuitBreaker extends CircuitBreaker with adaptive behavior
^
internal\cache\circuit_breaker.go:229:29: fieldalignment: struct with 64 pointer bytes could be 40 (govet)
type AdaptiveCircuitBreaker struct {
                            ^
internal\cache\circuit_breaker.go:243:1: Comment should end in a period (godot)
// NewAdaptiveCircuitBreaker creates an adaptive circuit breaker
^
internal\cache\circuit_breaker.go:269:1: Comment should end in a period (godot)
// RecordRequest records a request (for adaptive behavior)
^
internal\cache\circuit_breaker.go:286:1: Comment should end in a period (godot)
// RecordFailure records a failure with adaptive behavior
^
internal\cache\circuit_breaker.go:305:1: Comment should end in a period (godot)
// GetFailureRate returns the current failure rate
^
internal\cache\circuit_breaker.go:317:1: Comment should end in a period (godot)
// adaptiveAdjustment runs in background to adjust thresholds
^
internal\cache\circuit_breaker.go:351:1: Comment should end in a period (godot)
// Helper functions
^
internal\cache\circuit_breaker.go:352:1: redefines-builtin-id: redefinition of the built-in function max (revive)
func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}
internal\cache\circuit_breaker.go:352:6: builtinShadowDecl: shadowing of predeclared identifier: max (gocritic)
func max(a, b int) int {
     ^
internal\cache\circuit_breaker.go:359:1: redefines-builtin-id: redefinition of the built-in function min (revive)
func min(a, b int) int {
        if a < b {
                return a
        }
        return b
}
internal\cache\circuit_breaker.go:359:6: builtinShadowDecl: shadowing of predeclared identifier: min (gocritic)
func min(a, b int) int {
     ^
internal\cache\consistent_hash.go:11:1: Comment should end in a period (godot)
// ConsistentHash provides consistent hashing for distributed caching
^
internal\cache\consistent_hash.go:12:21: fieldalignment: struct with 72 pointer bytes could be 24 (govet)
type ConsistentHash struct {
                    ^
internal\cache\consistent_hash.go:20:1: Comment should end in a period (godot)
// NewConsistentHash creates a new consistent hash ring
^
internal\cache\consistent_hash.go:30:1: Comment should end in a period (godot)
// Add adds a node to the hash ring
^
internal\cache\consistent_hash.go:56:1: Comment should end in a period (godot)
// Remove removes a node from the hash ring
^
internal\cache\consistent_hash.go:80:1: Comment should end in a period (godot)
// Get returns the node responsible for the given key
^
internal\cache\consistent_hash.go:107:1: Comment should end in a period (godot)
// GetMultiple returns multiple nodes for replication
^
internal\cache\consistent_hash.go:145:1: Comment should end in a period (godot)
// GetNodes returns all nodes in the hash ring
^
internal\cache\consistent_hash.go:158:1: Comment should end in a period (godot)
// Size returns the number of nodes in the hash ring
^
internal\cache\consistent_hash.go:166:1: Comment should end in a period (godot)
// Distribution returns the distribution of keys across nodes
^
internal\cache\consistent_hash.go:205:1: Comment should end in a period (godot)
// hash generates a hash for the given key
^
internal\cache\consistent_hash.go:211:1: Comment should end in a period (godot)
// RebalanceInfo provides information about data that needs to be moved when nodes change
^
internal\cache\consistent_hash.go:219:1: Comment should end in a period (godot)
// KeyRange represents a range of keys
^
internal\cache\consistent_hash.go:225:1: Comment should end in a period (godot)
// GetRebalanceInfo returns information about what data needs to be moved when adding/removing nodes
^
internal\cache\consistent_hash.go:244:20: G115: integer overflow conversion uint64 -> uint32 (gosec)
                keyHash := uint32(uint64(i) * maxHash / uint64(samplePoints))
                                 ^
internal\cache\consistent_hash.go:244:27: G115: integer overflow conversion int -> uint64 (gosec)
                keyHash := uint32(uint64(i) * maxHash / uint64(samplePoints))
                                        ^
internal\cache\consistent_hash.go:244:49: G115: integer overflow conversion int -> uint64 (gosec)
                keyHash := uint32(uint64(i) * maxHash / uint64(samplePoints))
                                                              ^
internal\cache\consistent_hash.go:268:1: Comment should end in a period (godot)
// getNodeForHash returns the node for a given hash (internal method)
^
internal\cache\distributed.go:17:1: Comment should end in a period (godot)
// CacheStrategy represents different caching strategies
^
internal\cache\distributed.go:18:6: exported: type name will be used as cache.CacheStrategy by other packages, and that stutters; consider calling this Strategy (revive)
type CacheStrategy string
     ^
internal\cache\distributed.go:21:2: exported: exported const StrategyWriteThrough should have comment (or a comment on this block) or be unexported (revive)
        StrategyWriteThrough CacheStrategy = "write_through"
        ^
internal\cache\distributed.go:27:1: Comment should end in a period (godot)
// EvictionPolicy represents cache eviction policies
^
internal\cache\distributed.go:31:2: exported: exported const EvictionLRU should have comment (or a comment on this block) or be unexported (revive)
        EvictionLRU    EvictionPolicy = "lru"
        ^
internal\cache\distributed.go:37:1: Comment should end in a period (godot)
// CacheConfig configures the distributed cache system
^
internal\cache\distributed.go:38:6: exported: type name will be used as cache.CacheConfig by other packages, and that stutters; consider calling this Config (revive)
type CacheConfig struct {
     ^
internal\cache\distributed.go:38:18: fieldalignment: struct of size 296 could be 272 (govet)
type CacheConfig struct {
                 ^
internal\cache\distributed.go:83:1: Comment should end in a period (godot)
// DefaultCacheConfig returns default cache configuration
^
internal\cache\distributed.go:116:1: Comment should end in a period (godot)
// DistributedCache provides distributed caching capabilities
^
internal\cache\distributed.go:117:23: fieldalignment: struct with 552 pointer bytes could be 464 (govet)
type DistributedCache struct {
                      ^
internal\cache\distributed.go:139:1: Comment should end in a period (godot)
// CacheShard represents a cache shard
^
internal\cache\distributed.go:140:6: exported: type name will be used as cache.CacheShard by other packages, and that stutters; consider calling this Shard (revive)
type CacheShard struct {
     ^
internal\cache\distributed.go:140:17: fieldalignment: struct with 72 pointer bytes could be 48 (govet)
type CacheShard struct {
                ^
internal\cache\distributed.go:148:1: Comment should end in a period (godot)
// WriteOperation represents a write operation in write-behind mode
^
internal\cache\distributed.go:149:21: fieldalignment: struct with 80 pointer bytes could be 64 (govet)
type WriteOperation struct {
                    ^
internal\cache\distributed.go:157:1: Comment should end in a period (godot)
// CacheStats tracks cache performance metrics
^
internal\cache\distributed.go:158:6: exported: type name will be used as cache.CacheStats by other packages, and that stutters; consider calling this Stats (revive)
type CacheStats struct {
     ^
internal\cache\distributed.go:158:17: fieldalignment: struct with 104 pointer bytes could be 24 (govet)
type CacheStats struct {
                ^
internal\cache\distributed.go:174:1: Comment should end in a period (godot)
// CacheEntry represents a cached item with metadata
^
internal\cache\distributed.go:175:6: exported: type name will be used as cache.CacheEntry by other packages, and that stutters; consider calling this Entry (revive)
type CacheEntry struct {
     ^
internal\cache\distributed.go:175:17: fieldalignment: struct with 120 pointer bytes could be 96 (govet)
type CacheEntry struct {
                ^
internal\cache\distributed.go:187:1: Comment should end in a period (godot)
// NewDistributedCache creates a new distributed cache instance
^
internal\cache\distributed.go:188:26: hugeParam: config is heavy (296 bytes); consider passing it by pointer (gocritic)
func NewDistributedCache(config CacheConfig, log *logger.Logger, telemetry *observability.TelemetryService) (*DistributedCache, error) {
                         ^
internal\cache\distributed.go:252:1: Comment should end in a period (godot)
// Set stores a value in the cache with the specified TTL
^
internal\cache\distributed.go:315:1: Comment should end in a period (godot)
// Get retrieves a value from the cache
^
internal\cache\distributed.go:378:1: Comment should end in a period (godot)
// Delete removes a key from the cache
^
internal\cache\distributed.go:411:1: Comment should end in a period (godot)
// Exists checks if a key exists in the cache
^
internal\cache\distributed.go:427:1: Comment should end in a period (godot)
// Expire sets the TTL for a key
^
internal\cache\distributed.go:443:1: Comment should end in a period (godot)
// Clear removes all keys matching the pattern
^
internal\cache\distributed.go:499:1: Comment should end in a period (godot)
// GetStats returns cache performance statistics
^
internal\cache\distributed.go:529:1: Comment should end in a period (godot)
// ResetStats resets cache statistics
^
internal\cache\distributed.go:537:1: Comment should end in a period (godot)
// Close gracefully shuts down the cache
^
internal\cache\distributed.go:552:1: exported: comment on exported method DistributedCache.HealthCheck should be of the form "HealthCheck ..." (revive)
// Health check for the cache
^
internal\cache\distributed.go:552:1: Comment should end in a period (godot)
// Health check for the cache
^
internal\cache\distributed.go:626:5: comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error (errorlint)
        if err == redis.Nil {
           ^
internal\cache\distributed.go:635:44: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (dc *DistributedCache) getReadThrough(ctx context.Context, key string) (interface{}, bool, error) {
                                           ^
internal\cache\distributed.go:635:65: unused-parameter: parameter 'key' seems to be unused, consider removing or renaming it as _ (revive)
func (dc *DistributedCache) getReadThrough(ctx context.Context, key string) (interface{}, bool, error) {
                                                                ^
internal\cache\distributed_test.go:12:1: File is not properly formatted (gci)
        "github.com/stretchr/testify/require"
^
internal\cache\distributed_test.go:187:21: fieldalignment: struct with 32 pointer bytes could be 24 (govet)
        type ComplexObject struct {
                           ^
internal\config\config.go:1:1: package-comments: should have a package comment (revive)
package config
^
internal\config\config.go:14:1: Comment should end in a period (godot)
// Config represents the application configuration
^
internal\config\config.go:15:13: fieldalignment: struct with 1464 pointer bytes could be 1312 (govet)
type Config struct {
            ^
internal\config\config.go:29:1: Comment should end in a period (godot)
// ComplianceConfig holds all compliance-related configuration
^
internal\config\config.go:30:23: fieldalignment: struct with 528 pointer bytes could be 512 (govet)
type ComplianceConfig struct {
                      ^
internal\config\config.go:43:1: Comment should end in a period (godot)
// PIIDetectionConfig configures PII detection and classification
^
internal\config\config.go:44:25: fieldalignment: struct of size 64 could be 56 (govet)
type PIIDetectionConfig struct {
                        ^
internal\config\config.go:52:1: Comment should end in a period (godot)
// ConsentConfig configures consent management
^
internal\config\config.go:53:20: fieldalignment: struct with 48 pointer bytes could be 24 (govet)
type ConsentConfig struct {
                   ^
internal\config\config.go:60:1: Comment should end in a period (godot)
// DataRetentionConfig configures data retention policies
^
internal\config\config.go:61:26: fieldalignment: struct of size 40 could be 32 (govet)
type DataRetentionConfig struct {
                         ^
internal\config\config.go:69:1: Comment should end in a period (godot)
// AuditLoggingConfig configures compliance audit logging
^
internal\config\config.go:70:25: fieldalignment: struct of size 56 could be 48 (govet)
type AuditLoggingConfig struct {
                        ^
internal\config\config.go:79:1: Comment should end in a period (godot)
// LGPDConfig specific configuration for Brazilian LGPD compliance
^
internal\config\config.go:80:17: fieldalignment: struct of size 80 could be 72 (govet)
type LGPDConfig struct {
                ^
internal\config\config.go:88:1: Comment should end in a period (godot)
// GDPRConfig specific configuration for European GDPR compliance
^
internal\config\config.go:89:17: fieldalignment: struct of size 104 could be 96 (govet)
type GDPRConfig struct {
                ^
internal\config\config.go:98:1: Comment should end in a period (godot)
// AnonymizationConfig configures data anonymization
^
internal\config\config.go:99:26: fieldalignment: struct of size 72 could be 64 (govet)
type AnonymizationConfig struct {
                         ^
internal\config\config.go:108:1: Comment should end in a period (godot)
// DataRightsConfig configures individual data rights handling
^
internal\config\config.go:109:23: fieldalignment: struct of size 48 could be 40 (govet)
type DataRightsConfig struct {
                      ^
internal\config\config.go:117:1: Comment should end in a period (godot)
// ServerConfig holds HTTP server configuration
^
internal\config\config.go:125:1: Comment should end in a period (godot)
// GRPCConfig holds gRPC server configuration
^
internal\config\config.go:135:1: Comment should end in a period (godot)
// KeepaliveConfig holds gRPC keepalive configuration
^
internal\config\config.go:146:1: Comment should end in a period (godot)
// DatabaseConfig holds database connections configuration
^
internal\config\config.go:147:21: fieldalignment: struct with 136 pointer bytes could be 128 (govet)
type DatabaseConfig struct {
                    ^
internal\config\config.go:152:1: Comment should end in a period (godot)
// PostgreSQLConfig holds PostgreSQL configuration
^
internal\config\config.go:153:23: fieldalignment: struct with 80 pointer bytes could be 72 (govet)
type PostgreSQLConfig struct {
                      ^
internal\config\config.go:165:1: Comment should end in a period (godot)
// RedisConfig holds Redis configuration
^
internal\config\config.go:173:1: Comment should end in a period (godot)
// NATSConfig holds NATS configuration
^
internal\config\config.go:180:1: Comment should end in a period (godot)
// TelemetryConfig holds comprehensive telemetry configuration
^
internal\config\config.go:181:22: fieldalignment: struct of size 272 could be 264 (govet)
type TelemetryConfig struct {
                     ^
internal\config\config.go:198:1: Comment should end in a period (godot)
// TracingConfig holds distributed tracing configuration
^
internal\config\config.go:207:1: Comment should end in a period (godot)
// MetricsConfig holds metrics collection configuration
^
internal\config\config.go:208:20: fieldalignment: struct with 48 pointer bytes could be 24 (govet)
type MetricsConfig struct {
                   ^
internal\config\config.go:216:1: Comment should end in a period (godot)
// ExportersConfig holds exporter configurations
^
internal\config\config.go:217:22: fieldalignment: struct with 96 pointer bytes could be 88 (govet)
type ExportersConfig struct {
                     ^
internal\config\config.go:228:1: Comment should end in a period (godot)
// JaegerConfig holds Jaeger exporter configuration
^
internal\config\config.go:229:19: fieldalignment: struct with 48 pointer bytes could be 40 (govet)
type JaegerConfig struct {
                  ^
internal\config\config.go:236:1: Comment should end in a period (godot)
// OTLPConfig holds OTLP exporter configuration
^
internal\config\config.go:237:17: fieldalignment: struct of size 40 could be 32 (govet)
type OTLPConfig struct {
                ^
internal\config\config.go:244:1: Comment should end in a period (godot)
// ConsoleConfig holds console exporter configuration
^
internal\config\config.go:249:1: Comment should end in a period (godot)
// FeaturesConfig holds feature flags configuration
^
internal\config\config.go:255:1: Comment should end in a period (godot)
// SecurityConfig holds all security-related configuration
^
internal\config\config.go:256:21: fieldalignment: struct with 256 pointer bytes could be 240 (govet)
type SecurityConfig struct {
                    ^
internal\config\config.go:263:1: Comment should end in a period (godot)
// Load loads configuration from file and environment variables
^
internal\config\config.go:283:1: Comment should end in a period (godot)
// loadFromFile loads configuration from YAML file
^
internal\config\config.go:295:1: Comment should end in a period (godot)
// getEnv returns environment variable value or default
^
internal\config\config.go:303:1: Comment should end in a period (godot)
// DSN returns PostgreSQL connection string
^
internal\config\secrets\loader.go:1:1: package-comments: should have a package comment (revive)
package config
^
internal\config\secrets\loader.go:14:1: Comment should end in a period (godot)
// SecretsBackendType define o tipo de backend de secrets
^
internal\config\secrets\loader.go:18:2: exported: exported const SecretsBackendEnv should have comment (or a comment on this block) or be unexported (revive)
        SecretsBackendEnv   SecretsBackendType = "env"
        ^
internal\config\secrets\loader.go:23:1: Comment should end in a period (godot)
// SecretsConfig representa a configuração de secrets
^
internal\config\secrets\loader.go:24:20: fieldalignment: struct with 344 pointer bytes could be 328 (govet)
type SecretsConfig struct {
                   ^
internal\config\secrets\loader.go:36:1: Comment should end in a period (godot)
// SecretsBackendConfig configura o backend de secrets
^
internal\config\secrets\loader.go:37:27: fieldalignment: struct with 24 pointer bytes could be 16 (govet)
type SecretsBackendConfig struct {
                          ^
internal\config\secrets\loader.go:42:1: Comment should end in a period (godot)
// VaultConfig configuração do Vault
^
internal\config\secrets\loader.go:49:1: Comment should end in a period (godot)
// DatabaseSecrets secrets do banco de dados
^
internal\config\secrets\loader.go:59:1: Comment should end in a period (godot)
// NATSSecrets secrets do NATS
^
internal\config\secrets\loader.go:67:1: Comment should end in a period (godot)
// TelemetrySecrets secrets de telemetria
^
internal\config\secrets\loader.go:73:1: Comment should end in a period (godot)
// OTLPSecrets configuração OTLP
^
internal\config\secrets\loader.go:74:18: fieldalignment: struct with 24 pointer bytes could be 16 (govet)
type OTLPSecrets struct {
                 ^
internal\config\secrets\loader.go:79:1: Comment should end in a period (godot)
// PrometheusSecrets configuração Prometheus
^
internal\config\secrets\loader.go:85:1: Comment should end in a period (godot)
// AuthSecrets secrets de autenticação
^
internal\config\secrets\loader.go:91:1: Comment should end in a period (godot)
// EncryptionSecrets secrets de criptografia
^
internal\config\secrets\loader.go:97:1: Comment should end in a period (godot)
// SecretsLoader carrega secrets de diferentes fontes
^
internal\config\secrets\loader.go:98:20: fieldalignment: struct with 32 pointer bytes could be 24 (govet)
type SecretsLoader struct {
                   ^
internal\config\secrets\loader.go:106:1: Comment should end in a period (godot)
// NewSecretsLoader cria um novo loader de secrets
^
internal\config\secrets\loader.go:145:1: Comment should end in a period (godot)
// Load carrega todos os secrets
^
internal\config\secrets\loader.go:164:1: Comment should end in a period (godot)
// initVaultClient inicializa o cliente Vault
^
internal\config\secrets\loader.go:184:1: Comment should end in a period (godot)
// loadFromVault carrega secrets do Vault
^
internal\config\secrets\loader.go:206:1: Comment should end in a period (godot)
// loadFromK8s carrega secrets do Kubernetes
^
internal\config\secrets\loader.go:207:38: `(*SecretsLoader).loadFromK8s` - `ctx` is unused (unparam)
func (sl *SecretsLoader) loadFromK8s(ctx context.Context) (*SecretsConfig, error) {
                                     ^
internal\config\secrets\loader.go:207:38: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (sl *SecretsLoader) loadFromK8s(ctx context.Context) (*SecretsConfig, error) {
                                     ^
internal\config\secrets\loader.go:212:1: Comment should end in a period (godot)
// validateRequiredSecrets valida se todos os secrets obrigatórios estão presentes
^
internal\config\secrets\loader.go:212:76: `presentes` is a misspelling of `presents` (misspell)
// validateRequiredSecrets valida se todos os secrets obrigatórios estão presentes
                                                                           ^
internal\config\secrets\loader.go:230:1: Comment should end in a period (godot)
// GetDatabaseDSN retorna a DSN do banco de dados de forma segura
^
internal\config\secrets\loader.go:239:1: Comment should end in a period (godot)
// GetNATSConnection retorna a string de conexão NATS
^
internal\config\secrets\loader.go:252:1: Comment should end in a period (godot)
// Redact remove informações sensíveis para logs
^
internal\config\secrets\loader.go:254:5: emptyStringTest: replace `len(value) == 0` with `value == ""` (gocritic)
        if len(value) == 0 {
           ^
internal\config\secrets\loader.go:263:1: Comment should end in a period (godot)
// SecureString representa uma string segura que não aparece em logs
^
internal\config\secrets\loader.go:268:1: Comment should end in a period (godot)
// NewSecureString cria uma nova string segura
^
internal\config\secrets\loader.go:273:1: Comment should end in a period (godot)
// String implementa Stringer e redact o valor
^
internal\config\secrets\loader.go:283:1: Comment should end in a period (godot)
// MarshalJSON implementa json.Marshaler
^
internal\config\tls.go:13:6: exported: exported type TLSConfig should have comment or be unexported (revive)
type TLSConfig struct {
     ^
internal\config\tls.go:13:16: fieldalignment: struct of size 192 could be 176 (govet)
type TLSConfig struct {
               ^
internal\config\tls.go:34:6: exported: exported type TLSManager should have comment or be unexported (revive)
type TLSManager struct {
     ^
internal\config\tls.go:41:18: fieldalignment: struct with 80 pointer bytes could be 72 (govet)
type certWatcher struct {
                 ^
internal\config\tls.go:50:1: exported: exported function NewTLSManager should have comment or be unexported (revive)
func NewTLSManager(config *TLSConfig, logger *zap.Logger) (*TLSManager, error) {
^
internal\config\tls.go:96:16: G402: TLS MinVersion too low. (gosec)
        tlsConfig := &tls.Config{
                Certificates:             []tls.Certificate{cert},
                PreferServerCipherSuites: true,
                CurvePreferences: []tls.CurveID{
                        tls.X25519,
                        tls.CurveP256,
                        tls.CurveP384,
                        tls.CurveP521,
                },
        }
internal\config\tls.go:355:1: Comment should end in a period (godot)
// GetTLSConfig returns the current TLS configuration
^
internal\config\tls.go:366:1: Comment should end in a period (godot)
// IsEnabled returns whether TLS is enabled
^
internal\config\tls.go:371:1: Comment should end in a period (godot)
// Stop stops the certificate watcher
^
internal\config\tls.go:380:1: Comment should end in a period (godot)
// ValidateConfig validates the TLS configuration
^
internal\config\tls_test.go:341:45: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should handle multiple stops", func(t *testing.T) {
                                                   ^
internal\config\tls_test.go:351:1: Comment should end in a period (godot)
// Helper function to create temporary files for testing
^
internal\config\tls_test.go:370:1: Comment should end in a period (godot)
// Test certificate and key (for testing purposes only)
^
internal\constants\test_constants.go:1:1: package-comments: should have a package comment (revive)
package constants
^
internal\constants\test_constants.go:3:1: Comment should end in a period (godot)
// Non-sensitive test constants (not secrets)
^
internal\constants\test_constants.go:5:2: Comment should end in a period (godot)
        // JWT Testing Constants (non-secret)
        ^
internal\constants\test_constants.go:11:2: Comment should end in a period (godot)
        // Database Testing Constants (non-secret)
        ^
internal\constants\test_constants.go:16:1: Comment should end in a period (godot)
// Deprecated: Use GetTestSecret() for runtime-generated secrets instead
^
internal\constants\test_constants.go:29:1: Comment should end in a period (godot)
// TestCredentials provides a structured way to access test credentials
^
internal\constants\test_constants.go:38:1: Comment should end in a period (godot)
// GetTestCredentials returns test credentials for containerized testing
^
internal\constants\test_constants.go:51:1: Comment should end in a period (godot)
// IsTestEnvironment checks if we're in a test environment
^
internal\constants\test_secrets.go:31:1: Comment should end in a period (godot)
// generateRandomSecret creates a cryptographically random string of the specified byte length
^
internal\constants\test_secrets.go:41:1: Comment should end in a period (godot)
// ResetTestSecrets clears the cached secrets (useful for test isolation)
^
internal\dashboard\models.go:1:1: package-comments: should have a package comment (revive)
package dashboard
^
internal\dashboard\models.go:7:1: Comment should end in a period (godot)
// SystemOverview represents the overall system status
^
internal\dashboard\models.go:8:21: fieldalignment: struct with 176 pointer bytes could be 104 (govet)
type SystemOverview struct {
                    ^
internal\dashboard\models.go:16:1: Comment should end in a period (godot)
// SystemHealth represents overall system health status
^
internal\dashboard\models.go:17:19: fieldalignment: struct with 40 pointer bytes could be 16 (govet)
type SystemHealth struct {
                  ^
internal\dashboard\models.go:25:1: Comment should end in a period (godot)
// ComponentStatus represents individual component status
^
internal\dashboard\models.go:26:22: fieldalignment: struct with 128 pointer bytes could be 104 (govet)
type ComponentStatus struct {
                     ^
internal\dashboard\models.go:37:1: Comment should end in a period (godot)
// OverviewMetrics represents key system metrics
^
internal\dashboard\models.go:50:1: Comment should end in a period (godot)
// Alert represents system alerts
^
internal\dashboard\models.go:65:1: Comment should end in a period (godot)
// AlertType represents different types of alerts
^
internal\dashboard\models.go:69:2: exported: exported const AlertTypeSystem should have comment (or a comment on this block) or be unexported (revive)
        AlertTypeSystem      AlertType = "system"
        ^
internal\dashboard\models.go:76:1: Comment should end in a period (godot)
// AlertSeverity represents alert severity levels
^
internal\dashboard\models.go:80:2: exported: exported const AlertSeverityInfo should have comment (or a comment on this block) or be unexported (revive)
        AlertSeverityInfo      AlertSeverity = "info"
        ^
internal\dashboard\models.go:86:1: Comment should end in a period (godot)
// AlertStatus represents alert status
^
internal\dashboard\models.go:90:2: exported: exported const AlertStatusActive should have comment (or a comment on this block) or be unexported (revive)
        AlertStatusActive       AlertStatus = "active"
        ^
internal\dashboard\models.go:96:1: Comment should end in a period (godot)
// AlertAction represents available actions for alerts
^
internal\dashboard\models.go:104:1: Comment should end in a period (godot)
// RealtimeMetrics represents real-time system metrics
^
internal\dashboard\models.go:105:22: fieldalignment: struct with 624 pointer bytes could be 160 (govet)
type RealtimeMetrics struct {
                     ^
internal\dashboard\models.go:113:1: Comment should end in a period (godot)
// SystemMetrics represents system-level metrics
^
internal\dashboard\models.go:122:1: Comment should end in a period (godot)
// CPUMetrics represents CPU usage metrics
^
internal\dashboard\models.go:132:1: Comment should end in a period (godot)
// MemoryMetrics represents memory usage metrics
^
internal\dashboard\models.go:143:1: Comment should end in a period (godot)
// DiskMetrics represents disk usage metrics
^
internal\dashboard\models.go:154:1: Comment should end in a period (godot)
// NetworkMetrics represents network usage metrics
^
internal\dashboard\models.go:167:1: Comment should end in a period (godot)
// ProcessMetrics represents process-level metrics
^
internal\dashboard\models.go:176:1: Comment should end in a period (godot)
// PerformanceMetrics represents application performance metrics
^
internal\dashboard\models.go:187:1: Comment should end in a period (godot)
// ResponseTimeMetrics represents response time statistics
^
internal\dashboard\models.go:198:1: Comment should end in a period (godot)
// DatabaseMetrics represents database performance metrics
^
internal\dashboard\models.go:209:1: Comment should end in a period (godot)
// CacheMetricsData represents cache performance metrics
^
internal\dashboard\models.go:219:1: Comment should end in a period (godot)
// ErrorMetrics represents error tracking metrics
^
internal\dashboard\models.go:220:19: fieldalignment: struct with 48 pointer bytes could be 24 (govet)
type ErrorMetrics struct {
                  ^
internal\dashboard\models.go:229:1: Comment should end in a period (godot)
// RecentError represents recent error information
^
internal\dashboard\models.go:230:18: fieldalignment: struct with 120 pointer bytes could be 104 (govet)
type RecentError struct {
                 ^
internal\dashboard\models.go:241:1: Comment should end in a period (godot)
// TrafficMetrics represents traffic and usage metrics
^
internal\dashboard\models.go:242:21: fieldalignment: struct with 72 pointer bytes could be 48 (govet)
type TrafficMetrics struct {
                    ^
internal\dashboard\models.go:253:1: Comment should end in a period (godot)
// TrafficPeak represents peak traffic information
^
internal\dashboard\models.go:261:1: Comment should end in a period (godot)
// BandwidthMetrics represents bandwidth usage
^
internal\dashboard\models.go:270:1: Comment should end in a period (godot)
// ChartData represents time-series data for charts
^
internal\dashboard\models.go:276:1: Comment should end in a period (godot)
// Dataset represents a data series in a chart
^
internal\dashboard\models.go:277:14: fieldalignment: struct with 64 pointer bytes could be 56 (govet)
type Dataset struct {
             ^
internal\dashboard\models.go:285:1: Comment should end in a period (godot)
// DashboardWidget represents a dashboard widget configuration
^
internal\dashboard\models.go:286:6: exported: type name will be used as dashboard.DashboardWidget by other packages, and that stutters; consider calling this Widget (revive)
type DashboardWidget struct {
     ^
internal\dashboard\models.go:286:22: fieldalignment: struct with 104 pointer bytes could be 64 (govet)
type DashboardWidget struct {
                     ^
internal\dashboard\models.go:296:1: Comment should end in a period (godot)
// WidgetSize represents widget dimensions
^
internal\dashboard\models.go:302:1: Comment should end in a period (godot)
// WidgetPosition represents widget position
^
internal\dashboard\models.go:308:1: Comment should end in a period (godot)
// WebSocketMessage represents messages sent via WebSocket
^
internal\dashboard\models.go:316:1: Comment should end in a period (godot)
// SubscriptionRequest represents WebSocket subscription requests
^
internal\dashboard\models.go:317:26: fieldalignment: struct with 24 pointer bytes could be 16 (govet)
type SubscriptionRequest struct {
                         ^
internal\domain\models.go:1:1: package-comments: should have a package comment (revive)
package domain
^
internal\domain\models.go:9:1: Comment should end in a period (godot)
// Task represents a task in the system
^
internal\domain\models.go:10:11: fieldalignment: struct with 200 pointer bytes could be 152 (govet)
type Task struct {
          ^
internal\domain\models.go:26:1: Comment should end in a period (godot)
// TaskStatus represents the status of a task
^
internal\domain\models.go:30:2: exported: exported const TaskStatusPending should have comment (or a comment on this block) or be unexported (revive)
        TaskStatusPending    TaskStatus = "pending"
        ^
internal\domain\models.go:36:1: Comment should end in a period (godot)
// Priority represents task priority
^
internal\domain\models.go:40:2: exported: exported const PriorityLow should have comment (or a comment on this block) or be unexported (revive)
        PriorityLow    Priority = "low"
        ^
internal\domain\models.go:46:1: Comment should end in a period (godot)
// User represents a user in the system
^
internal\domain\models.go:47:11: fieldalignment: struct with 112 pointer bytes could be 88 (govet)
type User struct {
          ^
internal\domain\models.go:57:1: Comment should end in a period (godot)
// Role represents user role
^
internal\domain\models.go:61:2: exported: exported const RoleAdmin should have comment (or a comment on this block) or be unexported (revive)
        RoleAdmin Role = "admin"
        ^
internal\domain\models.go:65:1: Comment should end in a period (godot)
// Event represents a domain event
^
internal\domain\models.go:66:12: fieldalignment: struct with 80 pointer bytes could be 40 (govet)
type Event struct {
           ^
internal\domain\models.go:75:1: Comment should end in a period (godot)
// FeatureFlag represents a feature flag
^
internal\domain\models.go:76:18: fieldalignment: struct with 128 pointer bytes could be 112 (govet)
type FeatureFlag struct {
                 ^
internal\domain\models.go:87:1: Comment should end in a period (godot)
// TaskFilter represents filters for task queries
^
internal\domain\models.go:88:17: fieldalignment: struct with 104 pointer bytes could be 88 (govet)
type TaskFilter struct {
                ^
internal\domain\models.go:100:1: Comment should end in a period (godot)
// NewTask creates a new task with default values
^
internal\domain\models.go:116:1: Comment should end in a period (godot)
// Complete marks a task as completed
^
internal\domain\models.go:124:1: Comment should end in a period (godot)
// Cancel marks a task as cancelled
^
internal\domain\models.go:130:1: Comment should end in a period (godot)
// UpdateStatus updates task status
^
internal\domain\models.go:136:1: Comment should end in a period (godot)
// IsValidStatus checks if status transition is valid
^
internal\domain\repository.go:9:1: Comment should end in a period (godot)
// TaskRepository defines the interface for task data access
^
internal\domain\repository.go:20:1: Comment should end in a period (godot)
// UserRepository defines the interface for user data access
^
internal\domain\repository.go:30:1: Comment should end in a period (godot)
// EventRepository defines the interface for event data access
^
internal\domain\repository.go:37:1: Comment should end in a period (godot)
// FeatureFlagRepository defines the interface for feature flag data access
^
internal\domain\repository.go:46:1: Comment should end in a period (godot)
// CacheRepository defines the interface for cache operations
^
internal\dr\failover_stub.go:1:1: package-comments: package comment should be of the form "Package dr ..." (revive)
// internal/dr/failover_stub.go
^
internal\dr\failover_stub.go:8:6: exported: exported type Controller should have comment or be unexported (revive)
type Controller struct{}
     ^
internal\dr\failover_stub.go:10:1: exported: exported method Controller.Healthy should have comment or be unexported (revive)
func (c *Controller) Healthy() bool { return true }
^
internal\events\bus.go:1:1: package-comments: should have a package comment (revive)
package events
^
internal\events\nats_bus.go:15:1: Comment should end in a period (godot)
// NATSEventBus implements EventBus using NATS
^
internal\events\nats_bus.go:21:1: Comment should end in a period (godot)
// NewNATSEventBus creates a new NATS event bus
^
internal\events\nats_bus.go:27:34: unused-parameter: parameter 'nc' seems to be unused, consider removing or renaming it as _ (revive)
                nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
                                               ^
internal\events\nats_bus.go:49:1: Comment should end in a period (godot)
// Publish publishes an event to NATS
^
internal\events\nats_bus.go:50:34: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (bus *NATSEventBus) Publish(ctx context.Context, event *domain.Event) error {
                                 ^
internal\events\nats_bus.go:84:1: Comment should end in a period (godot)
// Subscribe subscribes to events of a specific type
^
internal\events\nats_bus.go:114:1: Comment should end in a period (godot)
// SubscribeQueue subscribes to events with queue group
^
internal\events\nats_bus.go:146:1: Comment should end in a period (godot)
// Close closes the NATS connection
^
internal\events\nats_bus.go:154:1: Comment should end in a period (godot)
// EventHandler defines the interface for event handlers
^
internal\events\nats_bus.go:159:1: Comment should end in a period (godot)
// EventHandlerFunc is an adapter to allow using regular functions as EventHandler
^
internal\events\nats_bus.go:162:1: Comment should end in a period (godot)
// Handle implements EventHandler interface
^
internal\events\nats_bus.go:167:1: Comment should end in a period (godot)
// TaskEventHandler handles task-related events
^
internal\events\nats_bus.go:172:1: Comment should end in a period (godot)
// NewTaskEventHandler creates a new task event handler
^
internal\events\nats_bus.go:177:1: Comment should end in a period (godot)
// Handle handles task events
^
internal\events\nats_bus.go:194:46: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (h *TaskEventHandler) handleTaskCreated(ctx context.Context, event *domain.Event) error {
                                             ^
internal\events\nats_bus.go:194:46: `(*TaskEventHandler).handleTaskCreated` - `ctx` is unused (unparam)
func (h *TaskEventHandler) handleTaskCreated(ctx context.Context, event *domain.Event) error {
                                             ^
internal\events\nats_bus.go:203:46: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (h *TaskEventHandler) handleTaskUpdated(ctx context.Context, event *domain.Event) error {
                                             ^
internal\events\nats_bus.go:203:46: `(*TaskEventHandler).handleTaskUpdated` - `ctx` is unused (unparam)
func (h *TaskEventHandler) handleTaskUpdated(ctx context.Context, event *domain.Event) error {
                                             ^
internal\events\nats_bus.go:212:48: `(*TaskEventHandler).handleTaskCompleted` - `ctx` is unused (unparam)
func (h *TaskEventHandler) handleTaskCompleted(ctx context.Context, event *domain.Event) error {
                                               ^
internal\events\nats_bus.go:212:48: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (h *TaskEventHandler) handleTaskCompleted(ctx context.Context, event *domain.Event) error {
                                               ^
internal\events\nats_bus.go:221:46: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (h *TaskEventHandler) handleTaskDeleted(ctx context.Context, event *domain.Event) error {
                                             ^
internal\events\nats_bus.go:221:46: `(*TaskEventHandler).handleTaskDeleted` - `ctx` is unused (unparam)
func (h *TaskEventHandler) handleTaskDeleted(ctx context.Context, event *domain.Event) error {
                                             ^
internal\features\context.go:5:18: fieldalignment: struct with 56 pointer bytes could be 48 (govet)
type UserContext struct {
                 ^
internal\features\flags.go:1:1: package-comments: package comment should be of the form "Package features ..." (revive)
// internal/features/flags.go
^
internal\features\flags.go:11:6: exported: exported type FlagType should have comment or be unexported (revive)
type FlagType string
     ^
internal\features\flags.go:14:2: exported: exported const Boolean should have comment (or a comment on this block) or be unexported (revive)
        Boolean    FlagType = "boolean"
        ^
internal\features\flags.go:20:6: exported: exported type Flag should have comment or be unexported (revive)
type Flag struct {
     ^
internal\features\flags.go:20:11: fieldalignment: struct with 72 pointer bytes could be 56 (govet)
type Flag struct {
          ^
internal\features\flags.go:30:6: exported: exported type EvalContext should have comment or be unexported (revive)
type EvalContext struct {
     ^
internal\features\flags.go:30:18: fieldalignment: struct with 48 pointer bytes could be 40 (govet)
type EvalContext struct {
                 ^
internal\features\flags.go:36:6: exported: exported type Manager should have comment or be unexported (revive)
type Manager interface {
     ^
internal\features\flags.go:40:6: exported: exported type InMemoryManager should have comment or be unexported (revive)
type InMemoryManager struct {
     ^
internal\features\flags.go:44:1: exported: exported function NewInMemoryManager should have comment or be unexported (revive)
func NewInMemoryManager() *InMemoryManager {
^
internal\features\flags.go:58:1: exported: exported method InMemoryManager.Evaluate should have comment or be unexported (revive)
func (m *InMemoryManager) Evaluate(key string, ctx EvalContext) any {
^
internal\features\manager.go:15:1: Comment should end in a period (godot)
// FlagManager manages feature flags with persistence
^
internal\features\manager.go:16:18: fieldalignment: struct with 88 pointer bytes could be 64 (govet)
type FlagManager struct {
                 ^
internal\features\manager.go:26:1: Comment should end in a period (godot)
// NewFlagManager creates a new feature flag manager
^
internal\features\manager.go:42:1: Comment should end in a period (godot)
// IsEnabled checks if a feature flag is enabled
^
internal\features\manager.go:53:1: Comment should end in a period (godot)
// IsEnabledWithDefault checks if a feature flag is enabled with a default value
^
internal\features\manager.go:63:1: Comment should end in a period (godot)
// GetFlag retrieves a feature flag
^
internal\features\manager.go:105:1: Comment should end in a period (godot)
// SetFlag creates or updates a feature flag
^
internal\features\manager.go:146:1: Comment should end in a period (godot)
// ListFlags returns all feature flags
^
internal\features\manager.go:151:1: Comment should end in a period (godot)
// DeleteFlag deletes a feature flag
^
internal\features\manager.go:172:1: Comment should end in a period (godot)
// RefreshFlags reloads all flags from the repository
^
internal\features\manager.go:201:1: Comment should end in a period (godot)
// startRefresh starts background refresh of feature flags
^
internal\features\manager.go:219:1: Comment should end in a period (godot)
// Stop stops the background refresh
^
internal\features\manager.go:227:1: Comment should end in a period (godot)
// EvaluateFlag evaluates a feature flag with strategy
^
internal\features\manager.go:228:1: paramTypeCombine: func(ctx context.Context, key string, userID string, attributes map[string]interface{}) bool could be replaced with func(ctx context.Context, key, userID string, attributes map[string]interface{}) bool (gocritic)
func (m *FlagManager) EvaluateFlag(ctx context.Context, key string, userID string, attributes map[string]interface{}) bool {
^
internal\features\manager.go:255:1: Comment should end in a period (godot)
// evaluatePercentage evaluates percentage-based rollout
^
internal\features\manager.go:268:1: Comment should end in a period (godot)
// evaluateUserList evaluates user list strategy
^
internal\features\manager.go:282:1: Comment should end in a period (godot)
// evaluateAttribute evaluates attribute-based strategy
^
internal\features\manager_test.go:14:1: Comment should end in a period (godot)
// Mock repositories
^
internal\handlers\health.go:1:1: package-comments: should have a package comment (revive)
package handlers
^
internal\handlers\health.go:8:6: exported: exported type HealthHandler should have comment or be unexported (revive)
type HealthHandler struct{}
     ^
internal\handlers\health.go:10:1: exported: exported function NewHealthHandler should have comment or be unexported (revive)
func NewHealthHandler() *HealthHandler {
^
internal\handlers\health.go:14:1: exported: exported method HealthHandler.Live should have comment or be unexported (revive)
func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
^
internal\handlers\health.go:14:53: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
                                                    ^
internal\handlers\health.go:17:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "alive"})
                                 ^
internal\handlers\health.go:20:1: exported: exported method HealthHandler.Ready should have comment or be unexported (revive)
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
^
internal\handlers\health.go:20:54: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
                                                     ^
internal\handlers\health.go:23:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
                                 ^
internal\handlers\health.go:26:1: exported: exported method HealthHandler.Health should have comment or be unexported (revive)
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
^
internal\handlers\health.go:26:55: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
                                                      ^
internal\handlers\health.go:29:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
                                 ^
internal\handlers\health.go:32:1: exported: exported method HealthHandler.Livez should have comment or be unexported (revive)
func (h *HealthHandler) Livez(w http.ResponseWriter, r *http.Request) {
^
internal\handlers\health.go:36:1: exported: exported method HealthHandler.Readyz should have comment or be unexported (revive)
func (h *HealthHandler) Readyz(w http.ResponseWriter, r *http.Request) {
^
internal\handlers\health.go:40:1: exported: exported method HealthHandler.Metrics should have comment or be unexported (revive)
func (h *HealthHandler) Metrics() http.Handler {
^
internal\handlers\health.go:41:54: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                                                            ^
internal\handlers\health.go:44:10: Error return value of `w.Write` is not checked (errcheck)
                w.Write([]byte("# Metrics placeholder\n"))
                       ^
internal\handlers\health_test.go:11:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/livez", nil)
               ^
internal\handlers\health_test.go:28:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
               ^
internal\handlers\health_test.go:40:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/health", nil)
               ^
internal\handlers\health_test.go:52:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
               ^
internal\handlers\health_test.go:69:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/livez", nil)
               ^
internal\handlers\health_test.go:86:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
               ^
internal\http\router.go:1:1: package-comments: package comment should be of the form "Package httpserver ..." (revive)
// internal/http/router.go
^
internal\http\router.go:11:1: exported: exported function RegisterRoutes should have comment or be unexported (revive)
func RegisterRoutes(mux *http.ServeMux) {
^
internal\http\router.go:16:35: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func hello(w http.ResponseWriter, r *http.Request) {
                                  ^
internal\http\router.go:19:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(resp)
                                 ^
internal\http\router.go:24:18: fieldalignment: struct with 40 pointer bytes could be 32 (govet)
type evalRequest struct {
                 ^
internal\http\router.go:38:27: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
        json.NewEncoder(w).Encode(map[string]any{"flag": req.Flag, "value": val})
                                 ^
internal\lifecycle\deployment.go:1:1: package-comments: should have a package comment (revive)
package lifecycle
^
internal\lifecycle\deployment.go:13:1: Comment should end in a period (godot)
// DeploymentStrategy represents deployment strategies
^
internal\lifecycle\deployment.go:17:2: exported: exported const DeploymentBlueGreen should have comment (or a comment on this block) or be unexported (revive)
        DeploymentBlueGreen DeploymentStrategy = "blue_green"
        ^
internal\lifecycle\deployment.go:23:1: Comment should end in a period (godot)
// DeploymentPhase represents deployment phases
^
internal\lifecycle\deployment.go:27:2: exported: exported const PhaseValidation should have comment (or a comment on this block) or be unexported (revive)
        PhaseValidation   DeploymentPhase = "validation"
        ^
internal\lifecycle\deployment.go:36:1: Comment should end in a period (godot)
// DeploymentConfig configures deployment automation
^
internal\lifecycle\deployment.go:37:23: fieldalignment: struct of size 344 could be 328 (govet)
type DeploymentConfig struct {
                      ^
internal\lifecycle\deployment.go:85:1: Comment should end in a period (godot)
// DeploymentHook represents a deployment hook
^
internal\lifecycle\deployment.go:86:21: fieldalignment: struct with 104 pointer bytes could be 80 (govet)
type DeploymentHook struct {
                    ^
internal\lifecycle\deployment.go:97:1: Comment should end in a period (godot)
// RollbackThresholds defines when to trigger auto-rollback
^
internal\lifecycle\deployment.go:105:1: Comment should end in a period (godot)
// DeploymentResult represents the result of a deployment
^
internal\lifecycle\deployment.go:106:23: fieldalignment: struct with 216 pointer bytes could be 184 (govet)
type DeploymentResult struct {
                      ^
internal\lifecycle\deployment.go:122:1: Comment should end in a period (godot)
// DeploymentAutomation manages automated deployments
^
internal\lifecycle\deployment.go:123:27: fieldalignment: struct with 368 pointer bytes could be 336 (govet)
type DeploymentAutomation struct {
                          ^
internal\lifecycle\deployment.go:133:1: Comment should end in a period (godot)
// NewDeploymentAutomation creates a new deployment automation system
^
internal\lifecycle\deployment.go:134:30: hugeParam: config is heavy (344 bytes); consider passing it by pointer (gocritic)
func NewDeploymentAutomation(config DeploymentConfig, logger logger.Logger) *DeploymentAutomation {
                             ^
internal\lifecycle\deployment.go:143:1: Comment should end in a period (godot)
// Deploy executes a deployment using the configured strategy
^
internal\lifecycle\deployment.go:194:1: Comment should end in a period (godot)
// Rollback rolls back to the previous version
^
internal\lifecycle\deployment.go:232:1: Comment should end in a period (godot)
// GetDeploymentHistory returns deployment history
^
internal\lifecycle\deployment.go:239:1: Comment should end in a period (godot)
// GetCurrentDeployment returns the current deployment status
^
internal\lifecycle\deployment.go:247:16: fieldalignment: struct with 24 pointer bytes could be 16 (govet)
        pipeline := []struct {
                      ^
internal\lifecycle\deployment.go:270:52: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) validateDeployment(ctx context.Context, result *DeploymentResult) error {
                                                   ^
internal\lifecycle\deployment.go:407:20: Error return value of `da.executeCommand` is not checked (errcheck)
                da.executeCommand(ctx, fmt.Sprintf("kubectl delete deployment mcp-ultra-canary --namespace=%s", da.config.Namespace), result)
                                 ^
internal\lifecycle\deployment.go:412:1: The line is 185 characters long, which exceeds the maximum of 140 characters. (lll)
        cmd = fmt.Sprintf("kubectl patch deployment mcp-ultra --patch '{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"mcp-ultra\",\"image\":\"%s:%s\"}]}}}}' --namespace=%s",
^
internal\lifecycle\deployment.go:420:19: Error return value of `da.executeCommand` is not checked (errcheck)
        da.executeCommand(ctx, fmt.Sprintf("kubectl delete deployment mcp-ultra-canary --namespace=%s", da.config.Namespace), result)
                         ^
internal\lifecycle\deployment.go:537:9: G204: Subprocess launched with a potential tainted input or cmd arguments (gosec)
        cmd := exec.CommandContext(ctx, parts[0], parts[1:]...)
               ^
internal\lifecycle\deployment.go:563:20: S1039: unnecessary use of fmt.Sprintf (gosimple)
        da.addLog(result, fmt.Sprintf("Script executed successfully"))
                          ^
internal\lifecycle\deployment.go:567:49: `(*DeploymentAutomation).executeHTTPHook` - `ctx` is unused (unparam)
func (da *DeploymentAutomation) executeHTTPHook(ctx context.Context, hook DeploymentHook, result *DeploymentResult) error {
                                                ^
internal\lifecycle\deployment.go:567:49: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) executeHTTPHook(ctx context.Context, hook DeploymentHook, result *DeploymentResult) error {
                                                ^
internal\lifecycle\deployment.go:567:117: (*DeploymentAutomation).executeHTTPHook - result 0 (error) is always nil (unparam)
func (da *DeploymentAutomation) executeHTTPHook(ctx context.Context, hook DeploymentHook, result *DeploymentResult) error {
                                                                                                                    ^
internal\lifecycle\deployment.go:579:53: unused-parameter: parameter 'version' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) validateDockerImage(version string) error {
                                                    ^
internal\lifecycle\deployment.go:589:55: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) validateCanaryMetrics(ctx context.Context, result *DeploymentResult) error {
                                                      ^
internal\lifecycle\deployment.go:589:76: unused-parameter: parameter 'result' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) validateCanaryMetrics(ctx context.Context, result *DeploymentResult) error {
                                                                           ^
internal\lifecycle\deployment.go:594:53: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) performHealthChecks(ctx context.Context, result *DeploymentResult) error {
                                                    ^
internal\lifecycle\deployment.go:594:74: unused-parameter: parameter 'result' seems to be unused, consider removing or renaming it as _ (revive)
func (da *DeploymentAutomation) performHealthChecks(ctx context.Context, result *DeploymentResult) error {
                                                                         ^
internal\lifecycle\health.go:14:1: Comment should end in a period (godot)
// HealthStatus represents the health status of a component
^
internal\lifecycle\health.go:18:2: exported: exported const HealthStatusHealthy should have comment (or a comment on this block) or be unexported (revive)
        HealthStatusHealthy   HealthStatus = "healthy"
        ^
internal\lifecycle\health.go:24:1: Comment should end in a period (godot)
// HealthCheck represents a health check result
^
internal\lifecycle\health.go:25:18: fieldalignment: struct with 96 pointer bytes could be 88 (govet)
type HealthCheck struct {
                 ^
internal\lifecycle\health.go:35:1: Comment should end in a period (godot)
// HealthReport represents the overall health status
^
internal\lifecycle\health.go:36:19: fieldalignment: struct with 120 pointer bytes could be 72 (govet)
type HealthReport struct {
                  ^
internal\lifecycle\health.go:46:1: Comment should end in a period (godot)
// HealthSummary provides a summary of health checks
^
internal\lifecycle\health.go:55:1: Comment should end in a period (godot)
// DependencyStatus represents the status of an external dependency
^
internal\lifecycle\health.go:56:23: fieldalignment: struct with 88 pointer bytes could be 72 (govet)
type DependencyStatus struct {
                      ^
internal\lifecycle\health.go:65:1: Comment should end in a period (godot)
// HealthChecker interface for health check implementations
^
internal\lifecycle\health.go:73:1: Comment should end in a period (godot)
// HealthMonitor provides comprehensive health monitoring
^
internal\lifecycle\health.go:74:20: fieldalignment: struct with 272 pointer bytes could be 232 (govet)
type HealthMonitor struct {
                   ^
internal\lifecycle\health.go:94:1: Comment should end in a period (godot)
// HealthConfig configures health monitoring
^
internal\lifecycle\health.go:95:19: fieldalignment: struct of size 128 could be 112 (govet)
type HealthConfig struct {
                  ^
internal\lifecycle\health.go:119:1: Comment should end in a period (godot)
// DependencyChecker checks external dependencies
^
internal\lifecycle\health.go:127:1: Comment should end in a period (godot)
// DefaultHealthConfig returns default health monitoring configuration
^
internal\lifecycle\health.go:133:27: commentedOutCode: may want to remove commented-out code (gocritic)
                DegradedThreshold:  25, // 25% failures = degraded
                                        ^
internal\lifecycle\health.go:134:27: commentedOutCode: may want to remove commented-out code (gocritic)
                UnhealthyThreshold: 50, // 50% failures = unhealthy
                                        ^
internal\lifecycle\health.go:146:1: Comment should end in a period (godot)
// NewHealthMonitor creates a new health monitor
^
internal\lifecycle\health.go:159:1: Comment should end in a period (godot)
// RegisterChecker registers a health checker
^
internal\lifecycle\health.go:172:1: Comment should end in a period (godot)
// RegisterDependency registers a dependency checker
^
internal\lifecycle\health.go:185:1: Comment should end in a period (godot)
// Start starts the health monitoring
^
internal\lifecycle\health.go:214:1: Comment should end in a period (godot)
// Stop stops the health monitoring
^
internal\lifecycle\health.go:234:1: Comment should end in a period (godot)
// GetHealth returns the current health status
^
internal\lifecycle\health.go:239:1: Comment should end in a period (godot)
// GetLastReport returns the last health report
^
internal\lifecycle\health.go:253:1: Comment should end in a period (godot)
// IsHealthy returns true if the system is healthy
^
internal\lifecycle\health.go:262:1: Comment should end in a period (godot)
// IsDegraded returns true if the system is degraded
^
internal\lifecycle\health.go:271:1: Comment should end in a period (godot)
// IsUnhealthy returns true if the system is unhealthy
^
internal\lifecycle\health.go:432:2: ifElseChain: rewrite if-else to switch statement (gocritic)
        if failures == 0 {
        ^
internal\lifecycle\health.go:476:28: Error return value of `(*encoding/json.Encoder).Encode` is not checked (errcheck)
                json.NewEncoder(w).Encode(report)
                                         ^
internal\lifecycle\health.go:480:55: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
        mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
                                                             ^
internal\lifecycle\health.go:483:11: Error return value of `w.Write` is not checked (errcheck)
                        w.Write([]byte("OK"))
                               ^
internal\lifecycle\health.go:486:11: Error return value of `w.Write` is not checked (errcheck)
                        w.Write([]byte("Not Ready"))
                               ^
internal\lifecycle\health.go:491:54: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
        mux.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
                                                            ^
internal\lifecycle\health.go:494:11: Error return value of `w.Write` is not checked (errcheck)
                        w.Write([]byte("OK"))
                               ^
internal\lifecycle\health.go:497:11: Error return value of `w.Write` is not checked (errcheck)
                        w.Write([]byte("Unhealthy"))
                               ^
internal\lifecycle\health.go:532:1: Comment should end in a period (godot)
// DatabaseHealthChecker checks database connectivity
^
internal\lifecycle\health.go:540:1: exported: exported function NewDatabaseHealthChecker should have comment or be unexported (revive)
func NewDatabaseHealthChecker(name string) *DatabaseHealthChecker {
^
internal\lifecycle\health.go:548:1: exported: exported method DatabaseHealthChecker.Name should have comment or be unexported (revive)
func (d *DatabaseHealthChecker) Name() string {
^
internal\lifecycle\health.go:552:1: exported: exported method DatabaseHealthChecker.IsRequired should have comment or be unexported (revive)
func (d *DatabaseHealthChecker) IsRequired() bool {
^
internal\lifecycle\health.go:556:1: exported: exported method DatabaseHealthChecker.Timeout should have comment or be unexported (revive)
func (d *DatabaseHealthChecker) Timeout() time.Duration {
^
internal\lifecycle\health.go:560:1: exported: exported method DatabaseHealthChecker.Check should have comment or be unexported (revive)
func (d *DatabaseHealthChecker) Check(ctx context.Context) HealthCheck {
^
internal\lifecycle\health.go:560:39: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (d *DatabaseHealthChecker) Check(ctx context.Context) HealthCheck {
                                      ^
internal\lifecycle\health.go:580:1: Comment should end in a period (godot)
// RedisHealthChecker checks Redis connectivity
^
internal\lifecycle\health.go:588:1: exported: exported function NewRedisHealthChecker should have comment or be unexported (revive)
func NewRedisHealthChecker(name string) *RedisHealthChecker {
^
internal\lifecycle\health.go:596:1: exported: exported method RedisHealthChecker.Name should have comment or be unexported (revive)
func (r *RedisHealthChecker) Name() string {
^
internal\lifecycle\health.go:600:1: exported: exported method RedisHealthChecker.IsRequired should have comment or be unexported (revive)
func (r *RedisHealthChecker) IsRequired() bool {
^
internal\lifecycle\health.go:604:1: exported: exported method RedisHealthChecker.Timeout should have comment or be unexported (revive)
func (r *RedisHealthChecker) Timeout() time.Duration {
^
internal\lifecycle\health.go:608:1: exported: exported method RedisHealthChecker.Check should have comment or be unexported (revive)
func (r *RedisHealthChecker) Check(ctx context.Context) HealthCheck {
^
internal\lifecycle\health.go:608:36: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (r *RedisHealthChecker) Check(ctx context.Context) HealthCheck {
                                   ^
internal\lifecycle\manager.go:14:1: Comment should end in a period (godot)
// LifecycleState represents the current state of the application
^
internal\lifecycle\manager.go:15:6: exported: type name will be used as lifecycle.LifecycleState by other packages, and that stutters; consider calling this State (revive)
type LifecycleState int32
     ^
internal\lifecycle\manager.go:18:2: exported: exported const StateInitializing should have comment (or a comment on this block) or be unexported (revive)
        StateInitializing LifecycleState = iota
        ^
internal\lifecycle\manager.go:37:10: string `healthy` has 3 occurrences, but such constant `HealthStatusHealthy` already exists (goconst)
                return "healthy"
                       ^
internal\lifecycle\manager.go:51:1: Comment should end in a period (godot)
// Component represents a lifecycle-managed component
^
internal\lifecycle\manager.go:62:1: Comment should end in a period (godot)
// LifecycleEvent represents events during lifecycle transitions
^
internal\lifecycle\manager.go:63:6: exported: type name will be used as lifecycle.LifecycleEvent by other packages, and that stutters; consider calling this Event (revive)
type LifecycleEvent struct {
     ^
internal\lifecycle\manager.go:63:21: fieldalignment: struct with 104 pointer bytes could be 88 (govet)
type LifecycleEvent struct {
                    ^
internal\lifecycle\manager.go:73:1: Comment should end in a period (godot)
// LifecycleManager manages application lifecycle and component orchestration
^
internal\lifecycle\manager.go:74:6: exported: type name will be used as lifecycle.LifecycleManager by other packages, and that stutters; consider calling this Manager (revive)
type LifecycleManager struct {
     ^
internal\lifecycle\manager.go:74:23: fieldalignment: struct with 312 pointer bytes could be 176 (govet)
type LifecycleManager struct {
                      ^
internal\lifecycle\manager.go:107:1: Comment should end in a period (godot)
// ComponentState tracks individual component state
^
internal\lifecycle\manager.go:108:21: fieldalignment: struct with 96 pointer bytes could be 80 (govet)
type ComponentState struct {
                    ^
internal\lifecycle\manager.go:117:1: Comment should end in a period (godot)
// Config configures the lifecycle manager
^
internal\lifecycle\manager.go:130:1: Comment should end in a period (godot)
// DefaultConfig returns default lifecycle manager configuration
^
internal\lifecycle\manager.go:145:1: Comment should end in a period (godot)
// NewLifecycleManager creates a new lifecycle manager
^
internal\lifecycle\manager.go:179:1: Comment should end in a period (godot)
// RegisterComponent registers a component for lifecycle management
^
internal\lifecycle\manager.go:204:1: Comment should end in a period (godot)
// RegisterEventHandler registers an event handler for lifecycle events
^
internal\lifecycle\manager.go:212:1: Comment should end in a period (godot)
// Start starts all registered components in priority order
^
internal\lifecycle\manager.go:259:1: Comment should end in a period (godot)
// Stop stops all components in reverse priority order
^
internal\lifecycle\manager.go:304:1: Comment should end in a period (godot)
// GetState returns the current lifecycle state
^
internal\lifecycle\manager.go:309:1: Comment should end in a period (godot)
// IsReady returns true if the application is ready to serve requests
^
internal\lifecycle\manager.go:315:1: Comment should end in a period (godot)
// IsHealthy returns true if the application is healthy
^
internal\lifecycle\manager.go:320:1: Comment should end in a period (godot)
// GetComponentStates returns the current state of all components
^
internal\lifecycle\manager.go:332:1: Comment should end in a period (godot)
// GetEventHistory returns recent lifecycle events
^
internal\lifecycle\manager.go:349:1: Comment should end in a period (godot)
// GetMetrics returns lifecycle metrics
^
internal\lifecycle\manager.go:386:1: Comment should end in a period (godot)
// LifecycleMetrics contains lifecycle metrics
^
internal\lifecycle\manager.go:387:6: exported: type name will be used as lifecycle.LifecycleMetrics by other packages, and that stutters; consider calling this Metrics (revive)
type LifecycleMetrics struct {
     ^
internal\lifecycle\manager.go:518:1: The line is 150 characters long, which exceeds the maximum of 140 characters. (lll)
func (lm *LifecycleManager) emitEvent(eventType, component string, state LifecycleState, message string, metadata map[string]interface{}, err error) {
^
internal\lifecycle\manager.go:621:2: ifElseChain: rewrite if-else to switch statement (gocritic)
        if errorCount == 0 && healthyCount == totalComponents {
        ^
internal\lifecycle\operations.go:12:1: Comment should end in a period (godot)
// OperationType represents different types of operations
^
internal\lifecycle\operations.go:16:2: exported: exported const OperationMaintenance should have comment (or a comment on this block) or be unexported (revive)
        OperationMaintenance   OperationType = "maintenance"
        ^
internal\lifecycle\operations.go:27:1: Comment should end in a period (godot)
// OperationStatus represents the status of an operation
^
internal\lifecycle\operations.go:31:2: exported: exported const StatusPending should have comment (or a comment on this block) or be unexported (revive)
        StatusPending   OperationStatus = "pending"
        ^
internal\lifecycle\operations.go:38:1: Comment should end in a period (godot)
// Operation represents a system operation
^
internal\lifecycle\operations.go:39:16: fieldalignment: struct with 304 pointer bytes could be 248 (govet)
type Operation struct {
               ^
internal\lifecycle\operations.go:77:1: Comment should end in a period (godot)
// OperationStep represents a step within an operation
^
internal\lifecycle\operations.go:78:20: fieldalignment: struct with 112 pointer bytes could be 96 (govet)
type OperationStep struct {
                   ^
internal\lifecycle\operations.go:93:1: Comment should end in a period (godot)
// OperationExecutor defines the interface for operation execution
^
internal\lifecycle\operations.go:100:1: Comment should end in a period (godot)
// OperationsManager manages system operations and procedures
^
internal\lifecycle\operations.go:101:24: fieldalignment: struct with 160 pointer bytes could be 48 (govet)
type OperationsManager struct {
                       ^
internal\lifecycle\operations.go:123:1: Comment should end in a period (godot)
// OperationsConfig configures operations management
^
internal\lifecycle\operations.go:135:1: Comment should end in a period (godot)
// DefaultOperationsConfig returns default operations configuration
^
internal\lifecycle\operations.go:149:1: Comment should end in a period (godot)
// NewOperationsManager creates a new operations manager
^
internal\lifecycle\operations.go:164:1: Comment should end in a period (godot)
// RegisterExecutor registers an operation executor
^
internal\lifecycle\operations.go:173:1: Comment should end in a period (godot)
// Start starts the operations manager
^
internal\lifecycle\operations.go:197:1: Comment should end in a period (godot)
// Stop stops the operations manager
^
internal\lifecycle\operations.go:220:1: Comment should end in a period (godot)
// CreateOperation creates a new operation
^
internal\lifecycle\operations.go:282:1: Comment should end in a period (godot)
// ExecuteOperation executes an operation asynchronously
^
internal\lifecycle\operations.go:306:1: Comment should end in a period (godot)
// CancelOperation cancels a running operation
^
internal\lifecycle\operations.go:340:1: Comment should end in a period (godot)
// GetOperation returns an operation by ID
^
internal\lifecycle\operations.go:355:1: Comment should end in a period (godot)
// ListOperations returns all operations with optional filtering
^
internal\lifecycle\operations.go:372:1: Comment should end in a period (godot)
// GetOperationHistory returns operation history
^
internal\lifecycle\operations.go:389:1: Comment should end in a period (godot)
// OperationFilter for filtering operations
^
internal\lifecycle\operations.go:398:1: Comment should end in a period (godot)
// Matches checks if an operation matches the filter
^
internal\lifecycle\operations.go:584:1: Comment should end in a period (godot)
// MaintenanceExecutor handles maintenance operations
^
internal\lifecycle\operations.go:589:1: exported: exported function NewMaintenanceExecutor should have comment or be unexported (revive)
func NewMaintenanceExecutor(logger logger.Logger) *MaintenanceExecutor {
^
internal\lifecycle\operations.go:593:1: exported: exported method MaintenanceExecutor.Execute should have comment or be unexported (revive)
func (me *MaintenanceExecutor) Execute(ctx context.Context, operation *Operation) error {
^
internal\lifecycle\operations.go:619:1: exported: exported method MaintenanceExecutor.Rollback should have comment or be unexported (revive)
func (me *MaintenanceExecutor) Rollback(ctx context.Context, operation *Operation) error {
^
internal\lifecycle\operations.go:619:41: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (me *MaintenanceExecutor) Rollback(ctx context.Context, operation *Operation) error {
                                        ^
internal\lifecycle\operations.go:624:1: exported: exported method MaintenanceExecutor.Validate should have comment or be unexported (revive)
func (me *MaintenanceExecutor) Validate(operation *Operation) error {
^
internal\metrics\business.go:1:1: package-comments: should have a package comment (revive)
package metrics
^
internal\metrics\business.go:13:1: Comment should end in a period (godot)
// MetricType represents different types of business metrics
^
internal\metrics\business.go:17:2: exported: exported const MetricCounter should have comment (or a comment on this block) or be unexported (revive)
        MetricCounter   MetricType = "counter"
        ^
internal\metrics\business.go:23:1: Comment should end in a period (godot)
// AggregationType represents how metrics should be aggregated
^
internal\metrics\business.go:27:2: exported: exported const AggregationSum should have comment (or a comment on this block) or be unexported (revive)
        AggregationSum   AggregationType = "sum"
        ^
internal\metrics\business.go:36:1: Comment should end in a period (godot)
// BusinessMetric defines a business metric configuration
^
internal\metrics\business.go:37:21: fieldalignment: struct with 160 pointer bytes could be 144 (govet)
type BusinessMetric struct {
                    ^
internal\metrics\business.go:52:1: Comment should end in a period (godot)
// BusinessMetricsConfig configures business metrics collection
^
internal\metrics\business.go:53:28: fieldalignment: struct of size 160 could be 144 (govet)
type BusinessMetricsConfig struct {
                           ^
internal\metrics\business.go:75:1: Comment should end in a period (godot)
// MetricAlertRule defines alerting rules for business metrics
^
internal\metrics\business.go:76:22: fieldalignment: struct with 80 pointer bytes could be 56 (govet)
type MetricAlertRule struct {
                     ^
internal\metrics\business.go:87:1: Comment should end in a period (godot)
// MetricValue represents a metric measurement
^
internal\metrics\business.go:88:18: fieldalignment: struct with 64 pointer bytes could be 56 (govet)
type MetricValue struct {
                 ^
internal\metrics\business.go:96:1: Comment should end in a period (godot)
// AggregatedMetric represents an aggregated metric value
^
internal\metrics\business.go:104:1: Comment should end in a period (godot)
// BusinessMetricsCollector collects and manages business metrics
^
internal\metrics\business.go:105:31: fieldalignment: struct with 288 pointer bytes could be 240 (govet)
type BusinessMetricsCollector struct {
                              ^
internal\metrics\business.go:126:1: Comment should end in a period (godot)
// AlertState tracks the state of metric alerts
^
internal\metrics\business.go:127:17: fieldalignment: struct with 104 pointer bytes could be 80 (govet)
type AlertState struct {
                ^
internal\metrics\business.go:138:1: Comment should end in a period (godot)
// MetricStorage interface for metric storage backends
^
internal\metrics\business.go:147:1: Comment should end in a period (godot)
// MetricQuery defines a metric query
^
internal\metrics\business.go:148:18: fieldalignment: struct with 72 pointer bytes could be 64 (govet)
type MetricQuery struct {
                 ^
internal\metrics\business.go:156:1: Comment should end in a period (godot)
// AggregationQuery defines an aggregation query
^
internal\metrics\business.go:164:1: Comment should end in a period (godot)
// DefaultBusinessMetricsConfig returns default configuration
^
internal\metrics\business.go:185:1: Comment should end in a period (godot)
// DefaultBusinessMetrics returns default business metrics
^
internal\metrics\business.go:312:1: Comment should end in a period (godot)
// DefaultAlertRules returns default alert rules
^
internal\metrics\business.go:363:1: Comment should end in a period (godot)
// NewBusinessMetricsCollector creates a new business metrics collector
^
internal\metrics\business.go:411:1: Comment should end in a period (godot)
// RecordCounter records a counter metric
^
internal\metrics\business.go:416:1: Comment should end in a period (godot)
// RecordGauge records a gauge metric
^
internal\metrics\business.go:421:1: Comment should end in a period (godot)
// RecordHistogram records a histogram metric
^
internal\metrics\business.go:426:1: Comment should end in a period (godot)
// RecordSummary records a summary metric
^
internal\metrics\business.go:431:1: Comment should end in a period (godot)
// recordMetric is the internal method to record any metric
^
internal\metrics\business.go:479:1: Comment should end in a period (godot)
// GetMetricValues returns raw metric values
^
internal\metrics\business.go:504:1: Comment should end in a period (godot)
// GetAggregatedMetrics returns aggregated metrics
^
internal\metrics\business.go:529:1: Comment should end in a period (godot)
// GetAlertStates returns current alert states
^
internal\metrics\business.go:542:1: Comment should end in a period (godot)
// GetMetrics returns all configured metrics
^
internal\metrics\business.go:555:1: Comment should end in a period (godot)
// Close gracefully shuts down the collector
^
internal\metrics\business.go:667:3: missing cases in switch of type metrics.AggregationType: metrics.AggregationP95, metrics.AggregationP99 (exhaustive)
                switch aggType {
                ^
internal\metrics\business.go:758:40: string `resolved` has 3 occurrences, make it a constant (goconst)
                if !exists || existingState.State == "resolved" {
                                                     ^
internal\metrics\business.go:890:1: Comment should end in a period (godot)
// NewMetricStorage creates a new metric storage backend
^
internal\metrics\business.go:891:39: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func NewMetricStorage(backend string, config map[string]interface{}) (MetricStorage, error) {
                                      ^
internal\metrics\storage.go:10:1: Comment should end in a period (godot)
// MemoryMetricStorage provides in-memory metric storage
^
internal\metrics\storage.go:11:26: fieldalignment: struct with 32 pointer bytes could be 8 (govet)
type MemoryMetricStorage struct {
                         ^
internal\metrics\storage.go:16:1: Comment should end in a period (godot)
// NewMemoryMetricStorage creates a new in-memory metric storage
^
internal\metrics\storage.go:23:1: Comment should end in a period (godot)
// Store stores metric values
^
internal\metrics\storage.go:24:39: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) Store(ctx context.Context, values []MetricValue) error {
                                      ^
internal\metrics\storage.go:39:1: Comment should end in a period (godot)
// Query queries metric values
^
internal\metrics\storage.go:40:39: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) Query(ctx context.Context, query MetricQuery) ([]MetricValue, error) {
                                      ^
internal\metrics\storage.go:69:1: Comment should end in a period (godot)
// Aggregate performs aggregations on metric values
^
internal\metrics\storage.go:112:1: Comment should end in a period (godot)
// Delete removes old metric values
^
internal\metrics\storage.go:113:40: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) Delete(ctx context.Context, before time.Time) error {
                                       ^
internal\metrics\storage.go:130:1: Comment should end in a period (godot)
// Close closes the storage (no-op for memory storage)
^
internal\metrics\storage.go:186:47: unused-parameter: parameter 'groupKey' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) extractLabels(groupKey string, groupBy []string) map[string]string {
                                              ^
internal\metrics\storage.go:186:64: unused-parameter: parameter 'groupBy' seems to be unused, consider removing or renaming it as _ (revive)
func (mms *MemoryMetricStorage) extractLabels(groupKey string, groupBy []string) map[string]string {
                                                               ^
internal\metrics\storage.go:216:3: builtinShadow: shadowing of predeclared identifier: max (gocritic)
                max := values[0].Value
                ^
internal\metrics\storage.go:216:3: redefines-builtin-id: redefinition of the built-in function max (revive)
                max := values[0].Value
                ^
internal\metrics\storage.go:219:5: redefines-builtin-id: redefinition of the built-in function max (revive)
                                max = value.Value
                                ^
internal\metrics\storage.go:225:3: redefines-builtin-id: redefinition of the built-in function min (revive)
                min := values[0].Value
                ^
internal\metrics\storage.go:225:3: builtinShadow: shadowing of predeclared identifier: min (gocritic)
                min := values[0].Value
                ^
internal\metrics\storage.go:228:5: redefines-builtin-id: redefinition of the built-in function min (revive)
                                min = value.Value
                                ^
internal\nats\publisher_error_handler.go:1:1: package-comments: should have a package comment (revive)
package natsx
^
internal\nats\publisher_error_handler.go:12:1: Comment should end in a period (godot)
// Publisher publishes messages to NATS with retry and error handling
^
internal\nats\publisher_error_handler.go:20:1: Comment should end in a period (godot)
// NewPublisher creates a new NATS publisher with error handling
^
internal\nats\publisher_error_handler.go:30:1: Comment should end in a period (godot)
// PublishWithRetry publishes a message with retry logic and error reporting
^
internal\nats\publisher_error_handler.go:58:6: Error return value of `p.js.Publish` is not checked (errcheck)
                _, _ = p.js.Publish(p.subjectErr, ev)
                   ^
internal\nats\publisher_error_handler.go:64:1: Comment should end in a period (godot)
// sanitizeErr prevents leaking credentials in logs
^
internal\observability\enhanced_telemetry.go:1:1: package-comments: should have a package comment (revive)
package observability
^
internal\observability\enhanced_telemetry.go:29:1: Comment should end in a period (godot)
// EnhancedTelemetryService provides comprehensive observability
^
internal\observability\enhanced_telemetry.go:30:31: fieldalignment: struct with 672 pointer bytes could be 568 (govet)
type EnhancedTelemetryService struct {
                              ^
internal\observability\enhanced_telemetry.go:67:2: field `spanMutex` is unused (unused)
        spanMutex   sync.RWMutex
        ^
internal\observability\enhanced_telemetry.go:75:1: Comment should end in a period (godot)
// MetricCollector interface for custom metric collection
^
internal\observability\enhanced_telemetry.go:81:1: Comment should end in a period (godot)
// HealthChecker interface for service health checks
^
internal\observability\enhanced_telemetry.go:87:1: Comment should end in a period (godot)
// HealthResult represents the result of a health check
^
internal\observability\enhanced_telemetry.go:88:19: fieldalignment: struct with 72 pointer bytes could be 56 (govet)
type HealthResult struct {
                  ^
internal\observability\enhanced_telemetry.go:96:1: Comment should end in a period (godot)
// AlertRule defines conditions for triggering alerts
^
internal\observability\enhanced_telemetry.go:97:16: fieldalignment: struct with 120 pointer bytes could be 96 (govet)
type AlertRule struct {
               ^
internal\observability\enhanced_telemetry.go:108:1: Comment should end in a period (godot)
// AlertNotification represents an alert notification
^
internal\observability\enhanced_telemetry.go:109:24: fieldalignment: struct with 160 pointer bytes could be 152 (govet)
type AlertNotification struct {
                       ^
internal\observability\enhanced_telemetry.go:116:1: Comment should end in a period (godot)
// NewEnhancedTelemetryService creates a new enhanced telemetry service
^
internal\observability\enhanced_telemetry.go:117:34: hugeParam: config is heavy (304 bytes); consider passing it by pointer (gocritic)
func NewEnhancedTelemetryService(config TelemetryConfig, logger *zap.Logger) (*EnhancedTelemetryService, error) {
                                 ^
internal\observability\enhanced_telemetry.go:171:1: Comment should end in a period (godot)
// initTracing initializes OpenTelemetry tracing
^
internal\observability\enhanced_telemetry.go:206:1: Comment should end in a period (godot)
// initMetrics initializes OpenTelemetry metrics
^
internal\observability\enhanced_telemetry.go:234:1: Comment should end in a period (godot)
// createMetrics creates all OpenTelemetry metrics
^
internal\observability\enhanced_telemetry.go:327:1: Comment should end in a period (godot)
// initPrometheusMetrics initializes Prometheus metrics
^
internal\observability\enhanced_telemetry.go:375:1: Comment should end in a period (godot)
// collectRuntimeMetrics collects runtime metrics
^
internal\observability\enhanced_telemetry.go:376:60: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (ets *EnhancedTelemetryService) collectRuntimeMetrics(ctx context.Context, observer metric.Observer) error {
                                                           ^
internal\observability\enhanced_telemetry.go:381:46: G115: integer overflow conversion uint64 -> int64 (gosec)
        observer.ObserveInt64(ets.memoryUsage, int64(m.Alloc))
                                                    ^
internal\observability\enhanced_telemetry.go:392:1: Comment should end in a period (godot)
// StartSpan starts a new tracing span
^
internal\observability\enhanced_telemetry.go:393:1: The line is 151 characters long, which exceeds the maximum of 140 characters. (lll)
func (ets *EnhancedTelemetryService) StartSpan(ctx context.Context, name string, opts ...oteltrace.SpanStartOption) (context.Context, oteltrace.Span) {
^
internal\observability\enhanced_telemetry.go:397:1: Comment should end in a period (godot)
// RecordRequest records HTTP request metrics
^
internal\observability\enhanced_telemetry.go:414:1: Comment should end in a period (godot)
// RecordError records error metrics
^
internal\observability\enhanced_telemetry.go:424:1: Comment should end in a period (godot)
// RecordTask records task processing metrics
^
internal\observability\enhanced_telemetry.go:434:1: Comment should end in a period (godot)
// UpdateConnectionCount updates active connection count
^
internal\observability\enhanced_telemetry.go:439:1: Comment should end in a period (godot)
// UpdateDatabaseConnections updates database connection metrics
^
internal\observability\enhanced_telemetry.go:444:1: Comment should end in a period (godot)
// UpdateCacheHitRatio updates cache hit ratio metrics
^
internal\observability\enhanced_telemetry.go:449:1: Comment should end in a period (godot)
// RegisterMetricCollector registers a custom metric collector
^
internal\observability\enhanced_telemetry.go:456:1: Comment should end in a period (godot)
// RegisterHealthChecker registers a health checker
^
internal\observability\enhanced_telemetry.go:463:1: Comment should end in a period (godot)
// RegisterAlertRule registers an alert rule
^
internal\observability\enhanced_telemetry.go:470:1: Comment should end in a period (godot)
// metricsCollectionWorker collects custom metrics periodically
^
internal\observability\enhanced_telemetry.go:502:1: Comment should end in a period (godot)
// healthCheckWorker performs health checks periodically
^
internal\observability\enhanced_telemetry.go:530:1: Comment should end in a period (godot)
// alertingWorker processes alert notifications
^
internal\observability\enhanced_telemetry.go:545:1: Comment should end in a period (godot)
// GetHealthStatus returns the current health status
^
internal\observability\enhanced_telemetry.go:557:1: Comment should end in a period (godot)
// CreateSpanWithError creates a span and records an error if present
^
internal\observability\enhanced_telemetry.go:575:1: Comment should end in a period (godot)
// HTTPMiddleware provides HTTP observability middleware
^
internal\observability\enhanced_telemetry.go:618:1: Comment should end in a period (godot)
// responseWriter wraps http.ResponseWriter to capture metrics
^
internal\observability\enhanced_telemetry.go:636:1: Comment should end in a period (godot)
// Shutdown gracefully shuts down the telemetry service
^
internal\observability\enhanced_telemetry.go:637:47: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (ets *EnhancedTelemetryService) Shutdown(ctx context.Context) error {
                                              ^
internal\observability\integration.go:14:1: Comment should end in a period (godot)
// Service aggregates all observability functionality
^
internal\observability\integration.go:21:1: Comment should end in a period (godot)
// NewService creates a new observability service
^
internal\observability\integration.go:22:17: hugeParam: cfg is heavy (272 bytes); consider passing it by pointer (gocritic)
func NewService(cfg config.TelemetryConfig, logger *zap.Logger) (*Service, error) {
                ^
internal\observability\integration.go:79:1: Comment should end in a period (godot)
// Start initializes the observability service
^
internal\observability\integration.go:101:1: Comment should end in a period (godot)
// Stop gracefully shuts down the observability service
^
internal\observability\integration.go:116:1: Comment should end in a period (godot)
// HTTPMiddleware returns the HTTP telemetry middleware
^
internal\observability\integration.go:126:1: Comment should end in a period (godot)
// GetTelemetryService returns the underlying telemetry service
^
internal\observability\integration.go:131:1: Comment should end in a period (godot)
// HealthCheck returns the observability service health status
^
internal\observability\integration.go:152:1: Comment should end in a period (godot)
// RecordTaskOperation records a task-related operation for telemetry
^
internal\observability\integration.go:166:1: Comment should end in a period (godot)
// RecordDatabaseOperation records a database operation for telemetry
^
internal\observability\integration.go:176:1: Comment should end in a period (godot)
// RecordCacheOperation records a cache operation for telemetry
^
internal\observability\integration.go:185:1: Comment should end in a period (godot)
// RecordMessageQueueOperation records a message queue operation for telemetry
^
internal\observability\integration.go:194:1: Comment should end in a period (godot)
// LogWithTrace logs a message with tracing context
^
internal\observability\middleware.go:17:1: Comment should end in a period (godot)
// HTTPTelemetryMiddleware provides HTTP request instrumentation
^
internal\observability\middleware.go:110:39: unused-parameter: parameter 'operation' seems to be unused, consider removing or renaming it as _ (revive)
                otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
                                                    ^
internal\observability\middleware.go:116:1: Comment should end in a period (godot)
// middlewareResponseWriter wraps http.ResponseWriter to capture response data
^
internal\observability\middleware.go:134:1: Comment should end in a period (godot)
// DatabaseTelemetryWrapper provides database operation instrumentation
^
internal\observability\middleware.go:139:1: Comment should end in a period (godot)
// NewDatabaseTelemetryWrapper creates a new database telemetry wrapper
^
internal\observability\middleware.go:146:1: Comment should end in a period (godot)
// WrapDatabaseOperation wraps a database operation with telemetry
^
internal\observability\middleware.go:188:1: Comment should end in a period (godot)
// CacheOperation wrapper for cache operations
^
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
internal\observability\middleware.go:227:1: Comment should end in a period (godot)
// MessageQueueOperation wrapper for message queue operations
^
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
internal\observability\middleware.go:266:1: Comment should end in a period (godot)
// BusinessOperation wrapper for general business operations
^
internal\observability\middleware.go:300:1: Comment should end in a period (godot)
// AddSpanEvent adds an event to the current span
^
internal\observability\middleware.go:310:1: Comment should end in a period (godot)
// LogEvent logs a structured event with tracing context
^
internal\observability\middleware.go:311:1: paramTypeCombine: func(ctx context.Context, level string, message string, fields ...zap.Field) could be replaced with func(ctx context.Context, level, message string, fields ...zap.Field) (gocritic)
func (ts *TelemetryService) LogEvent(ctx context.Context, level string, message string, fields ...zap.Field) {
^
internal\observability\telemetry.go:26:1: Comment should end in a period (godot)
// TelemetryConfig holds telemetry configuration
^
internal\observability\telemetry.go:27:22: fieldalignment: struct of size 304 could be 264 (govet)
type TelemetryConfig struct {
                     ^
internal\observability\telemetry.go:65:1: Comment should end in a period (godot)
// TelemetryService manages OpenTelemetry instrumentation
^
internal\observability\telemetry.go:66:23: fieldalignment: struct with 496 pointer bytes could be 488 (govet)
type TelemetryService struct {
                      ^
internal\observability\telemetry.go:87:1: Comment should end in a period (godot)
// TaskMetrics holds task-specific metrics
^
internal\observability\telemetry.go:98:1: Comment should end in a period (godot)
// NewTelemetryService creates a new telemetry service
^
internal\observability\telemetry.go:99:26: hugeParam: config is heavy (304 bytes); consider passing it by pointer (gocritic)
func NewTelemetryService(config TelemetryConfig, logger *zap.Logger) (*TelemetryService, error) {
                         ^
internal\observability\telemetry.go:147:1: Comment should end in a period (godot)
// initResource creates the OpenTelemetry resource
^
internal\observability\telemetry.go:148:65: (*TelemetryService).initResource - result 1 (error) is always nil (unparam)
func (ts *TelemetryService) initResource() (*resource.Resource, error) {
                                                                ^
internal\observability\telemetry.go:161:1: Comment should end in a period (godot)
// initTracing sets up distributed tracing
^
internal\observability\telemetry.go:167:2: ifElseChain: rewrite if-else to switch statement (gocritic)
        if ts.config.JaegerEndpoint != "" {
        ^
internal\observability\telemetry.go:223:1: Comment should end in a period (godot)
// initMetrics sets up metrics collection
^
internal\observability\telemetry.go:249:1: Comment should end in a period (godot)
// initBusinessMetrics creates business-specific metrics
^
internal\observability\telemetry.go:300:1: Comment should end in a period (godot)
// initTaskMetrics creates task-specific metrics
^
internal\observability\telemetry.go:364:1: Comment should end in a period (godot)
// initSystemMetrics creates system-level metrics
^
internal\observability\telemetry.go:409:1: Comment should end in a period (godot)
// Start is a no-op method for compatibility with Service.Start()
^
internal\observability\telemetry.go:411:35: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (ts *TelemetryService) Start(ctx context.Context) error {
                                  ^
internal\observability\telemetry.go:418:1: Comment should end in a period (godot)
// Stop gracefully shuts down telemetry
^
internal\observability\telemetry.go:423:1: Comment should end in a period (godot)
// Tracer returns the configured tracer
^
internal\observability\telemetry.go:431:1: Comment should end in a period (godot)
// GetTracer returns a named tracer from the tracer provider
^
internal\observability\telemetry.go:443:1: Comment should end in a period (godot)
// Meter returns the configured meter
^
internal\observability\telemetry.go:451:1: Comment should end in a period (godot)
// GetMeter returns a named meter from the meter provider
^
internal\observability\telemetry.go:463:1: Comment should end in a period (godot)
// StartSpan starts a new trace span
^
internal\observability\telemetry.go:471:1: Comment should end in a period (godot)
// RecordHTTPRequest records HTTP request metrics
^
internal\observability\telemetry.go:487:1: Comment should end in a period (godot)
// RecordError records application errors
^
internal\observability\telemetry.go:501:1: Comment should end in a period (godot)
// RecordTaskCreated records task creation metrics
^
internal\observability\telemetry.go:514:1: Comment should end in a period (godot)
// RecordTaskCompleted records task completion metrics
^
internal\observability\telemetry.go:528:1: Comment should end in a period (godot)
// RecordTaskFailed records task failure metrics
^
internal\observability\telemetry.go:529:1: paramTypeCombine: func(priority string, reason string) could be replaced with func(priority, reason string) (gocritic)
func (ts *TelemetryService) RecordTaskFailed(priority string, reason string) {
^
internal\observability\telemetry.go:542:1: Comment should end in a period (godot)
// IncrementActiveConnections increments active connections counter
^
internal\observability\telemetry.go:550:1: Comment should end in a period (godot)
// DecrementActiveConnections decrements active connections counter
^
internal\observability\telemetry.go:558:1: Comment should end in a period (godot)
// IncrementRequestCounter increments the HTTP request counter
^
internal\observability\telemetry.go:574:1: Comment should end in a period (godot)
// RecordRequestDuration records HTTP request duration
^
internal\observability\telemetry.go:589:1: Comment should end in a period (godot)
// IncrementErrorCounter increments the error counter
^
internal\observability\telemetry.go:604:1: Comment should end in a period (godot)
// RecordProcessingTime records processing time for an operation
^
internal\observability\telemetry.go:628:1: Comment should end in a period (godot)
// HTTPMiddleware returns an HTTP middleware that instruments requests with tracing and metrics
^
internal\observability\telemetry.go:661:30: Error return value of `ts.IncrementRequestCounter` is not checked (errcheck)
                        ts.IncrementRequestCounter(ctx, r.Method, r.URL.Path, statusCode)
                                                  ^
internal\observability\telemetry.go:662:28: Error return value of `ts.RecordRequestDuration` is not checked (errcheck)
                        ts.RecordRequestDuration(ctx, r.Method, r.URL.Path, duration)
                                                ^
internal\observability\telemetry.go:673:1: Comment should end in a period (godot)
// HealthCheck returns health status of the telemetry service
^
internal\observability\telemetry.go:690:16: Error return value is not checked (errcheck)
        components := health["components"].(map[string]interface{})
                      ^
internal\observability\telemetry.go:718:1: Comment should end in a period (godot)
// collectSystemMetrics collects system-level metrics
^
internal\observability\telemetry.go:719:50: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (ts *TelemetryService) collectSystemMetrics(ctx context.Context, observer metric.Observer) error {
                                                 ^
internal\observability\telemetry.go:735:1: Comment should end in a period (godot)
// Shutdown gracefully shuts down the telemetry service
^
internal\observability\telemetry.go:765:1: Comment should end in a period (godot)
// generateInstanceID generates a unique instance identifier
^
internal\observability\telemetry_shim.go:14:1: exported: exported method TelemetryService.RecordCounter should have comment or be unexported (revive)
func (ts *TelemetryService) RecordCounter(name string, value float64, labels map[string]string) {
^
internal\observability\telemetry_shim.go:18:1: exported: exported method TelemetryService.RecordGauge should have comment or be unexported (revive)
func (ts *TelemetryService) RecordGauge(name string, value float64, labels map[string]string) {
^
internal\observability\telemetry_shim.go:22:1: exported: exported method TelemetryService.RecordHistogram should have comment or be unexported (revive)
func (ts *TelemetryService) RecordHistogram(name string, value float64, labels map[string]string) {
^
internal\observability\telemetry_shim.go:28:1: Comment should end in a period (godot)
// RecordCounterWithContext increments a counter metric with context propagation
^
internal\observability\telemetry_shim.go:48:1: Comment should end in a period (godot)
// RecordGaugeWithContext sets a gauge metric with context propagation
^
internal\observability\telemetry_shim.go:69:1: Comment should end in a period (godot)
// RecordHistogramWithContext records a histogram observation with context propagation
^
internal\observability\telemetry_shim.go:89:1: Comment should end in a period (godot)
// labelsToAttributes converts map[string]string to []attribute.KeyValue
^
internal\observability\telemetry_test.go:206:62: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
        testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                                                                    ^
internal\observability\telemetry_test.go:217:9: httpNoBody: http.NoBody should be preferred to the nil request body (gocritic)
        req := httptest.NewRequest("GET", "/test/endpoint", nil)
               ^
internal\observability\telemetry_test.go:298:11: unused-parameter: parameter 'i' seems to be unused, consider removing or renaming it as _ (revive)
                go func(i int) {
                        ^
internal\observability\telemetry_test.go:328:2: ineffectual assignment to ctx (ineffassign)
        ctx, span := tracer.Start(ctx, "test-operation",
        ^
internal\observability\telemetry_test.go:328:2: SA4006: this value of `ctx` is never used (staticcheck)
        ctx, span := tracer.Start(ctx, "test-operation",
        ^
internal\ratelimit\distributed.go:1:1: package-comments: should have a package comment (revive)
package ratelimit
^
internal\ratelimit\distributed.go:16:1: Comment should end in a period (godot)
// Algorithm represents different rate limiting algorithms
^
internal\ratelimit\distributed.go:20:2: exported: exported const AlgorithmTokenBucket should have comment (or a comment on this block) or be unexported (revive)
        AlgorithmTokenBucket   Algorithm = "token_bucket"
        ^
internal\ratelimit\distributed.go:28:1: Comment should end in a period (godot)
// DistributedRateLimiter provides distributed rate limiting capabilities
^
internal\ratelimit\distributed.go:29:29: fieldalignment: struct with 208 pointer bytes could be 104 (govet)
type DistributedRateLimiter struct {
                            ^
internal\ratelimit\distributed.go:36:2: field `mu` is unused (unused)
        mu       sync.RWMutex
        ^
internal\ratelimit\distributed.go:46:1: Comment should end in a period (godot)
// Config configures the distributed rate limiter
^
internal\ratelimit\distributed.go:47:13: fieldalignment: struct of size 112 could be 96 (govet)
type Config struct {
            ^
internal\ratelimit\distributed.go:77:1: Comment should end in a period (godot)
// Rule defines a rate limiting rule
^
internal\ratelimit\distributed.go:78:11: fieldalignment: struct with 248 pointer bytes could be 184 (govet)
type Rule struct {
          ^
internal\ratelimit\distributed.go:109:1: Comment should end in a period (godot)
// Condition represents a condition for rule application
^
internal\ratelimit\distributed.go:117:1: Comment should end in a period (godot)
// Request represents a rate limiting request
^
internal\ratelimit\distributed.go:118:14: fieldalignment: struct with 120 pointer bytes could be 112 (govet)
type Request struct {
             ^
internal\ratelimit\distributed.go:129:1: Comment should end in a period (godot)
// Response represents a rate limiting response
^
internal\ratelimit\distributed.go:130:15: fieldalignment: struct of size 144 could be 136 (govet)
type Response struct {
              ^
internal\ratelimit\distributed.go:149:1: Comment should end in a period (godot)
// Limiter interface for different rate limiting algorithms
^
internal\ratelimit\distributed.go:156:1: Comment should end in a period (godot)
// TokenBucketLimiter implements token bucket algorithm
^
internal\ratelimit\distributed.go:162:1: Comment should end in a period (godot)
// SlidingWindowLimiter implements sliding window algorithm
^
internal\ratelimit\distributed.go:168:1: Comment should end in a period (godot)
// AdaptiveLimiter implements adaptive rate limiting
^
internal\ratelimit\distributed.go:169:22: fieldalignment: struct with 168 pointer bytes could be 64 (govet)
type AdaptiveLimiter struct {
                     ^
internal\ratelimit\distributed.go:178:1: Comment should end in a period (godot)
// AdaptiveState tracks adaptive rate limiting state
^
internal\ratelimit\distributed.go:179:20: fieldalignment: struct with 72 pointer bytes could be 24 (govet)
type AdaptiveState struct {
                   ^
internal\ratelimit\distributed.go:190:1: Comment should end in a period (godot)
// LuaScripts contains Lua scripts for atomic operations
^
internal\ratelimit\distributed.go:199:1: Comment should end in a period (godot)
// DefaultConfig returns default rate limiter configuration
^
internal\ratelimit\distributed.go:221:1: Comment should end in a period (godot)
// NewDistributedRateLimiter creates a new distributed rate limiter
^
internal\ratelimit\distributed.go:222:1: The line is 167 characters long, which exceeds the maximum of 140 characters. (lll)
func NewDistributedRateLimiter(client redis.Cmdable, config Config, logger logger.Logger, telemetry *observability.TelemetryService) (*DistributedRateLimiter, error) {
^
internal\ratelimit\distributed.go:275:1: Comment should end in a period (godot)
// Allow checks if a request should be allowed
^
internal\ratelimit\distributed.go:312:1: Comment should end in a period (godot)
// AllowWithRule checks if a request should be allowed using a specific rule
^
internal\ratelimit\distributed.go:379:1: Comment should end in a period (godot)
// Reset resets the rate limit for a key
^
internal\ratelimit\distributed.go:390:1: Comment should end in a period (godot)
// GetUsage returns current usage for a key
^
internal\ratelimit\distributed.go:400:1: Comment should end in a period (godot)
// GetStats returns rate limiting statistics
^
internal\ratelimit\distributed.go:413:1: Comment should end in a period (godot)
// Close gracefully shuts down the rate limiter
^
internal\ratelimit\distributed.go:423:1: Comment should end in a period (godot)
// Stats contains rate limiting statistics
^
internal\ratelimit\distributed.go:424:12: fieldalignment: struct with 64 pointer bytes could be 24 (govet)
type Stats struct {
           ^
internal\ratelimit\distributed.go:499:10: emptyStringTest: replace `len(requestValue) > 0` with `requestValue != ""` (gocritic)
                return len(requestValue) > 0 && len(fmt.Sprintf("%v", condition.Value)) > 0
                       ^
internal\ratelimit\distributed.go:499:35: emptyStringTest: replace `len(fmt.Sprintf("%v", condition.Value)) > 0` with `fmt.Sprintf("%v", condition.Value) != ""` (gocritic)
                return len(requestValue) > 0 && len(fmt.Sprintf("%v", condition.Value)) > 0
                                                ^
internal\ratelimit\distributed.go:501:10: emptyStringTest: replace `len(requestValue) > 0` with `requestValue != ""` (gocritic)
                return len(requestValue) > 0 && fmt.Sprintf("%v", condition.Value) != ""
                       ^
internal\ratelimit\distributed.go:503:10: emptyStringTest: replace `len(requestValue) > 0` with `requestValue != ""` (gocritic)
                return len(requestValue) > 0 && fmt.Sprintf("%v", condition.Value) != ""
                       ^
internal\ratelimit\distributed.go:526:86: `(*DistributedRateLimiter).recordMetrics` - `key` is unused (unparam)
func (drl *DistributedRateLimiter) recordMetrics(status string, algorithm Algorithm, key string, remaining int64) {
                                                                                     ^
internal\ratelimit\distributed.go:526:86: unused-parameter: parameter 'key' seems to be unused, consider removing or renaming it as _ (revive)
func (drl *DistributedRateLimiter) recordMetrics(status string, algorithm Algorithm, key string, remaining int64) {
                                                                                     ^
internal\ratelimit\distributed.go:598:1: exported: exported method TokenBucketLimiter.Allow should have comment or be unexported (revive)
func (tbl *TokenBucketLimiter) Allow(ctx context.Context, key string, limit int64, window time.Duration) (*Response, error) {
^
internal\ratelimit\distributed.go:605:12: Error return value is not checked (errcheck)
        values := result.([]interface{})
                  ^
internal\ratelimit\distributed.go:606:13: Error return value is not checked (errcheck)
        allowed := values[0].(int64) == 1
                   ^
internal\ratelimit\distributed.go:607:15: Error return value is not checked (errcheck)
        remaining := values[1].(int64)
                     ^
internal\ratelimit\distributed.go:608:25: Error return value is not checked (errcheck)
        resetTime := time.Unix(values[2].(int64), 0)
                               ^
internal\ratelimit\distributed.go:625:1: exported: exported method TokenBucketLimiter.Reset should have comment or be unexported (revive)
func (tbl *TokenBucketLimiter) Reset(ctx context.Context, key string) error {
^
internal\ratelimit\distributed.go:629:1: exported: exported method TokenBucketLimiter.GetUsage should have comment or be unexported (revive)
func (tbl *TokenBucketLimiter) GetUsage(ctx context.Context, key string) (int64, error) {
^
internal\ratelimit\distributed.go:631:5: comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error (errorlint)
        if err == redis.Nil {
           ^
internal\ratelimit\distributed.go:648:1: exported: exported method SlidingWindowLimiter.Allow should have comment or be unexported (revive)
func (swl *SlidingWindowLimiter) Allow(ctx context.Context, key string, limit int64, window time.Duration) (*Response, error) {
^
internal\ratelimit\distributed.go:655:12: Error return value is not checked (errcheck)
        values := result.([]interface{})
                  ^
internal\ratelimit\distributed.go:656:13: Error return value is not checked (errcheck)
        allowed := values[0].(int64) == 1
                   ^
internal\ratelimit\distributed.go:657:11: Error return value is not checked (errcheck)
        count := values[1].(int64)
                 ^
internal\ratelimit\distributed.go:676:1: exported: exported method SlidingWindowLimiter.Reset should have comment or be unexported (revive)
func (swl *SlidingWindowLimiter) Reset(ctx context.Context, key string) error {
^
internal\ratelimit\distributed.go:680:1: exported: exported method SlidingWindowLimiter.GetUsage should have comment or be unexported (revive)
func (swl *SlidingWindowLimiter) GetUsage(ctx context.Context, key string) (int64, error) {
^
internal\ratelimit\distributed.go:688:1: exported: exported method AdaptiveLimiter.Allow should have comment or be unexported (revive)
func (al *AdaptiveLimiter) Allow(ctx context.Context, key string, limit int64, window time.Duration) (*Response, error) {
^
internal\ratelimit\distributed.go:698:1: exported: exported method AdaptiveLimiter.Reset should have comment or be unexported (revive)
func (al *AdaptiveLimiter) Reset(ctx context.Context, key string) error {
^
internal\ratelimit\distributed.go:706:1: exported: exported method AdaptiveLimiter.GetUsage should have comment or be unexported (revive)
func (al *AdaptiveLimiter) GetUsage(ctx context.Context, key string) (int64, error) {
^
internal\ratelimit\distributed.go:733:52: unused-parameter: parameter 'rule' seems to be unused, consider removing or renaming it as _ (revive)
func (al *AdaptiveLimiter) updateState(key string, rule Rule, allowed bool) {
                                                   ^
internal\ratelimit\distributed.go:733:52: `(*AdaptiveLimiter).updateState` - `rule` is unused (unparam)
func (al *AdaptiveLimiter) updateState(key string, rule Rule, allowed bool) {
                                                   ^
internal\repository\postgres\connection.go:1:1: package-comments: should have a package comment (revive)
package postgres
^
internal\repository\postgres\connection.go:7:1: File is not properly formatted (gci)
        _ "github.com/lib/pq"
^
internal\repository\postgres\connection.go:11:1: Comment should end in a period (godot)
// Connect creates a PostgreSQL database connection
^
internal\repository\postgres\task_repository.go:11:1: File is not properly formatted (gci)
        "github.com/google/uuid"
^
internal\repository\postgres\task_repository.go:15:1: Comment should end in a period (godot)
// TaskRepository implements domain.TaskRepository using PostgreSQL
^
internal\repository\postgres\task_repository.go:20:1: Comment should end in a period (godot)
// NewTaskRepository creates a new PostgreSQL task repository
^
internal\repository\postgres\task_repository.go:25:1: Comment should end in a period (godot)
// Create inserts a new task
^
internal\repository\postgres\task_repository.go:28:1: The line is 143 characters long, which exceeds the maximum of 140 characters. (lll)
                INSERT INTO tasks (id, title, description, status, priority, assignee_id, created_by, created_at, updated_at, due_date, tags, metadata)
^
internal\repository\postgres\task_repository.go:32:12: Error return value of `json.Marshal` is not checked (errcheck)
        tagsJSON, _ := json.Marshal(task.Tags)
                  ^
internal\repository\postgres\task_repository.go:33:16: Error return value of `json.Marshal` is not checked (errcheck)
        metadataJSON, _ := json.Marshal(task.Metadata)
                      ^
internal\repository\postgres\task_repository.go:48:1: Comment should end in a period (godot)
// GetByID retrieves a task by ID
^
internal\repository\postgres\task_repository.go:60:1: Comment should end in a period (godot)
// Update updates an existing task
^
internal\repository\postgres\task_repository.go:70:12: Error return value of `json.Marshal` is not checked (errcheck)
        tagsJSON, _ := json.Marshal(task.Tags)
                  ^
internal\repository\postgres\task_repository.go:71:16: Error return value of `json.Marshal` is not checked (errcheck)
        metadataJSON, _ := json.Marshal(task.Metadata)
                      ^
internal\repository\postgres\task_repository.go:83:12: Error return value of `result.RowsAffected` is not checked (errcheck)
        affected, _ := result.RowsAffected()
                  ^
internal\repository\postgres\task_repository.go:91:1: Comment should end in a period (godot)
// Delete removes a task
^
internal\repository\postgres\task_repository.go:100:12: Error return value of `result.RowsAffected` is not checked (errcheck)
        affected, _ := result.RowsAffected()
                  ^
internal\repository\postgres\task_repository.go:108:1: Comment should end in a period (godot)
// List retrieves tasks with filtering and pagination
^
internal\repository\postgres\task_repository.go:173:11: G202: SQL string concatenation (gosec)
        query := `
                SELECT id, title, description, status, priority, assignee_id, created_by,
                       created_at, updated_at, completed_at, due_date, tags, metadata
                FROM tasks ` + whereClause + `
                ORDER BY created_at DESC
                LIMIT $` + strconv.Itoa(argIndex) + ` OFFSET $` + strconv.Itoa(argIndex+1)
internal\repository\postgres\task_repository.go:190:32: rows.Err must be checked (rowserrcheck)
        rows, err := r.db.QueryContext(ctx, query, args...)
                                      ^
internal\repository\postgres\task_repository.go:208:1: Comment should end in a period (godot)
// GetByStatus retrieves tasks by status
^
internal\repository\postgres\task_repository.go:217:32: rows.Err must be checked (rowserrcheck)
        rows, err := r.db.QueryContext(ctx, query, status)
                                      ^
internal\repository\postgres\task_repository.go:235:1: Comment should end in a period (godot)
// GetByAssignee retrieves tasks assigned to a specific user
^
internal\repository\postgres\task_repository.go:244:32: rows.Err must be checked (rowserrcheck)
        rows, err := r.db.QueryContext(ctx, query, assigneeID)
                                      ^
internal\repository\postgres\task_repository.go:262:1: Comment should end in a period (godot)
// scanTask scans a database row into a Task struct
^
internal\repository\postgres\task_repository.go:276:6: comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error (errorlint)
                if err == sql.ErrNoRows {
                   ^
internal\repository\postgres\task_repository.go:284:17: Error return value of `json.Unmarshal` is not checked (errcheck)
                json.Unmarshal(tagsJSON, &task.Tags)
                              ^
internal\repository\postgres\task_repository.go:290:17: Error return value of `json.Unmarshal` is not checked (errcheck)
                json.Unmarshal(metadataJSON, &task.Metadata)
                              ^
internal\repository\redis\cache_repository.go:1:1: package-comments: should have a package comment (revive)
package redis
^
internal\repository\redis\cache_repository.go:12:1: Comment should end in a period (godot)
// CacheRepository implements domain.CacheRepository using Redis
^
internal\repository\redis\cache_repository.go:17:1: Comment should end in a period (godot)
// NewCacheRepository creates a new Redis cache repository
^
internal\repository\redis\cache_repository.go:22:1: Comment should end in a period (godot)
// Set stores a value in cache with TTL
^
internal\repository\redis\cache_repository.go:42:1: Comment should end in a period (godot)
// Get retrieves a value from cache
^
internal\repository\redis\cache_repository.go:45:5: comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error (errorlint)
        if err == redis.Nil {
           ^
internal\repository\redis\cache_repository.go:55:1: Comment should end in a period (godot)
// Delete removes a key from cache
^
internal\repository\redis\cache_repository.go:65:1: Comment should end in a period (godot)
// Exists checks if a key exists in cache
^
internal\repository\redis\cache_repository.go:75:1: Comment should end in a period (godot)
// Increment increments a counter
^
internal\repository\redis\cache_repository.go:85:1: Comment should end in a period (godot)
// SetNX sets a value only if the key doesn't exist (atomic operation)
^
internal\repository\redis\cache_repository.go:105:1: Comment should end in a period (godot)
// GetJSON retrieves and unmarshals a JSON value from cache
^
internal\repository\redis\cache_repository.go:120:1: Comment should end in a period (godot)
// SetWithExpiry sets a value with a specific expiry time
^
internal\repository\redis\cache_repository.go:135:1: Comment should end in a period (godot)
// GetTTL returns the remaining time-to-live of a key
^
internal\repository\redis\cache_repository.go:145:1: Comment should end in a period (godot)
// FlushAll removes all keys (use with caution)
^
internal\repository\redis\connection.go:7:1: File is not properly formatted (gci)
        "github.com/redis/go-redis/v9"
^
internal\repository\redis\connection.go:11:1: Comment should end in a period (godot)
// NewClient creates a new Redis client
^
internal\repository\redis\connection.go:23:1: Comment should end in a period (godot)
// Ping tests Redis connection
^
internal\slo\alerting.go:1:1: package-comments: should have a package comment (revive)
package slo
^
internal\slo\alerting.go:16:1: Comment should end in a period (godot)
// AlertSeverity represents different alert severity levels
^
internal\slo\alerting.go:20:2: exported: exported const SeverityInfo should have comment (or a comment on this block) or be unexported (revive)
        SeverityInfo     AlertSeverity = "info"
        ^
internal\slo\alerting.go:25:1: Comment should end in a period (godot)
// AlertChannel represents different alerting channels
^
internal\slo\alerting.go:29:2: exported: exported const ChannelSlack should have comment (or a comment on this block) or be unexported (revive)
        ChannelSlack     AlertChannel = "slack"
        ^
internal\slo\alerting.go:37:1: Comment should end in a period (godot)
// AlertingConfig holds configuration for the alerting system
^
internal\slo\alerting.go:38:21: fieldalignment: struct with 112 pointer bytes could be 72 (govet)
type AlertingConfig struct {
                    ^
internal\slo\alerting.go:48:1: Comment should end in a period (godot)
// ChannelConfig holds configuration for specific alert channels
^
internal\slo\alerting.go:59:1: Comment should end in a period (godot)
// TemplateConfig holds message templates for different channels
^
internal\slo\alerting.go:68:1: Comment should end in a period (godot)
// RateLimitConfig configures rate limiting for alerts
^
internal\slo\alerting.go:76:1: Comment should end in a period (godot)
// EscalationPolicy defines how alerts should be escalated
^
internal\slo\alerting.go:84:1: Comment should end in a period (godot)
// EscalationStep defines a single step in an escalation policy
^
internal\slo\alerting.go:85:21: fieldalignment: struct with 40 pointer bytes could be 24 (govet)
type EscalationStep struct {
                    ^
internal\slo\alerting.go:91:1: Comment should end in a period (godot)
// SilenceRule defines when alerts should be silenced
^
internal\slo\alerting.go:101:1: Comment should end in a period (godot)
// AlertManager manages SLO-based alerting
^
internal\slo\alerting.go:102:19: fieldalignment: struct with 208 pointer bytes could be 168 (govet)
type AlertManager struct {
                  ^
internal\slo\alerting.go:118:1: Comment should end in a period (godot)
// NewAlertManager creates a new alert manager
^
internal\slo\alerting.go:132:1: Comment should end in a period (godot)
// Start begins the alert processing
^
internal\slo\alerting.go:150:1: Comment should end in a period (godot)
// Stop stops the alert manager
^
internal\slo\alerting.go:155:1: Comment should end in a period (godot)
// SendAlert queues an alert for processing
^
internal\slo\alerting.go:166:1: Comment should end in a period (godot)
// processAlerts processes incoming alerts
^
internal\slo\alerting.go:187:1: Comment should end in a period (godot)
// processAlert processes a single alert
^
internal\slo\alerting.go:229:1: Comment should end in a period (godot)
// shouldSilence checks if an alert should be silenced
^
internal\slo\alerting.go:230:1: cyclomatic complexity 21 of func `(*AlertManager).shouldSilence` is high (> 20) (gocyclo)
func (am *AlertManager) shouldSilence(alert AlertEvent) bool {
^
internal\slo\alerting.go:230:1: calculated cyclomatic complexity for function shouldSilence is 21, max is 20 (cyclop)
func (am *AlertManager) shouldSilence(alert AlertEvent) bool {
^
internal\slo\alerting.go:230:1: cognitive complexity 47 of func `(*AlertManager).shouldSilence` is high (> 30) (gocognit)
func (am *AlertManager) shouldSilence(alert AlertEvent) bool {
^
internal\slo\alerting.go:302:1: Comment should end in a period (godot)
// isRateLimited checks if an alert is rate limited
^
internal\slo\alerting.go:335:1: Comment should end in a period (godot)
// storeAlertHistory stores alert in history
^
internal\slo\alerting.go:349:1: Comment should end in a period (godot)
// getChannelsForSeverity returns channels for a given severity
^
internal\slo\alerting.go:357:1: Comment should end in a period (godot)
// sendToChannel sends an alert to a specific channel
^
internal\slo\alerting.go:383:1: Comment should end in a period (godot)
// sendToSlack sends alert to Slack
^
internal\slo\alerting.go:426:1: Comment should end in a period (godot)
// sendToDiscord sends alert to Discord
^
internal\slo\alerting.go:464:1: Comment should end in a period (godot)
// sendToWebhook sends alert to a generic webhook
^
internal\slo\alerting.go:479:1: Comment should end in a period (godot)
// sendToEmail sends alert via email (placeholder implementation)
^
internal\slo\alerting.go:480:55: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToEmail(alert AlertEvent, config ChannelConfig) error {
                                                      ^
internal\slo\alerting.go:480:55: `(*AlertManager).sendToEmail` - `config` is unused (unparam)
func (am *AlertManager) sendToEmail(alert AlertEvent, config ChannelConfig) error {
                                                      ^
internal\slo\alerting.go:487:1: Comment should end in a period (godot)
// sendToPagerDuty sends alert to PagerDuty (placeholder implementation)
^
internal\slo\alerting.go:488:59: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToPagerDuty(alert AlertEvent, config ChannelConfig) error {
                                                          ^
internal\slo\alerting.go:488:59: `(*AlertManager).sendToPagerDuty` - `config` is unused (unparam)
func (am *AlertManager) sendToPagerDuty(alert AlertEvent, config ChannelConfig) error {
                                                          ^
internal\slo\alerting.go:495:1: Comment should end in a period (godot)
// sendToMSTeams sends alert to Microsoft Teams (placeholder implementation)
^
internal\slo\alerting.go:496:57: unused-parameter: parameter 'config' seems to be unused, consider removing or renaming it as _ (revive)
func (am *AlertManager) sendToMSTeams(alert AlertEvent, config ChannelConfig) error {
                                                        ^
internal\slo\alerting.go:496:57: `(*AlertManager).sendToMSTeams` - `config` is unused (unparam)
func (am *AlertManager) sendToMSTeams(alert AlertEvent, config ChannelConfig) error {
                                                        ^
internal\slo\alerting.go:503:1: Comment should end in a period (godot)
// sendHTTPPayload sends a JSON payload via HTTP POST
^
internal\slo\alerting.go:511:29: should rewrite http.NewRequestWithContext or add (*Request).WithContext (noctx)
        req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
                                   ^
internal\slo\alerting.go:534:1: Comment should end in a period (godot)
// startEscalation starts escalation process for an alert
^
internal\slo\alerting.go:559:1: Comment should end in a period (godot)
// executeEscalation executes an escalation policy
^
internal\slo\alerting.go:587:1: Comment should end in a period (godot)
// cleanup performs periodic cleanup of old data
^
internal\slo\alerting.go:604:1: Comment should end in a period (godot)
// performCleanup cleans up old rate limiter and history data
^
internal\slo\alerting.go:651:7: string `critical` has 3 occurrences, but such constant `SLOStatusCritical` already exists (goconst)
        case "critical":
             ^
internal\slo\alerting.go:653:7: string `warning` has 3 occurrences, but such constant `SeverityWarning` already exists (goconst)
        case "warning":
             ^
internal\slo\alerting.go:675:1: Comment should end in a period (godot)
// GetAlertHistory returns alert history for an SLO
^
internal\slo\alerting.go:683:1: Comment should end in a period (godot)
// GetAllAlertHistory returns all alert history
^
internal\slo\config.go:7:1: Comment should end in a period (godot)
// DefaultSLOs returns the default SLO configuration for MCP Ultra
^
internal\slo\config.go:8:6: Function 'DefaultSLOs' is too long (363 > 150) (funlen)
func DefaultSLOs() []*SLO {
     ^
internal\slo\config.go:114:1: The line is 145 characters long, which exceeds the maximum of 140 characters. (lll)
                                        Expression: "histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket{job=\"mcp-ultra\"}[5m])) by (le)) > 0.5",
^
internal\slo\config.go:387:1: Comment should end in a period (godot)
// GetSLOsByService returns SLOs filtered by service name
^
internal\slo\config.go:398:1: Comment should end in a period (godot)
// GetSLOsByComponent returns SLOs filtered by component name
^
internal\slo\config.go:409:1: Comment should end in a period (godot)
// GetSLOsByType returns SLOs filtered by type
^
internal\slo\config.go:420:1: Comment should end in a period (godot)
// GetCriticalSLOs returns SLOs marked as critical
^
internal\slo\monitor.go:15:1: Comment should end in a period (godot)
// SLOType represents the type of SLO being monitored
^
internal\slo\monitor.go:16:6: exported: type name will be used as slo.SLOType by other packages, and that stutters; consider calling this Type (revive)
type SLOType string
     ^
internal\slo\monitor.go:19:2: exported: exported const SLOTypeAvailability should have comment (or a comment on this block) or be unexported (revive)
        SLOTypeAvailability SLOType = "availability"
        ^
internal\slo\monitor.go:26:1: Comment should end in a period (godot)
// SLOStatus represents the current status of an SLO
^
internal\slo\monitor.go:27:6: exported: type name will be used as slo.SLOStatus by other packages, and that stutters; consider calling this Status (revive)
type SLOStatus string
     ^
internal\slo\monitor.go:30:2: exported: exported const SLOStatusHealthy should have comment (or a comment on this block) or be unexported (revive)
        SLOStatusHealthy   SLOStatus = "healthy"
        ^
internal\slo\monitor.go:36:1: Comment should end in a period (godot)
// SLO represents a Service Level Objective
^
internal\slo\monitor.go:37:10: fieldalignment: struct with 248 pointer bytes could be 192 (govet)
type SLO struct {
         ^
internal\slo\monitor.go:69:1: Comment should end in a period (godot)
// SLOResult represents the result of an SLO evaluation
^
internal\slo\monitor.go:70:6: exported: type name will be used as slo.SLOResult by other packages, and that stutters; consider calling this Result (revive)
type SLOResult struct {
     ^
internal\slo\monitor.go:70:16: fieldalignment: struct with 160 pointer bytes could be 112 (govet)
type SLOResult struct {
               ^
internal\slo\monitor.go:81:1: Comment should end in a period (godot)
// ErrorBudget represents the error budget information
^
internal\slo\monitor.go:82:18: fieldalignment: struct with 56 pointer bytes could be 24 (govet)
type ErrorBudget struct {
                 ^
internal\slo\monitor.go:90:1: Comment should end in a period (godot)
// BurnRate represents burn rate information
^
internal\slo\monitor.go:100:1: Comment should end in a period (godot)
// CompliancePoint represents a point in time compliance measurement
^
internal\slo\monitor.go:101:22: fieldalignment: struct with 40 pointer bytes could be 32 (govet)
type CompliancePoint struct {
                     ^
internal\slo\monitor.go:107:1: Comment should end in a period (godot)
// AlertRule represents an alerting rule for an SLO
^
internal\slo\monitor.go:108:16: fieldalignment: struct with 72 pointer bytes could be 56 (govet)
type AlertRule struct {
               ^
internal\slo\monitor.go:118:1: Comment should end in a period (godot)
// Monitor manages SLO monitoring and evaluation
^
internal\slo\monitor.go:119:14: fieldalignment: struct with 104 pointer bytes could be 64 (govet)
type Monitor struct {
             ^
internal\slo\monitor.go:136:1: Comment should end in a period (godot)
// AlertEvent represents an SLO alert event
^
internal\slo\monitor.go:137:17: fieldalignment: struct with 104 pointer bytes could be 96 (govet)
type AlertEvent struct {
                ^
internal\slo\monitor.go:147:1: Comment should end in a period (godot)
// StatusEvent represents an SLO status change event
^
internal\slo\monitor.go:156:1: Comment should end in a period (godot)
// NewMonitor creates a new SLO monitor
^
internal\slo\monitor.go:173:1: Comment should end in a period (godot)
// AddSLO adds an SLO to the monitor
^
internal\slo\monitor.go:210:1: Comment should end in a period (godot)
// RemoveSLO removes an SLO from monitoring
^
internal\slo\monitor.go:220:1: Comment should end in a period (godot)
// GetSLO retrieves an SLO by name
^
internal\slo\monitor.go:229:1: Comment should end in a period (godot)
// GetAllSLOs returns all configured SLOs
^
internal\slo\monitor.go:241:1: Comment should end in a period (godot)
// GetSLOResult retrieves the latest SLO evaluation result
^
internal\slo\monitor.go:250:1: Comment should end in a period (godot)
// GetAllSLOResults returns all SLO evaluation results
^
internal\slo\monitor.go:262:1: Comment should end in a period (godot)
// Start begins SLO monitoring
^
internal\slo\monitor.go:283:1: Comment should end in a period (godot)
// Stop stops SLO monitoring
^
internal\slo\monitor.go:288:1: Comment should end in a period (godot)
// AlertChannel returns the alert event channel
^
internal\slo\monitor.go:293:1: Comment should end in a period (godot)
// StatusChannel returns the status change event channel
^
internal\slo\monitor.go:298:1: Comment should end in a period (godot)
// evaluateAllSLOs evaluates all configured SLOs
^
internal\slo\monitor.go:318:1: Comment should end in a period (godot)
// evaluateSLO evaluates a single SLO
^
internal\slo\monitor.go:371:1: Comment should end in a period (godot)
// queryPrometheus executes a Prometheus query
^
internal\slo\monitor.go:403:1: Comment should end in a period (godot)
// calculateErrorBudget calculates the error budget for an SLO
^
internal\slo\monitor.go:445:1: Comment should end in a period (godot)
// calculateBurnRate calculates the burn rate for an SLO
^
internal\slo\monitor.go:489:1: Comment should end in a period (godot)
// determineStatus determines the SLO status based on current metrics
^
internal\slo\monitor.go:509:1: Comment should end in a period (godot)
// getComplianceHistory retrieves historical compliance data
^
internal\slo\monitor.go:558:1: Comment should end in a period (godot)
// storeResult stores an SLO evaluation result and checks for status changes
^
internal\slo\monitor.go:586:1: Comment should end in a period (godot)
// checkAndGenerateAlerts checks if alerts should be generated for an SLO result
^
internal\telemetry\metrics.go:1:1: package-comments: package comment should be of the form "Package telemetry ..." (revive)
// internal/telemetry/metrics.go
^
internal\telemetry\metrics.go:13:2: exported: exported var RequestDuration should have comment or be unexported (revive)
        RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        ^
internal\telemetry\metrics.go:30:1: exported: exported function Middleware should have comment or be unexported (revive)
func Middleware(next http.Handler) http.Handler {
^
internal\telemetry\telemetry.go:23:2: Comment should end in a period (godot)
        // HTTP Metrics
        ^
internal\telemetry\telemetry.go:41:2: Comment should end in a period (godot)
        // Business Metrics
        ^
internal\telemetry\telemetry.go:59:2: Comment should end in a period (godot)
        // System Metrics
        ^
internal\telemetry\telemetry.go:77:1: Comment should end in a period (godot)
// Telemetry holds telemetry configuration and clients
^
internal\telemetry\telemetry.go:83:1: Comment should end in a period (godot)
// Init initializes telemetry system
^
internal\telemetry\telemetry.go:84:11: hugeParam: cfg is heavy (272 bytes); consider passing it by pointer (gocritic)
func Init(cfg config.TelemetryConfig) (*Telemetry, error) {
          ^
internal\telemetry\telemetry.go:84:11: unused-parameter: parameter 'cfg' seems to be unused, consider removing or renaming it as _ (revive)
func Init(cfg config.TelemetryConfig) (*Telemetry, error) {
          ^
internal\telemetry\telemetry.go:109:1: Comment should end in a period (godot)
// HTTPMetrics middleware for HTTP request metrics
^
internal\telemetry\telemetry.go:129:1: Comment should end in a period (godot)
// RecordTaskCreated records task creation metrics
^
internal\telemetry\telemetry.go:134:1: Comment should end in a period (godot)
// RecordTaskProcessingTime records task processing time
^
internal\telemetry\telemetry.go:139:1: Comment should end in a period (godot)
// RecordDatabaseConnections records database connection metrics
^
internal\telemetry\telemetry.go:144:1: Comment should end in a period (godot)
// RecordCacheOperation records cache operation metrics
^
internal\telemetry\telemetry.go:149:1: Comment should end in a period (godot)
// TaskMetrics handles task-related metrics
^
internal\telemetry\telemetry.go:157:1: Comment should end in a period (godot)
// NewTaskMetrics creates new task metrics
^
internal\telemetry\telemetry.go:192:1: Comment should end in a period (godot)
// RecordTaskCreated records a task creation
^
internal\telemetry\telemetry.go:202:1: Comment should end in a period (godot)
// RecordTaskCompleted records a task completion
^
internal\telemetry\telemetry.go:217:1: Comment should end in a period (godot)
// FeatureFlagMetrics handles feature flag metrics
^
internal\telemetry\telemetry.go:223:1: Comment should end in a period (godot)
// NewFeatureFlagMetrics creates new feature flag metrics
^
internal\telemetry\telemetry.go:239:1: Comment should end in a period (godot)
// RecordEvaluation records a feature flag evaluation
^
internal\telemetry\tracing.go:22:6: exported: exported type TracingConfig should have comment or be unexported (revive)
type TracingConfig struct {
     ^
internal\telemetry\tracing.go:22:20: fieldalignment: struct with 168 pointer bytes could be 136 (govet)
type TracingConfig struct {
                   ^
internal\telemetry\tracing.go:44:6: exported: exported type TracingProvider should have comment or be unexported (revive)
type TracingProvider struct {
     ^
internal\telemetry\tracing.go:51:1: exported: exported function NewTracingProvider should have comment or be unexported (revive)
func NewTracingProvider(config *TracingConfig, logger *zap.Logger) (*TracingProvider, error) {
^
internal\telemetry\tracing.go:170:3: appendCombine: can combine chain of 2 appends into one (gocritic)
                opts = append(opts, jaeger.WithUsername(config.JaegerUser))
                ^
internal\telemetry\tracing.go:184:1: Comment should end in a period (godot)
// GetTracer returns a tracer for the given name
^
internal\telemetry\tracing.go:192:1: Comment should end in a period (godot)
// Shutdown gracefully shuts down the tracing provider
^
internal\telemetry\tracing.go:203:1: Comment should end in a period (godot)
// TraceFunction wraps a function with tracing
^
internal\telemetry\tracing.go:217:1: Comment should end in a period (godot)
// TraceFunctionWithResult wraps a function with tracing and returns a result
^
internal\telemetry\tracing.go:233:1: Comment should end in a period (godot)
// AddSpanAttributes adds multiple attributes to the current span
^
internal\telemetry\tracing.go:241:1: Comment should end in a period (godot)
// AddSpanEvent adds an event to the current span
^
internal\telemetry\tracing.go:249:1: Comment should end in a period (godot)
// SetSpanError sets error status on the current span
^
internal\telemetry\tracing.go:258:1: Comment should end in a period (godot)
// GetTraceID returns the trace ID from the current context
^
internal\telemetry\tracing.go:267:1: Comment should end in a period (godot)
// GetSpanID returns the span ID from the current context
^
internal\telemetry\tracing.go:276:1: Comment should end in a period (godot)
// InjectTraceContext injects trace context into a map (for cross-service calls)
^
internal\telemetry\tracing.go:281:1: Comment should end in a period (godot)
// ExtractTraceContext extracts trace context from a map
^
internal\telemetry\tracing.go:286:1: Comment should end in a period (godot)
// mapCarrier implements the TextMapCarrier interface
^
internal\telemetry\tracing.go:307:1: Comment should end in a period (godot)
// noopExporter is a no-op span exporter for disabled tracing
^
internal\telemetry\tracing.go:318:1: Comment should end in a period (godot)
// Span naming conventions
^
internal\telemetry\tracing.go:331:1: Comment should end in a period (godot)
// Common span attributes
^
internal\telemetry\tracing_test.go:134:77: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
                err := TraceFunction(context.Background(), tracer, "test-operation", func(ctx context.Context) error {
                                                                                          ^
internal\telemetry\tracing_test.go:145:80: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
                err := TraceFunction(context.Background(), tracer, "failing-operation", func(ctx context.Context) error {
                                                                                             ^
internal\telemetry\tracing_test.go:169:95: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
                result, err := TraceFunctionWithResult(context.Background(), tracer, "test-operation", func(ctx context.Context) (string, error) {
                                                                                                            ^
internal\telemetry\tracing_test.go:179:1: The line is 141 characters long, which exceeds the maximum of 140 characters. (lll)
                result, err := TraceFunctionWithResult(context.Background(), tracer, "failing-operation", func(ctx context.Context) (string, error) {
^
internal\telemetry\tracing_test.go:179:98: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
                result, err := TraceFunctionWithResult(context.Background(), tracer, "failing-operation", func(ctx context.Context) (string, error) {
                                                                                                               ^
internal\telemetry\tracing_test.go:202:43: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should add span attributes", func(t *testing.T) {
                                                 ^
internal\telemetry\tracing_test.go:218:39: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should add span events", func(t *testing.T) {
                                             ^
internal\telemetry\tracing_test.go:231:38: unused-parameter: parameter 't' seems to be unused, consider removing or renaming it as _ (revive)
        t.Run("should set span error", func(t *testing.T) {
                                            ^
internal\testhelpers\helpers.go:1:1: package-comments: should have a package comment (revive)
package testhelpers
^
internal\testhelpers\helpers.go:8:1: Comment should end in a period (godot)
// GetTestJWTSecret returns a safe test JWT secret
^
internal\testhelpers\helpers.go:13:1: Comment should end in a period (godot)
// GenerateTestSecret generates a random test secret
^
internal\testhelpers\helpers.go:22:1: Comment should end in a period (godot)
// GetTestDatabaseURL returns a test database URL
^
internal\testhelpers\helpers.go:27:1: Comment should end in a period (godot)
// GetTestRedisURL returns a test Redis URL
^
internal\testhelpers\helpers.go:32:1: Comment should end in a period (godot)
// GetTestNATSURL returns a test NATS URL
^
internal\tracing\business.go:1:1: package-comments: should have a package comment (revive)
package tracing
^
internal\tracing\business.go:20:1: Comment should end in a period (godot)
// BusinessTransactionTracer provides advanced tracing for critical business transactions
^
internal\tracing\business.go:21:32: fieldalignment: struct with 336 pointer bytes could be 200 (govet)
type BusinessTransactionTracer struct {
                               ^
internal\tracing\business.go:39:1: Comment should end in a period (godot)
// TracingConfig configures business transaction tracing
^
internal\tracing\business.go:40:6: exported: type name will be used as tracing.TracingConfig by other packages, and that stutters; consider calling this Config (revive)
type TracingConfig struct {
     ^
internal\tracing\business.go:40:20: fieldalignment: struct of size 232 could be 200 (govet)
type TracingConfig struct {
                   ^
internal\tracing\business.go:76:1: Comment should end in a period (godot)
// AlertThresholds defines alerting thresholds
^
internal\tracing\business.go:85:1: Comment should end in a period (godot)
// BusinessTransaction represents a high-level business transaction
^
internal\tracing\business.go:86:26: fieldalignment: struct with 472 pointer bytes could be 360 (govet)
type BusinessTransaction struct {
                         ^
internal\tracing\business.go:129:1: Comment should end in a period (godot)
// TransactionType represents different types of business transactions
^
internal\tracing\business.go:133:2: exported: exported const TransactionTypeAPI should have comment (or a comment on this block) or be unexported (revive)
        TransactionTypeAPI        TransactionType = "api"
        ^
internal\tracing\business.go:145:1: Comment should end in a period (godot)
// TransactionStatus represents transaction status
^
internal\tracing\business.go:149:2: exported: exported const TransactionStatusStarted should have comment (or a comment on this block) or be unexported (revive)
        TransactionStatusStarted    TransactionStatus = "started"
        ^
internal\tracing\business.go:157:1: Comment should end in a period (godot)
// TransactionStep represents a step within a business transaction
^
internal\tracing\business.go:158:22: fieldalignment: struct with 144 pointer bytes could be 136 (govet)
type TransactionStep struct {
                     ^
internal\tracing\business.go:173:1: Comment should end in a period (godot)
// TransactionEvent represents an event within a transaction
^
internal\tracing\business.go:183:1: Comment should end in a period (godot)
// TransactionError represents an error within a transaction
^
internal\tracing\business.go:195:1: Comment should end in a period (godot)
// TransactionMetrics contains transaction performance metrics
^
internal\tracing\business.go:208:1: Comment should end in a period (godot)
// TransactionTemplate defines a template for transaction creation
^
internal\tracing\business.go:209:26: fieldalignment: struct with 128 pointer bytes could be 96 (govet)
type TransactionTemplate struct {
                         ^
internal\tracing\business.go:227:1: Comment should end in a period (godot)
// EventLevel represents the severity level of an event
^
internal\tracing\business.go:231:2: exported: exported const EventLevelDebug should have comment (or a comment on this block) or be unexported (revive)
        EventLevelDebug    EventLevel = "debug"
        ^
internal\tracing\business.go:238:1: Comment should end in a period (godot)
// DefaultTracingConfig returns default tracing configuration
^
internal\tracing\business.go:270:1: Comment should end in a period (godot)
// NewBusinessTransactionTracer creates a new business transaction tracer
^
internal\tracing\business.go:271:1: The line is 158 characters long, which exceeds the maximum of 140 characters. (lll)
func NewBusinessTransactionTracer(config TracingConfig, logger logger.Logger, telemetry *observability.TelemetryService) (*BusinessTransactionTracer, error) {
^
internal\tracing\business.go:306:1: Comment should end in a period (godot)
// StartTransaction starts a new business transaction
^
internal\tracing\business.go:307:1: The line is 198 characters long, which exceeds the maximum of 140 characters. (lll)
func (btt *BusinessTransactionTracer) StartTransaction(ctx context.Context, transactionType TransactionType, name string, attributes map[string]interface{}) (*BusinessTransaction, context.Context) {
^
internal\tracing\business.go:389:1: Comment should end in a period (godot)
// EndTransaction ends a business transaction
^
internal\tracing\business.go:455:1: Comment should end in a period (godot)
// StartStep starts a new step within a transaction
^
internal\tracing\business.go:456:1: The line is 162 characters long, which exceeds the maximum of 140 characters. (lll)
func (btt *BusinessTransactionTracer) StartStep(transaction *BusinessTransaction, stepName, stepType string, attributes map[string]interface{}) *TransactionStep {
^
internal\tracing\business.go:489:1: Comment should end in a period (godot)
// EndStep ends a transaction step
^
internal\tracing\business.go:533:1: Comment should end in a period (godot)
// AddEvent adds an event to a transaction
^
internal\tracing\business.go:534:1: The line is 164 characters long, which exceeds the maximum of 140 characters. (lll)
func (btt *BusinessTransactionTracer) AddEvent(transaction *BusinessTransaction, eventType, eventName string, level EventLevel, attributes map[string]interface{}) {
^
internal\tracing\business.go:568:1: Comment should end in a period (godot)
// AddError adds an error to a transaction
^
internal\tracing\business.go:573:1: Comment should end in a period (godot)
// GetTransaction retrieves a transaction by ID
^
internal\tracing\business.go:588:1: Comment should end in a period (godot)
// ListActiveTransactions returns all currently active transactions
^
internal\tracing\business.go:604:1: Comment should end in a period (godot)
// GetTransactionMetrics returns aggregated metrics for transactions
^
internal\tracing\business.go:625:1: `if transaction.Duration > 0` has complex nested blocks (complexity: 6) (nestif)
                if transaction.Duration > 0 {
^
internal\tracing\business.go:650:1: Comment should end in a period (godot)
// RegisterTemplate registers a transaction template
^
internal\tracing\business.go:664:1: Comment should end in a period (godot)
// Close gracefully shuts down the tracer
^
internal\tracing\business.go:674:1: Comment should end in a period (godot)
// TransactionAnalytics contains transaction analytics
^
internal\tracing\business.go:675:27: fieldalignment: struct with 24 pointer bytes could be 16 (govet)
type TransactionAnalytics struct {
                          ^
internal\tracing\business.go:735:83: `(*BusinessTransactionTracer).shouldSample` - `attributes` is unused (unparam)
func (btt *BusinessTransactionTracer) shouldSample(template *TransactionTemplate, attributes map[string]interface{}) bool {
                                                                                  ^
internal\tracing\business.go:735:83: unused-parameter: parameter 'attributes' seems to be unused, consider removing or renaming it as _ (revive)
func (btt *BusinessTransactionTracer) shouldSample(template *TransactionTemplate, attributes map[string]interface{}) bool {
                                                                                  ^
internal\tracing\business.go:749:1: The line is 210 characters long, which exceeds the maximum of 140 characters. (lll)
func (btt *BusinessTransactionTracer) createLightweightTransaction(ctx context.Context, transactionType TransactionType, name string, attributes map[string]interface{}) (*BusinessTransaction, context.Context) {
^
internal\tracing\business.go:802:7: Error return value of `baggage.Parse` is not checked (errcheck)
        bag, _ := baggage.Parse(fmt.Sprintf("transaction.id=%s,transaction.type=%s,transaction.name=%s",
             ^
internal\tracing\business.go:806:11: Error return value of `baggage.NewMember` is not checked (errcheck)
                member, _ := baggage.NewMember("user.id", transaction.UserID)
                        ^
internal\tracing\business.go:807:8: Error return value of `bag.SetMember` is not checked (errcheck)
                bag, _ = bag.SetMember(member)
                     ^
main.go:27:19: Error return value of `logger.Sync` is not checked (errcheck)
        defer logger.Sync()
                         ^
main.go:74:30: unnecessary conversion (unconvert)
                ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
                                           ^
main.go:75:30: unnecessary conversion (unconvert)
                WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
                                           ^
main.go:76:30: unnecessary conversion (unconvert)
                IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
                                           ^
scripts\generate-secrets.go:5:1: package-comments: should have a package comment (revive)
package main
^
scripts\generate-secrets.go:15:1: Comment should end in a period (godot)
// generateRandomHex creates a cryptographically secure random hex string
^
test\mocks\mocks.go:1:1: package-comments: should have a package comment (revive)
package mocks
^
test\mocks\mocks.go:8:1: File is not properly formatted (gci)
        "github.com/stretchr/testify/mock"
^
test\mocks\mocks.go:12:1: Comment should end in a period (godot)
// MockTaskRepository is a mock implementation of TaskRepository
^
test\mocks\mocks.go:17:1: exported: exported method MockTaskRepository.Create should have comment or be unexported (revive)
func (m *MockTaskRepository) Create(ctx context.Context, task *domain.Task) error {
^
test\mocks\mocks.go:22:1: exported: exported method MockTaskRepository.GetByID should have comment or be unexported (revive)
func (m *MockTaskRepository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
^
test\mocks\mocks.go:34:1: exported: exported method MockTaskRepository.Update should have comment or be unexported (revive)
func (m *MockTaskRepository) Update(ctx context.Context, task *domain.Task) error {
^
test\mocks\mocks.go:39:1: exported: exported method MockTaskRepository.Delete should have comment or be unexported (revive)
func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
^
test\mocks\mocks.go:44:1: exported: exported method MockTaskRepository.List should have comment or be unexported (revive)
func (m *MockTaskRepository) List(ctx context.Context, limit, offset int) ([]*domain.Task, error) {
^
test\mocks\mocks.go:49:9: Error return value is not checked (errcheck)
        return args.Get(0).([]*domain.Task), args.Error(1)
               ^
test\mocks\mocks.go:52:1: Comment should end in a period (godot)
// MockCacheRepository is a mock implementation of CacheRepository
^
test\mocks\mocks.go:57:1: exported: exported method MockCacheRepository.Get should have comment or be unexported (revive)
func (m *MockCacheRepository) Get(ctx context.Context, key string) (interface{}, error) {
^
test\mocks\mocks.go:62:1: exported: exported method MockCacheRepository.Set should have comment or be unexported (revive)
func (m *MockCacheRepository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
^
test\mocks\mocks.go:67:1: exported: exported method MockCacheRepository.Delete should have comment or be unexported (revive)
func (m *MockCacheRepository) Delete(ctx context.Context, key string) error {
^
test\mocks\mocks.go:72:1: exported: exported method MockCacheRepository.Exists should have comment or be unexported (revive)
func (m *MockCacheRepository) Exists(ctx context.Context, key string) (bool, error) {
^
test\mocks\mocks.go:77:1: Comment should end in a period (godot)
// MockEventBus is a mock implementation of EventBus
^
test\mocks\mocks.go:82:1: exported: exported method MockEventBus.Publish should have comment or be unexported (revive)
func (m *MockEventBus) Publish(ctx context.Context, subject string, data []byte) error {
^
test\mocks\mocks.go:87:1: exported: exported method MockEventBus.Subscribe should have comment or be unexported (revive)
func (m *MockEventBus) Subscribe(ctx context.Context, subject string, handler func([]byte)) error {
^
test\mocks\mocks.go:92:1: exported: exported method MockEventBus.Close should have comment or be unexported (revive)
func (m *MockEventBus) Close() error {
^
test\mocks\mocks.go:97:1: Comment should end in a period (godot)
// MockValidator is a mock implementation of Validator
^
test\mocks\mocks.go:102:1: exported: exported method MockValidator.Validate should have comment or be unexported (revive)
func (m *MockValidator) Validate(ctx context.Context, data interface{}) error {
^
PS E:\vertikon\business\SaaS\templates\mcp-ultra> go test ./... -count=1 -cover
ok      github.com/vertikon/mcp-ultra   0.572s  coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/api/grpc/gen/compliance/v1                coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/api/grpc/gen/system/v1            coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/api/grpc/gen/task/v1              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/automation                coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/ai/events        0.558s  coverage: 100.0% of statements
        github.com/vertikon/mcp-ultra/internal/ai/router                coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/ai/telemetry     0.599s  coverage: 87.9% of statements
# github.com/vertikon/mcp-ultra/internal/compliance [github.com/vertikon/mcp-ultra/internal/compliance.test]
internal\compliance\framework_test.go:202:3: unknown field UserID in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:203:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:204:3: unknown field Purpose in struct literal of type DataAccessRequest
internal\compliance\framework_test.go:208:17: assignment mismatch: 2 variables but framework.ProcessDataAccessRequest returns 1 value
internal\compliance\framework_test.go:230:3: unknown field UserID in struct literal of type DataDeletionRequest
internal\compliance\framework_test.go:231:14: cannot use uuid.New() (value of array type uuid.UUID) as string value in struct literal
internal\compliance\framework_test.go:236:17: assignment mismatch: 2 variables but framework.ProcessDataDeletionRequest returns 1 value
internal\compliance\framework_test.go:254:25: assignment mismatch: 2 variables but framework.AnonymizeData returns 1 value
internal\compliance\framework_test.go:254:68: too many arguments in call to framework.AnonymizeData
        have (context.Context, map[string]interface{}, string)
        want (context.Context, string)
internal\compliance\framework_test.go:279:46: too many arguments in call to framework.LogAuditEvent
        have (context.Context, uuid.UUID, string, map[string]interface{})
        want (context.Context, AuditEvent)
internal\compliance\framework_test.go:279:46: too many errors
ok      github.com/vertikon/mcp-ultra/internal/ai/wiring        0.543s  coverage: 80.0% of statements
--- FAIL: TestCircuitBreaker_HalfOpenMaxRequests (0.06s)
    circuit_breaker_test.go:260:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/circuit_breaker_test.go:260
                Error:          Should be false
                Test:           TestCircuitBreaker_HalfOpenMaxRequests
                Messages:       Request should be denied after max half-open requests
--- FAIL: TestDistributedCache_SetAndGet (0.02s)
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
--- FAIL: TestDistributedCache_Delete (0.03s)
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
--- FAIL: TestDistributedCache_SetComplexObject (0.00s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:181
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_SetComplexObject
--- FAIL: TestDistributedCache_ConcurrentOperations (0.01s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:232
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_ConcurrentOperations
--- FAIL: TestDistributedCache_Namespace (0.07s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:268
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_Namespace
--- FAIL: TestCacheStrategy_WriteThrough (0.16s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:297
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestCacheStrategy_WriteThrough
--- FAIL: TestDistributedCache_InvalidKey (0.03s)
    distributed_test.go:63:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:63
                                                        E:/vertikon/business/SaaS/templates/mcp-ultra/internal/cache/distributed_test.go:316
                Error:          Received unexpected error:
                                failed to connect to Redis cluster: ERR unknown command `readonly`, with args beginning with:
                Test:           TestDistributedCache_InvalidKey
FAIL
coverage: 17.7% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/cache    1.237s
FAIL    github.com/vertikon/mcp-ultra/internal/compliance [build failed]
--- FAIL: TestNewTLSManager (0.11s)
    logger.go:146: 2025-10-17T15:35:34.382-0300 INFO    TLS is disabled
    --- FAIL: TestNewTLSManager/should_create_manager_with_valid_TLS_config (0.03s)
        tls_test.go:120:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/config/tls_test.go:120
                Error:          Received unexpected error:
                                failed to load TLS configuration: failed to load certificate pair: tls: failed to find any PEM data in key input
                Test:           TestNewTLSManager/should_create_manager_with_valid_TLS_config
--- FAIL: TestTLSManager_GetTLSConfig (0.03s)
    --- FAIL: TestTLSManager_GetTLSConfig/should_return_copy_of_TLS_config (0.03s)
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
FAIL    github.com/vertikon/mcp-ultra/internal/config   0.868s
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
FAIL    github.com/vertikon/mcp-ultra/internal/domain   0.546s
        github.com/vertikon/mcp-ultra/internal/dr               coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/events           coverage: 0.0% of statements
ok      github.com/vertikon/mcp-ultra/internal/features 0.524s  coverage: 22.0% of statements
# github.com/vertikon/mcp-ultra/internal/middleware [github.com/vertikon/mcp-ultra/internal/middleware.test]
internal\middleware\auth_test.go:96:30: undefined: testhelpers.GetTestAPIKeys
ok      github.com/vertikon/mcp-ultra/internal/handlers 0.690s  coverage: 100.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/handlers/http [build failed]
        github.com/vertikon/mcp-ultra/internal/http             coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/lifecycle                coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/metrics          coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/middleware [build failed]
        github.com/vertikon/mcp-ultra/internal/nats             coverage: 0.0% of statements
# github.com/vertikon/mcp-ultra/internal/services [github.com/vertikon/mcp-ultra/internal/services.test]
internal\services\task_service_test.go:104:70: undefined: domain.UserFilter
internal\services\task_service_test.go:171:28: cannot use taskRepo (variable of type *mockTaskRepository) as domain.TaskRepository value in argument to NewTaskService: *mockTaskRepository does not implement domain.TaskRepository (wrong type for method List)
                have List(context.Context, domain.TaskFilter) ([]*domain.Task, error)
                want List(context.Context, domain.TaskFilter) ([]*domain.Task, int, error)
internal\services\task_service_test.go:171:48: cannot use eventRepo (variable of type *mockEventRepository) as domain.EventRepository value in argument to NewTaskService: *mockEventRepository does not implement domain.EventRepository (missing method GetByType)
internal\services\task_service_test.go:171:59: cannot use cacheRepo (variable of type *mockCacheRepository) as domain.CacheRepository value in argument to NewTaskService: *mockCacheRepository does not implement domain.CacheRepository (missing method Exists)
internal\services\task_service_test.go:199:31: declared and not used: eventRepo
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
    logger.go:146: 2025-10-17T15:35:39.210-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T15:35:39.210-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T15:35:39.210-0300 DEBUG   TelemetryService.Start called (initialization already complete)
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
    logger.go:146: 2025-10-17T15:35:39.211-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryService_SpanAttributes (0.00s)
    logger.go:146: 2025-10-17T15:35:39.229-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T15:35:39.229-0300 INFO    Telemetry initialized successfully      {"service": "test-service", "version": "1.0.0", "environment": "test"}
    logger.go:146: 2025-10-17T15:35:39.229-0300 DEBUG   TelemetryService.Start called (initialization already complete)
    telemetry_test.go:345:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:345
                Error:          Should be true
                Test:           TestTelemetryService_SpanAttributes
    logger.go:146: 2025-10-17T15:35:39.229-0300 INFO    Telemetry service shutdown complete
--- FAIL: TestTelemetryConfig_Validation (0.00s)
    logger.go:146: 2025-10-17T15:35:39.229-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T15:35:39.229-0300 INFO    Telemetry initialized successfully      {"service": "test", "version": "", "environment": ""}
    logger.go:146: 2025-10-17T15:35:39.229-0300 DEBUG   No tracing exporter configured, using no-op tracer
    logger.go:146: 2025-10-17T15:35:39.230-0300 INFO    Telemetry initialized successfully      {"service": "", "version": "", "environment": ""}
    telemetry_test.go:376:
                Error Trace:    E:/vertikon/business/SaaS/templates/mcp-ultra/internal/observability/telemetry_test.go:376
                Error:          Should NOT be empty, but was
                Test:           TestTelemetryConfig_Validation
FAIL
coverage: 22.1% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/observability    0.971s
        github.com/vertikon/mcp-ultra/internal/ratelimit                coverage: 0.0% of statements
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
        github.com/vertikon/mcp-ultra/internal/repository/postgres              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/repository/redis         coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/internal/security [build failed]
FAIL    github.com/vertikon/mcp-ultra/internal/services [build failed]
        github.com/vertikon/mcp-ultra/internal/slo              coverage: 0.0% of statements
# github.com/vertikon/mcp-ultra/test/property [github.com/vertikon/mcp-ultra/test/property.test]
test\property\task_properties_test.go:231:4: declared and not used: originalTitle
--- FAIL: TestNewTracingProvider (0.04s)
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
    logger.go:146: 2025-10-17T15:35:41.133-0300 INFO    Distributed tracing is disabled
    logger.go:146: 2025-10-17T15:35:41.172-0300 INFO    Shutting down tracing provider
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
    logger.go:146: 2025-10-17T15:35:41.172-0300 INFO    Distributed tracing is disabled
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
FAIL    github.com/vertikon/mcp-ultra/internal/telemetry        0.789s
        github.com/vertikon/mcp-ultra/internal/testhelpers              coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/internal/tracing          coverage: 0.0% of statements
        github.com/vertikon/mcp-ultra/scripts           coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/test/component [build failed]
        github.com/vertikon/mcp-ultra/test/mocks                coverage: 0.0% of statements
FAIL    github.com/vertikon/mcp-ultra/test/property [build failed]
ok      github.com/vertikon/mcp-ultra/tests/smoke       0.416s  coverage: [no statements]
FAIL
PS E:\vertikon\business\SaaS\templates\mcp-ultra>
PS E:\vertikon\business\SaaS\templates\mcp-ultra> # 4. Validar novamente
PS E:\vertikon\business\SaaS\templates\mcp-ultra> vertikon-validator run --project mcp-ultra
vertikon-validator : O termo 'vertikon-validator' não é reconhecido como nome de cmdlet, função, arquivo de script ou programa
operável. Verifique a grafia do nome ou, se um caminho tiver sido incluído, veja se o caminho está correto e tente novamente.
No linha:1 caractere:1
+ vertikon-validator run --project mcp-ultra
+ ~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : ObjectNotFound: (vertikon-validator:String) [], CommandNotFoundException
    + FullyQualifiedErrorId : CommandNotFoundException

PS E:\vertikon\business\SaaS\templates\mcp-ultra>
