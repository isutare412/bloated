package http

import (
	"fmt"
	"time"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type todo struct {
	ID         int       `json:"id"`
	UserID     string    `json:"userId"`
	Task       string    `json:"task"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (t *todo) fromEntity(todo *ent.Todo) {
	t.ID = todo.ID
	t.UserID = todo.UserID
	t.Task = todo.Task
	t.CreateTime = todo.CreateTime
	t.UpdateTime = todo.UpdateTime
}

type createTodoRequest struct {
	UserID string `json:"userId"`
	Task   string `json:"task"`
}

func (req *createTodoRequest) validate() error {
	if req.UserID == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("userId should not be empty"),
		}
	}
	if req.Task == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("task should not be empty"),
		}
	}
	return nil
}

func (req *createTodoRequest) toEntity() *ent.Todo {
	return &ent.Todo{
		UserID: req.UserID,
		Task:   req.Task,
	}
}

type createTodoResponse struct {
	todo
}

func (resp *createTodoResponse) fromEntity(todo *ent.Todo) {
	resp.todo.fromEntity(todo)
}

type listTodosReponse struct {
	Todos []todo `json:"todos"`
}

func (resp *listTodosReponse) fromEntities(todos []*ent.Todo) {
	resp.Todos = make([]todo, len(todos))
	for i, todo := range todos {
		resp.Todos[i].fromEntity(todo)
	}
}
