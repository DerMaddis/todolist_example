package transport

import (
	"errors"
	"log"
	"net/http"

	"github.com/dermaddis/todolist_example/internal/errs"
	"github.com/dermaddis/todolist_example/internal/templates"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getIndex(c echo.Context) error {
	todos, err := h.service.GetTodos()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return Render(c, http.StatusOK, templates.Index(todos))
}

type GetTodoById struct {
	Id *int `param:"id" validate:"required"`
}

func (h *Handler) getTodoById(c echo.Context) error {
	var data GetTodoById
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	todo, err := h.service.GetTodoById(*data.Id)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return Render(c, http.StatusOK, templates.Todo(todo))
}

type PostTodo struct {
	Title string `form:"title" validate:"required"`
}

func (h *Handler) postTodo(c echo.Context) error {
	var data PostTodo
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := h.service.AddTodo(data.Title)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	todos, err := h.service.GetTodos()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return Render(c, http.StatusOK, templates.TodoList(todos))
}

type PostTodoId struct {
	Id              *int    `param:"id" validate:"required"`
	Title           string  `form:"title" validate:"required"`
	CompletedString *string `form:"completed_string"`
}

func (h *Handler) postTodoId(c echo.Context) error {
	var data PostTodoId
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := h.service.UpdateTodo(*data.Id, data.Title, data.CompletedString != nil)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	todo, err := h.service.GetTodoById(*data.Id)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return Render(c, http.StatusOK, templates.Todo(todo))
}

type DeleteTodoId struct {
	Id *int `param:"id" validate:"required"`
}

func (h *Handler) deleteTodoId(c echo.Context) error {
	var data DeleteTodoId
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := h.service.DeleteTodo(*data.Id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, errs.ErrorNotFound) {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusOK, "") // Replaces the todo that was deleted
}
