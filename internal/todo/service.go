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

func (s *TodoService) Add() db.Todo {
	newTodo := db.Todo{
		Id:        5,
		Context:   "this is new todo",
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
