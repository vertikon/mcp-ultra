// Package domain contém DTOs mínimos exigidos por handlers e testes.
package domain

import "github.com/vertikon/mcp-ultra/pkg/types"

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
	ID    types.UUID
	Title string
}

type TaskList struct {
	Items []*Task
	Total int
}
