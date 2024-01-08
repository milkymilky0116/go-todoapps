package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/milkymilky0116/go-todoapps/internal/db"
)

type Application struct {
	DB      *db.DB
	Port    int
	Handler *echo.Echo
}

func InitApp() {
	app := Application{}
	app.Port = 3000
	app.DB = db.InitDB()
	app.Handler = echo.New()
	app.InitRoutes()
	app.Run()
}

func (app *Application) Run() {
	addr := fmt.Sprintf(":%d", app.Port)

	if err := app.Handler.Start(addr); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
