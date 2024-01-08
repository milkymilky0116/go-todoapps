package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/milkymilky0116/go-todoapps/internal/db"
	"github.com/milkymilky0116/go-todoapps/internal/types"
)

type StateMiddleware struct {
	DB *db.DB
}

func (m *StateMiddleware) StateLoading(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newState := types.State{
			Length: len(m.DB.Todos),
		}
		c.Set("state", &newState)
		return next(c)
	}
}

func NewStateLoaderMiddleware(db *db.DB) *StateMiddleware {
	return &StateMiddleware{
		DB: db,
	}
}
