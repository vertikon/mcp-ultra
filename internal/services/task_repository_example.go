package services

import (
	"context"
)

type TaskFilters struct {
	TenantKey string
	Limit int
	Offset int
}

type Task struct {
	ID string
	Title string
}

type TaskRepository interface {
	// v11.1 — assinatura final
	List(ctx context.Context, filter TaskFilters) ([]*Task, int, error)
	Create(ctx context.Context, t *Task) (string, error)
	Exists(ctx context.Context, id string) (bool, error) // necessário para alguns testes/mocks
}