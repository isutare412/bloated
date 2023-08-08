package port

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/model"
)

//go:generate mockgen -package portmock -destination portmock/mock_service.go github.com/isutare412/bloated/api/pkg/core/port AuthService,TodoService,IPService

type AuthService interface {
	IssueCustomTokenFromGoogle(ctx context.Context, tokenString string) (string, error)
	VerifyCustomToken(ctx context.Context, tokenString string) (model.CustomToken, error)
}

type TodoService interface {
	TodosOfUser(ctx context.Context, userID string) ([]*ent.Todo, error)
	AddTodo(context.Context, *ent.Todo) (*ent.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}

type IPService interface {
	AllBannedIPs(context.Context) ([]*ent.BannedIP, error)
}
