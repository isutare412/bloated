package todo

import (
	"context"
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/port"
)

type Service struct {
	txMgr    port.TransactionManager
	todoRepo port.TodoRepository
}

func NewService(txMgr port.TransactionManager, todoRepo port.TodoRepository) *Service {
	return &Service{
		txMgr:    txMgr,
		todoRepo: todoRepo,
	}
}

func (s *Service) TodosOfUser(ctx context.Context, userID string) ([]*ent.Todo, error) {
	todos, err := s.todoRepo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("finding all todos by user ID(%s): %w", userID, err)
	}

	return todos, nil
}

func (s *Service) AddTodo(ctx context.Context, todo *ent.Todo) (created *ent.Todo, err error) {
	if err := validateTodo(todo); err != nil {
		return nil, fmt.Errorf("validating todo: %w", err)
	}

	tx, err := s.txMgr.WithTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting tx: %w", err)
	}
	defer func() { tx.RollbackIfError(err) }()

	created, err = s.todoRepo.Create(tx.Ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("creating todo: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commiting transaction: %w", err)
	}

	return created, nil
}

func (s *Service) DeleteTodo(ctx context.Context, id int) error {
	tx, err := s.txMgr.WithTx(ctx)
	if err != nil {
		return fmt.Errorf("starting tx: %w", err)
	}
	defer func() { tx.RollbackIfError(err) }()

	if err := s.todoRepo.DeleteByID(tx.Ctx, id); err != nil {
		return fmt.Errorf("deleting todo by ID(%d): %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}
