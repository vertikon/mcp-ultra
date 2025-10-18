//go:build example || testpatch
// +build example testpatch

// Arquivo EXEMPLO de assinatura do router para alinhar com os testes.
// Não compila em produção por causa da build tag acima.
package http

import (
	"net/http"

	"github.com/vertikon/mcp-ultra/internal/services"
)

// HealthService é um alias semântico para o checker de saúde.
type HealthService interface {
	Status() services.HealthStatus
}

// TaskService representa os métodos realmente usados no router pelos testes.
// Adicione aqui apenas os métodos que os testes invocam.
type TaskService interface {
	// Exemplo: Create, List, etc. (adapte ao seu projeto real)
}

// NewRouter define uma assinatura canônica a ser usada nos testes.
// Ajuste a implementação real no seu router principal.
func NewRouter(
	taskSvc TaskService,
	healthSvc HealthService,
) http.Handler {
	return http.NewServeMux() // troque pelo seu router real (ex.: chi/httpx)
}
