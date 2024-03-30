package handlers

import (
	"net/http"
	"github.com/dermaddis/todolist_example/templates"
	"github.com/dermaddis/todolist_example/todo"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", getIndex)
	e.POST("/todo", postTodo)
	e.POST("/todo/:id", postTodoId)
}

var todos = []todo.Todo{
	{Id: 0, Title: "Todo 1", Completed: false},
	{Id: 1, Title: "Todo 2", Completed: true},
	{Id: 2, Title: "Todo 3", Completed: false},
}

func getIndex(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Index(todos))
}

func postTodo(c echo.Context) error {
	var todo todo.Todo
	if err := c.Bind(&todo); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	todo.Id = len(todos)

	todos = append(todos, todo)
	return Render(c, http.StatusOK, templates.TodoList(todos))
}

func postTodoId(c echo.Context) error {
	var updatedTodo todo.Todo
	if err := c.Bind(&updatedTodo); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if updatedTodo.Id >= len(todos) {
		return c.String(http.StatusNotFound, "not found")
	}

	updatedTodo.Completed = updatedTodo.CompletedString == "on"
	todos[updatedTodo.Id] = updatedTodo

	return Render(c, http.StatusOK, templates.Todo(updatedTodo))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
