package todo

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type Service struct {
	txManager port.TransactionManager
	todoRepo  port.TodoRepository

	maxTodoCountPerUser int
}

func NewService(cfg Config, txMgr port.TransactionManager, todoRepo port.TodoRepository) *Service {
	return &Service{
		txManager: txMgr,
		todoRepo:  todoRepo,

		maxTodoCountPerUser: cfg.MaxTodoCountPerUser,
	}
}

func (s *Service) TodosOfUser(ctx context.Context, userID uuid.UUID) ([]*ent.Todo, error) {
	todos, err := s.todoRepo.FindAllByOwnerIDOrderByCreateTimeAsc(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("finding all todos by user ID(%s): %w", userID, err)
	}

	return todos, nil
}

func (s *Service) AddTodo(ctx context.Context, todo *ent.Todo) (created *ent.Todo, err error) {
	if err := validateTodo(todo); err != nil {
		return nil, fmt.Errorf("validating todo: %w", err)
	}

	tx, err := s.txManager.WithTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting tx: %w", err)
	}
	defer func() { tx.RollbackIfError(err) }()

	todoCount, err := s.todoRepo.CountByOwnerID(tx.Ctx, todo.OwnerID)
	if err != nil {
		return nil, fmt.Errorf("counting todo of user '%s': %w", todo.UserID, err)
	}
	if todoCount >= s.maxTodoCountPerUser {
		return nil, pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("cannot create more than %d todos", s.maxTodoCountPerUser),
		}
	}

	created, err = s.todoRepo.Create(tx.Ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("creating todo: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commiting transaction: %w", err)
	}

	return created, nil
}

func (s *Service) DeleteTodo(ctx context.Context, id int, requesterID uuid.UUID) error {
	tx, err := s.txManager.WithTx(ctx)
	if err != nil {
		return fmt.Errorf("starting tx: %w", err)
	}
	defer func() { tx.RollbackIfError(err) }()

	foundTodo, err := s.todoRepo.FindByID(tx.Ctx, id)
	if err != nil {
		return fmt.Errorf("finding todo: %w", err)
	}
	if foundTodo.OwnerID != requesterID {
		return pkgerror.Known{
			Code:   pkgerror.CodeForbidden,
			Simple: fmt.Errorf("cannot delete todo of another user"),
		}
	}

	if err := s.todoRepo.DeleteByID(tx.Ctx, id); err != nil {
		return fmt.Errorf("deleting todo by ID(%d): %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}
