package transport

import (
	"errors"
	"net/http"
	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/templates"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getIndex(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Index(h.service.GetTodos()))
}

type PostTodo struct {
	Title string `form:"title"`
}

func (h *Handler) postTodo(c echo.Context) error {
	var data PostTodo
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	h.service.AddTodo(data.Title)

	return Render(c, http.StatusOK, templates.TodoList(h.service.GetTodos()))
}

type PostTodoId struct {
	Id              int    `form:"id"`
	Title           string `form:"title"`
	CompletedString string `form:"completed_string"`
}

func (h *Handler) postTodoId(c echo.Context) error {
	var data PostTodoId
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := h.service.UpdateTodo(data.Id, data.Title, data.CompletedString == "on")
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	todo, err := h.service.GetTodoById(data.Id)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return Render(c, http.StatusOK, templates.Todo(todo))
}