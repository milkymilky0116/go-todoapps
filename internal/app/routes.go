package app

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/milkymilky0116/go-todoapps/internal/handlers"
	"github.com/milkymilky0116/go-todoapps/internal/todo"
)

func (app *Application) InitRoutes() {
	todoService := todo.TodoService{
		DB: app.DB,
	}
	homeHandler := handlers.NewHomeHandler(todoService)
	app.Handler.Use(middleware.Logger())
	app.Handler.GET("/", homeHandler.Home)
	app.Handler.POST("/add", homeHandler.Add)
	app.Handler.DELETE("/delete", homeHandler.Delete)
}
