package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/milkymilky0116/go-todoapps/internal/todo"
	"github.com/milkymilky0116/go-todoapps/internal/utils"
	"github.com/milkymilky0116/go-todoapps/views/components"
	"github.com/milkymilky0116/go-todoapps/views/layout"
)

type HomeHandler struct {
	TodoService todo.TodoService
}

func (h *HomeHandler) Home(ctx echo.Context) error {
	todos := h.TodoService.List()

	layout := layout.Base(todos)

	return utils.Render(ctx, layout)
}

func (h *HomeHandler) Add(ctx echo.Context) error {
	newTodo := h.TodoService.Add()
	todoComponents := components.Card(components.CardProps{
		Todo: newTodo,
	})
	return utils.Render(ctx, todoComponents)
}

func (h *HomeHandler) Delete(ctx echo.Context) error {
	h.TodoService.Delete()

	todos := h.TodoService.List()

	layout := layout.Base(todos)

	return utils.Render(ctx, layout)
}

func NewHomeHandler(service todo.TodoService) *HomeHandler {
	return &HomeHandler{
		TodoService: service,
	}
}
