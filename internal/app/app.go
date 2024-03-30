package app

import (
	"github.com/dermaddis/todolist_example/internal/database/inmem"
	"github.com/dermaddis/todolist_example/internal/services"
	"github.com/dermaddis/todolist_example/internal/transport"

	"github.com/labstack/echo/v4"
)

func Run() {
    e := echo.New()
    database := inmem.New()
    service := services.New(&database)
    handler := transport.New(service)

    handler.RegisterRoutes(e)

    e.Start(":3000")
}
