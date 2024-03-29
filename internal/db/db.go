package db

import "time"

type Todo struct {
	Id         int
	Context    string
	IsComplete bool
	IsEditable bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DB struct {
	Todos []Todo
}

func InitDB() *DB {
	todos := []Todo{
		{
			Id:         1,
			Context:    "Test",
			IsComplete: false,
			IsEditable: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         2,
			Context:    "Test2",
			IsComplete: false,
			IsEditable: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         3,
			Context:    "Test3",
			IsComplete: false,
			IsEditable: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Id:         4,
			Context:    "Test4",
			IsComplete: true,
			IsEditable: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	db := DB{
		Todos: todos,
	}
	return &db
}
