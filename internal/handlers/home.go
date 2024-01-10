package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/milkymilky0116/go-todoapps/internal/todo"
	"github.com/milkymilky0116/go-todoapps/internal/types"
	"github.com/milkymilky0116/go-todoapps/internal/utils"
	"github.com/milkymilky0116/go-todoapps/views/components"
	"github.com/milkymilky0116/go-todoapps/views/components/card"
	"github.com/milkymilky0116/go-todoapps/views/layout"
)

type HomeHandler struct {
	TodoService todo.TodoService
}

func (h *HomeHandler) Home(ctx echo.Context) error {
	todos := h.TodoService.List()

	for index := range todos {
		todos[index].IsEditable = false
	}

	layout := layout.Base(todos)

	return utils.Render(ctx, layout, http.StatusOK)
}

func (h *HomeHandler) Add(ctx echo.Context) error {
	context := ctx.FormValue("context")
	log.Println(context)
	if context == "" {
		errorMsg := components.ErrorMsg("invalid form value")
		return utils.Render(ctx, errorMsg, http.StatusInternalServerError)
	}
	state, ok := ctx.Get("state").(*types.State)
	if !ok {
		ctx.Error(errors.New("context error"))
	}
	log.Println(state)
	h.TodoService.Add(state.Length+1, context)

	todosComponents := components.Todos(h.TodoService.List())
	return utils.Render(ctx, todosComponents, http.StatusOK)
}

func (h *HomeHandler) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	h.TodoService.Delete(id)
	todos := h.TodoService.List()

	todoBoard := components.Todos(todos)
	return utils.Render(ctx, todoBoard, http.StatusOK)
}

func (h *HomeHandler) ToggleComplete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	h.TodoService.ToggleComplete(id)

	todo := h.TodoService.Find(id)

	todoComponent := card.Card(card.CardProps{
		Todo: *todo,
	})
	return utils.Render(ctx, todoComponent, http.StatusOK)
}

func (h *HomeHandler) ToggleEdit(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	todo := h.TodoService.ToggleEdit(id)

	todoComponent := card.Card(card.CardProps{
		Todo: *todo,
	})

	return utils.Render(ctx, todoComponent, http.StatusOK)
}

func (h *HomeHandler) UpdateContext(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	context := ctx.FormValue("update-context")
	log.Println(context)
	todo := h.TodoService.Update(id, context)
	todo.IsEditable = false
	todo.IsComplete = false
	todoComponent := card.Card(card.CardProps{
		Todo: *todo,
	})

	return utils.Render(ctx, todoComponent, http.StatusOK)
}

func NewHomeHandler(service todo.TodoService) *HomeHandler {
	return &HomeHandler{
		TodoService: service,
	}
}
