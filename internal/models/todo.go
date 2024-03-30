package models

import "fmt"

type Todo struct {
	Title           string
	Completed       bool
	CompletedString string
	Id              int
}

func (t Todo) String() string {
	return fmt.Sprintf("{ Title: %q, Completed: %t, Id: %d }", t.Title, t.Completed, t.Id)
}
