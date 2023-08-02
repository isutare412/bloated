package port

import (
	"context"
	"database/sql"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/transaction"
)

//go:generate mockgen -package portmock -destination portmock/mock_repository.go github.com/isutare412/bloated/api/pkg/core/port TransactionManager,IPRepository,TodoRepository

type TransactionManager interface {
	WithTx(context.Context) (transaction.Context, error)
	WithTxOption(context.Context, *sql.TxOptions) (transaction.Context, error)
}

type IPRepository interface {
	CreateAll(context.Context, []*ent.BannedIP) ([]*ent.BannedIP, error)
	FindAll(context.Context) ([]*ent.BannedIP, error)
}

type TodoRepository interface {
	Create(context.Context, *ent.Todo) (*ent.Todo, error)
	FindAllByUserID(ctx context.Context, id string) ([]*ent.Todo, error)
	DeleteByID(ctx context.Context, id int) error
}
