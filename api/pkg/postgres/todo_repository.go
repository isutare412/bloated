package postgres

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/ent/todo"
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
		SetUserID(td.UserID).
		SetTask(td.Task).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *TodoRepository) FindByUserID(ctx context.Context, id string) ([]*ent.Todo, error) {
	todos, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.UserID(id)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
