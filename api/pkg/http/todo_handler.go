package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	"github.com/isutare412/bloated/api/pkg/contextbag"
	"github.com/isutare412/bloated/api/pkg/core/constant"
	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
	"github.com/isutare412/bloated/api/pkg/validation"
)

type todoHandler struct {
	validator   *validation.Validator
	pathGetter  *pathGetter
	queryGetter *queryGetter
	todoService port.TodoService
}

func newTodoHandler(
	validator *validation.Validator,
	pathGetter *pathGetter,
	queryGetter *queryGetter,
	todoService port.TodoService,
) *todoHandler {
	return &todoHandler{
		validator:   validator,
		pathGetter:  pathGetter,
		queryGetter: queryGetter,
		todoService: todoService,
	}
}

func (h *todoHandler) registerRoutes(r chi.Router) {
	jsonContent := middleware.AllowContentType("application/json")

	r.Post("/todos", jsonContent(http.HandlerFunc(h.createTodo)).ServeHTTP)
	r.Get("/todos", h.listTodos)
	r.Delete("/todos/{todoID}", h.deleteTodo)
}

func (h *todoHandler) extractCustomToken(ctx context.Context) (model.CustomToken, error) {
	token, ok := contextbag.Bag(ctx).Get(constant.BagKeyCustomToken).(model.CustomToken)
	if !ok {
		return model.CustomToken{}, pkgerror.Known{
			Code:   pkgerror.CodeUnauthorized,
			Simple: fmt.Errorf("need token"),
		}
	}
	return token, nil
}

func (h *todoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := h.extractCustomToken(ctx)
	if err != nil {
		responseError(w, r, fmt.Errorf("extracting custom token: %w", err))
		return
	}

	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := h.validator.Validate(&req); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
		return
	}
	if token.UserID != req.UserID {
		responseError(w, r, pkgerror.Known{
			Code:   pkgerror.CodeForbidden,
			Simple: fmt.Errorf("cannot create todo of other user"),
		})
		return
	}

	todo, err := h.todoService.AddTodo(ctx, req.toEntity())
	if err != nil {
		responseError(w, r, fmt.Errorf("adding todo: %w", err))
		return
	}

	var resp createTodoResponse
	resp.fromEntity(todo)
	responseJSON(w, r, &resp)
}

func (h *todoHandler) listTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := h.extractCustomToken(ctx)
	if err != nil {
		responseError(w, r, fmt.Errorf("extracting custom token: %w", err))
		return
	}

	userID, err := h.queryGetter.userID(r)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting query param: %w", err))
		return
	}

	if userID != token.UserID {
		responseError(w, r, pkgerror.Known{
			Code:   http.StatusForbidden,
			Simple: fmt.Errorf("cannot list todos of other user"),
		})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		responseError(w, r, pkgerror.Known{
			Code:   http.StatusBadRequest,
			Simple: fmt.Errorf("user ID is not a valid UUID"),
		})
		return
	}

	todos, err := h.todoService.TodosOfUser(ctx, userUUID)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting todos of user %s: %w", userID, err))
		return
	}

	var resp listTodosReponse
	resp.fromEntities(todos)
	responseJSON(w, r, &resp)
}

func (h *todoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := h.extractCustomToken(ctx)
	if err != nil {
		responseError(w, r, fmt.Errorf("extracting custom token: %w", err))
		return
	}

	todoID, err := h.pathGetter.todoID(r)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting path param: %w", err))
		return
	}

	userID, err := uuid.Parse(token.UserID)
	if err != nil {
		responseError(w, r, pkgerror.Known{
			Code:   http.StatusBadRequest,
			Simple: fmt.Errorf("user ID is not a valid UUID"),
		})
		return
	}

	if err := h.todoService.DeleteTodo(ctx, todoID, userID); err != nil {
		responseError(w, r, fmt.Errorf("deleting todo(%d): %w", todoID, err))
		return
	}

	responseStatus(w, r, http.StatusOK)
}
