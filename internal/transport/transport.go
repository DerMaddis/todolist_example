package transport

import (
	"github.com/dermaddis/todolist_example/internal/services"

	"github.com/labstack/echo/v4"
)

type Handler struct{
    service *services.Service
}

func New(service *services.Service) Handler {
	return Handler{
        service: service,
    }
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.getIndex)
	e.POST("/todo", h.postTodo)
	e.POST("/todo/:id", h.postTodoId)
}
