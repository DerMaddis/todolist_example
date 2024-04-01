package postgres

import (
	"fmt"
	"log"

	"github.com/dermaddis/todolist_example/internal/config"
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
		log.Fatalln("rror connecting to postgres", err)
	}
	return PostgresDatabase{
		db: db,
	}
}

func (p *PostgresDatabase) GetTodos() ([]models.Todo, error) {
	todos := []models.Todo{}
	err := p.db.Select(&todos, "SELECT * FROM todos ORDER BY id")
	if err != nil {
		return []models.Todo{}, fmt.Errorf("GetTodos: %w", err)
	}
	return todos, nil
}

func (p *PostgresDatabase) NumTodos() (int, error) {
	count := 0
	err := p.db.Get(&count, "SELECT COUNT(*) FROM todos")
	if err != nil {
		return 0, fmt.Errorf("NumTodos: %w", err)
	}
	return count, nil
}

func (p *PostgresDatabase) TodoExists(id int) (bool, error) {
	numTodos, err := p.NumTodos()
	if err != nil {
		return false, fmt.Errorf("TodoExists: %w", err)
	}
	return id >= 0 && id < numTodos, nil
}

func (p *PostgresDatabase) GetTodoById(id int) (models.Todo, error) {
	var todo models.Todo
	err := p.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return models.Todo{}, fmt.Errorf("GetTodoById: %w", err)
	}
	return todo, nil
}

func (p *PostgresDatabase) AddTodo(title string) error {
	_, err := p.db.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", title, false)
	if err != nil {
		return fmt.Errorf("AddTodo: %w", err)
	}
	return nil
}

func (p *PostgresDatabase) UpdateTodo(id int, title string, completed bool) error {
	_, err := p.db.Exec("UPDATE todos SET title = $1, completed = $2 WHERE id = $3", title, completed, id)
	if err != nil {
		return fmt.Errorf("UpdateTodo: %w", err)
	}
	return nil
}
