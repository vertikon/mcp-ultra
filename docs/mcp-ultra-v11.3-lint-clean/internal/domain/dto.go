// Package domain contém DTOs mínimos exigidos por handlers e testes.
package domain

// Troque para seu alias de UUID, se usar pkg/types:
// import "github.com/vertikon/mcp-ultra-fix/pkg/types"
// type UUID = types.UUID

import "github.com/google/uuid"

type CreateTaskRequest struct {
	Title       string
	Description string
}

type UpdateTaskRequest struct {
	Title       *string
	Description *string
}

type TaskFilters struct {
	TenantKey string
	Limit     int
	Offset    int
}

type Task struct {
	ID    uuid.UUID
	Title string
}

type TaskList struct {
	Items []*Task
	Total int
}
