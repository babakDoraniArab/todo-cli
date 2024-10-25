package domain

import (
	"time"


)


type Todo struct {
	Title      string    `json:"title"`
	AssigneeID int       `json:"assignee_id"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TodoList []Todo

