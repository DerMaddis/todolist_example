package database

import "github.com/dermaddis/todolist_example/internal/models"

type Database interface {
	GetTodos() []models.Todo
    GetTodoById(id int) (models.Todo, error)
    TodoExists(id int) bool
	NumTodos() int
	AddTodo(title string) error
	UpdateTodo(id int, title string, completed bool) error
	// todo: Delete
}

