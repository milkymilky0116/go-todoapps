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
		Id:         id,
		Context:    context,
		IsComplete: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	s.DB.Todos = append(s.DB.Todos, newTodo)

	return newTodo
}

func (s *TodoService) Find(id int) *db.Todo {
	for _, todo := range s.DB.Todos {
		if todo.Id == id {
			return &todo
		}
	}
	return nil
}

func (s *TodoService) Delete(id int) *db.Todo {
	for index, todo := range s.DB.Todos {
		if todo.Id == id {
			s.DB.Todos = append(s.DB.Todos[:index], s.DB.Todos[index+1:]...)
			return &todo
		}
	}
	return nil
}

func (s *TodoService) ToggleComplete(id int) *db.Todo {
	for index, todo := range s.DB.Todos {
		if todo.Id == id {
			s.DB.Todos[index].IsComplete = !todo.IsComplete
			return &s.DB.Todos[index]
		}
	}
	return nil
}

func (s *TodoService) ToggleEdit(id int) *db.Todo {
	for index, todo := range s.DB.Todos {
		if todo.Id == id {
			s.DB.Todos[index].IsEditable = !todo.IsEditable
			return &s.DB.Todos[index]
		}
	}
	return nil
}

func (s *TodoService) Update(id int, context string) *db.Todo {
	for index, todo := range s.DB.Todos {
		if todo.Id == id {
			s.DB.Todos[index].Context = context
			return &s.DB.Todos[index]
		}
	}

	return nil
}
