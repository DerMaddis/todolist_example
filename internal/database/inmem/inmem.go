package inmem

import (
	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/models"
)

type InmemDatabase struct {
	todos []models.Todo
}

func New() InmemDatabase {
	return InmemDatabase{
		todos: []models.Todo{
			{Id: 0, Title: "Todo 1", Completed: false},
			{Id: 1, Title: "Todo 2", Completed: true},
			{Id: 2, Title: "Todo 3", Completed: false},
		},
	}
}

func (d *InmemDatabase) todoExists(id int) bool {
	return 0 <= id && id < len(d.todos)
}

func (d *InmemDatabase) GetTodos() ([]models.Todo, error) {
	return d.todos, nil
}

func (d *InmemDatabase) NumTodos() (int, error) {
	return len(d.todos), nil
}

func (d *InmemDatabase) GetTodoById(id int) (models.Todo, error) {
	return d.todos[id], nil
}

func (d *InmemDatabase) AddTodo(title string) error {
	d.todos = append(d.todos, models.Todo{Id: len(d.todos), Title: title, Completed: false})
	return nil
}

func (d *InmemDatabase) UpdateTodo(id int, title string, completed bool) error {
	exists := d.todoExists(id)
	if !exists {
		return errs.ErrorNotFound
	}

	updated := models.Todo{Id: id, Title: title, Completed: completed}
	d.todos[updated.Id] = updated
	return nil
}
