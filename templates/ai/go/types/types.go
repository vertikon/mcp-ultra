package types

import "time"

// UseCase define os casos de uso de IA
type UseCase string

const (
	UseCaseClassification UseCase = "classification"
	UseCaseGeneration     UseCase = "generation"
	UseCaseRerank         UseCase = "rerank"
	UseCaseSummarize      UseCase = "summarize"
)

// AIMode define os modos de operacao
type AIMode string

const (
	AIModeBalanced   AIMode = "balanced"
	AIModeStrict     AIMode = "strict"
	AIModeAggressive AIMode = "aggressive"
)

// Provider define provedores de IA
type Provider string

const (
	ProviderOpenAI Provider = "openai"
	ProviderQwen   Provider = "qwen"
	ProviderLocal  Provider = "local"
)

// Context contem informacoes de contexto para requisicoes IA
type Context struct {
	TenantID  string
	MCPID     string
	SDKName   string
	Timestamp time.Time
}

// InferenceRequest representa uma solicitacao de inferencia
type InferenceRequest struct {
	Context   Context
	UseCase   UseCase
	Prompt    string
	MaxTokens int
	Timeout   time.Duration
}

// InferenceResponse representa a resposta de uma inferencia
type InferenceResponse struct {
	Content   string
	TokensIn  int
	TokensOut int
	LatencyMS int64
	CostBRL   float64
	Cached    bool
	Provider  Provider
	Model     string
}

// RouterDecision representa uma decisao do router
type RouterDecision struct {
	Context   Context
	UseCase   UseCase
	Provider  Provider
	Model     string
	Reason    string
	Timestamp time.Time
}

// PolicyBlock representa um bloqueio de policy
type PolicyBlock struct {
	Context   Context
	Rule      string
	Severity  string
	Sample    string
	Timestamp time.Time
}

// BudgetBreach representa uma violacao de budget
type BudgetBreach struct {
	Scope     string // global|tenant|mcp
	TenantID  string
	MCPID     string
	CapBRL    float64
	UsedBRL   float64
	Action    string // degrade|block|alert_only
	Timestamp time.Time
}
