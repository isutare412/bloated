package todo

import (
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

func validateTodo(todo *ent.Todo) error {
	if todo == nil {
		return fmt.Errorf("todo is nil")
	}

	if todo.Task == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("task should not be empty"),
		}
	}
	return nil
}
