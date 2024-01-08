package todo

import (
	"fmt"
	"time"

	"github.com/milkymilky0116/go-todoapps/internal/db"
)

type TodoService struct {
	DB *db.DB
}

func (s *TodoService) List() []db.Todo {
	return s.DB.Todos
}

func (s *TodoService) Add(length int) db.Todo {
	newContext := fmt.Sprintf("This is new todo #%d", length)
	newTodo := db.Todo{
		Id:        5,
		Context:   newContext,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.DB.Todos = append(s.DB.Todos, newTodo)

	return newTodo
}

func (s *TodoService) Delete() db.Todo {
	deletedTodo := s.DB.Todos[len(s.DB.Todos)-1]

	s.DB.Todos = s.DB.Todos[:len(s.DB.Todos)-1]

	return deletedTodo
}
