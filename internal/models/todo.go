package models

import "fmt"

type Todo struct {
	Title     string `db:"title"`
	Completed bool   `db:"completed"`
	Id        int    `db:"id"`
}

func (t Todo) String() string {
	return fmt.Sprintf("{ Title: %q, Completed: %t, Id: %d }", t.Title, t.Completed, t.Id)
}
