package postgres

import (
	"context"
	"fmt"

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
		SetUserID(td.UserID).
		SetTask(td.Task).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *TodoRepository) FindAllByUserID(ctx context.Context, id string) ([]*ent.Todo, error) {
	todos, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.UserID(id)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) CountByUserID(ctx context.Context, id string) (int, error) {
	count, err := r.conn.txClient(ctx).Todo.
		Query().
		Where(todo.UserID(id)).
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
