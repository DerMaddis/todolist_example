package inmem

import (
	"fmt"
	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/models"
)

type InmemDatabase struct {
	Todos []models.Todo
}

func New() InmemDatabase {
	return InmemDatabase{
		Todos: []models.Todo{
			{Id: 0, Title: "Todo 1", Completed: false},
			{Id: 1, Title: "Todo 2", Completed: true},
			{Id: 2, Title: "Todo 3", Completed: false},
		},
	}
}

func (d *InmemDatabase) GetTodos() []models.Todo {
	return d.Todos
}

func (d *InmemDatabase) NumTodos() int {
	return len(d.Todos)
}

func (d *InmemDatabase) TodoExists(id int) bool {
	return id >= 0 && id < d.NumTodos()
}

func (d *InmemDatabase) GetTodoById(id int) (models.Todo, error) {
	if !d.TodoExists(id) {
		return models.Todo{}, fmt.Errorf("GetTodoById: %w", errs.ErrorNotFound)
	}
	return d.Todos[id], nil
}

func (d *InmemDatabase) AddTodo(title string) error {
	d.Todos = append(d.Todos, models.Todo{Id: len(d.Todos), Title: title, Completed: false})
    return nil
}

func (d *InmemDatabase) UpdateTodo(id int, title string, completed bool) error {
	if !d.TodoExists(id) {
		return fmt.Errorf("UpdateTodo: %w", errs.ErrorNotFound)
	}

	updated := models.Todo{Id: id, Title: title, Completed: completed}
	d.Todos[updated.Id] = updated
	return nil
}
