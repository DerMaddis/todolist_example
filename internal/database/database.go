package database

import "github.com/dermaddis/todolist_example/internal/models"

type Database interface {
	GetTodos() ([]models.Todo, error)
	GetTodoById(id int) (models.Todo, error)
	NumTodos() (int, error)
	AddTodo(title string) error
	UpdateTodo(id int, title string, completed bool) error
	// todo: Delete
}
