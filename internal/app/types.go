package app

import (
	"github.com/labstack/echo/v4"
	"github.com/milkymilky0116/go-todoapps/internal/db"
)

type Application struct {
	DB      *db.DB
	Port    int
	Handler *echo.Echo
}
