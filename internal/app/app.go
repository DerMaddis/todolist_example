package app

import (
	"log"

	"github.com/dermaddis/todolist_example/internal/config"
	"github.com/dermaddis/todolist_example/internal/database"
	"github.com/dermaddis/todolist_example/internal/database/inmem"
	"github.com/dermaddis/todolist_example/internal/database/postgres"
	"github.com/dermaddis/todolist_example/internal/services"
	"github.com/dermaddis/todolist_example/internal/transport"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func Run() {
	godotenv.Load()

	e := echo.New()

	dbMode := config.LoadEnvVar("db_mode")
	var database database.Database

	switch dbMode {
	case "postgres":
        log.Println("Using postgres")
		postgres := postgres.New()
		database = &postgres
	case "inmem":
        log.Println("Using inmem")
		inmem := inmem.New()
		database = &inmem
	default:
		log.Fatalln("specify valid db_mode in env")
	}

    service := services.New(database)
	handler := transport.New(service)

	handler.RegisterRoutes(e)

	e.Start(":3000")
}
