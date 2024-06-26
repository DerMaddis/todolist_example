package inmem

import (
	"slices"
	"time"

	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/models"
)

type InmemDatabase struct {
	todos  map[int]models.Todo
	nextId int
}

func New() InmemDatabase {
	now := time.Now()
	return InmemDatabase{
		todos: map[int]models.Todo{
			0: {Id: 0, Title: "Todo 1", Completed: false, Created: now},
			1: {Id: 1, Title: "Todo 2", Completed: true, Created: now},
			2: {Id: 2, Title: "Todo 3", Completed: false, Created: now},
		},
		nextId: 3,
	}
}

func (d *InmemDatabase) todoExists(id int) bool {
	return 0 <= id && id < d.nextId
}

func (d *InmemDatabase) GetTodos() ([]models.Todo, error) {
	todos := make([]models.Todo, 0, len(d.todos))
	for _, v := range d.todos {
		todos = append(todos, v)
	}

	slices.SortFunc(todos, func(i, j models.Todo) int {
		return i.Created.Compare(j.Created)
	})

	return todos, nil
}

func (d *InmemDatabase) NumTodos() (int, error) {
	return len(d.todos), nil
}

func (d *InmemDatabase) GetTodoById(id int) (models.Todo, error) {
	todo, ok := d.todos[id]
	if !ok {
		return models.Todo{}, errs.ErrorNotFound
	}
	return todo, nil
}

func (d *InmemDatabase) AddTodo(title string) error {
	d.todos[d.nextId] = models.Todo{Id: d.nextId, Title: title, Completed: false, Created: time.Now()}
	d.nextId++
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

func (d *InmemDatabase) DeleteTodo(id int) error {
	exists := d.todoExists(id)
	if !exists {
		return errs.ErrorNotFound
	}

	delete(d.todos, id)
	return nil
}
