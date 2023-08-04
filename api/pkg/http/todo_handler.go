package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/bloated/api/pkg/core/port"
)

type todoHandler struct {
	queryGetter *queryGetter
	todoService port.TodoService
}

func newTodoHandler(queryGetter *queryGetter, todoService port.TodoService) *todoHandler {
	return &todoHandler{
		queryGetter: queryGetter,
		todoService: todoService,
	}
}

func (h *todoHandler) router() http.Handler {
	jsonContent := middleware.AllowContentType("application/json")

	r := chi.NewRouter()
	r.Post("/", jsonContent(http.HandlerFunc(h.createTodo)).ServeHTTP)
	r.Get("/", h.listTodos)

	return r
}

func (h *todoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := req.validate(); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
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

	userID, err := h.queryGetter.userID(r)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting userID query: %w", err))
		return
	}

	todos, err := h.todoService.TodosOfUser(ctx, userID)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting todos of user %s: %w", userID, err))
		return
	}

	var resp listTodosReponse
	resp.fromEntities(todos)
	responseJSON(w, r, &resp)
}
