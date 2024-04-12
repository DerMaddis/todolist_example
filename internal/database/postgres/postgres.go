package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dermaddis/todolist_example/internal/config"
	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	db *sqlx.DB
}

func New() PostgresDatabase {
	// loading env
	host := config.LoadEnvVar("host")
	port := config.LoadEnvVar("port")
	user := config.LoadEnvVar("user")
	password := config.LoadEnvVar("password")
	database := config.LoadEnvVar("database")
	sslmode := config.LoadEnvVar("sslmode")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=%s", host, port, user, password, database, sslmode))
	if err != nil {
		log.Fatalln("error connecting to postgres", err)
	}
	return PostgresDatabase{
		db: db,
	}
}

func (p *PostgresDatabase) GetTodos() ([]models.Todo, error) {
	todos := []models.Todo{}
	err := p.db.Select(&todos, "SELECT * FROM todos ORDER BY created ASC")
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

func (p *PostgresDatabase) NumTodos() (int, error) {
	count := 0
	err := p.db.Get(&count, "SELECT COUNT(*) FROM todos")
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PostgresDatabase) GetTodoById(id int) (models.Todo, error) {
	var todo models.Todo
	err := p.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, errs.ErrorNotFound
		}
		return models.Todo{}, err
	}
	return todo, nil
}

func (p *PostgresDatabase) AddTodo(title string) error {
	_, err := p.db.Exec("INSERT INTO todos (title, completed, created) VALUES ($1, $2, $3)", title, false, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDatabase) UpdateTodo(id int, title string, completed bool) error {
	res, err := p.db.Exec("UPDATE todos SET title = $1, completed = $2 WHERE id = $3", title, completed, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errs.ErrorNotFound
	}
	return nil
}

func (p *PostgresDatabase) DeleteTodo(id int) error {
	res, err := p.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errs.ErrorNotFound
	}
	return nil
}
