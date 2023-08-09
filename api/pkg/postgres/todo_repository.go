package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/ent/todo"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type TodoRepository struct {
	conn *Connection
}

func NewTodoRepository(conn *Connection) *TodoRepository {
	return &TodoRepository{
		conn: conn,
	}
}

func (r *TodoRepository) Create(ctx context.Context, td *ent.Todo) (*ent.Todo, error) {
	created, err := r.conn.txClient(ctx).Todo.
		Create().
		SetOwnerID(td.OwnerID).
		SetTask(td.Task).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *TodoRepository) FindByID(ctx context.Context, id int) (*ent.Todo, error) {
	found, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.ID(id)).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, pkgerror.Known{
			Code:   pkgerror.CodeNotFound,
			Simple: fmt.Errorf("todo not found with ID(%d)", id),
		}
	case err != nil:
		return nil, err
	}
	return found, nil
}

func (r *TodoRepository) FindAllByOwnerIDOrderByCreateTimeAsc(ctx context.Context, id uuid.UUID) ([]*ent.Todo, error) {
	todos, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.OwnerID(id)).
		Order(todo.ByCreateTime()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) CountByOwnerID(ctx context.Context, id uuid.UUID) (int, error) {
	count, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.OwnerID(id)).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TodoRepository) DeleteByID(ctx context.Context, id int) error {
	err := r.conn.txClient(ctx).Todo.
		DeleteOneID(id).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return pkgerror.Known{
			Code:   pkgerror.CodeNotFound,
			Origin: err,
			Simple: fmt.Errorf("todo with id(%d) does not exist", id),
		}
	case err != nil:
		return err
	}
	return nil
}
