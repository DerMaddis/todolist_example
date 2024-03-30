package models

import "fmt"

type Todo struct {
	Title           string `json:"title" form:"title"`
	Completed       bool   `json:"completed" form:"completed"`
	CompletedString string `json:"completed_string" form:"completed_string"`
	Id              int    `json:"id" form:"id"`
}

func (t Todo) String() string {
	return fmt.Sprintf("{ Title: %q, Completed: %t, Id: %d }", t.Title, t.Completed, t.Id)
}
