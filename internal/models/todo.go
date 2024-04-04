package models

import (
	"fmt"
	"time"
)

type Todo struct {
	Title     string    `db:"title"`
	Completed bool      `db:"completed"`
	Id        int       `db:"id"`
	Created   time.Time `db:"created"`
}

func (t Todo) String() string {
    return fmt.Sprintf("{ Title: %q, Completed: %t, Id: %d, Created: %s }", t.Title, t.Completed, t.Id, t.Created.Format("02.01.2006 15:04:05"))
}
