package port

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/transaction"
)

//go:generate mockgen -package portmock -destination portmock/mock_repository.go github.com/isutare412/bloated/api/pkg/core/port TransactionManager,UserRepository,IPRepository,TodoRepository

type TransactionManager interface {
	WithTx(context.Context) (transaction.Context, error)
	WithTxOption(context.Context, *sql.TxOptions) (transaction.Context, error)
}

type UserRepository interface {
	FindByID(context.Context, uuid.UUID) (*ent.User, error)
	Upsert(context.Context, *ent.User) (*ent.User, error)
}

type IPRepository interface {
	CreateAll(context.Context, []*ent.BannedIP) ([]*ent.BannedIP, error)
	FindAllOrderByCountryAscIPAsc(context.Context) ([]*ent.BannedIP, error)
}

type TodoRepository interface {
	Create(context.Context, *ent.Todo) (*ent.Todo, error)
	FindByID(ctx context.Context, id int) (*ent.Todo, error)
	FindAllByOwnerIDOrderByCreateTimeAsc(ctx context.Context, id uuid.UUID) ([]*ent.Todo, error)
	CountByOwnerID(ctx context.Context, id uuid.UUID) (int, error)
	DeleteByID(ctx context.Context, id int) error
}
