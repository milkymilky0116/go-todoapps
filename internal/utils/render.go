package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, component templ.Component, status int) error {
	ctx.Response().Status = status
	return component.Render(ctx.Request().Context(), ctx.Response())
}
