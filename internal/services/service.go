package services

import (
	"github.com/dermaddis/todolist_example/internal/database"
	"github.com/dermaddis/todolist_example/internal/models"
)

type Service struct {
	db database.Database
}

func New(db database.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetTodos() ([]models.Todo, error) {
	todos, err := s.db.GetTodos()
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func (s *Service) GetTodoById(id int) (models.Todo, error) {
	todo, err := s.db.GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (s *Service) AddTodo(title string) error {
	err := s.db.AddTodo(title)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTodo(id int, title string, completed bool) error {
	err := s.db.UpdateTodo(id, title, completed)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTodo(id int) error {
    err := s.db.DeleteTodo(id)
    if err != nil {
        return err
    }
    return nil
}
