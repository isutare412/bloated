package port

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
)

//go:generate mockgen -package portmock -destination portmock/mock_service.go github.com/isutare412/bloated/api/pkg/core/port TodoService

type TodoService interface {
	TodosOfUser(ctx context.Context, userID string) ([]*ent.Todo, error)
	AddTodo(context.Context, *ent.Todo) (*ent.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}
