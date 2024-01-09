package todo

import (
	"time"

	"github.com/milkymilky0116/go-todoapps/internal/db"
)

type TodoService struct {
	DB *db.DB
}

func (s *TodoService) List() []db.Todo {
	return s.DB.Todos
}

func (s *TodoService) Add(id int, context string) db.Todo {
	newTodo := db.Todo{
		Id:        id,
		Context:   context,
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
