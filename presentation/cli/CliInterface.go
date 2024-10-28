package cli

import "github.com/babakDoraniArab/todo-cli/domain"

type CliInterface interface {
	//******************** these are related to Assignee Service
	GetAllAssignee() (domain.AssigneeList, error)
	ShowAllAssignee() error
	AddAssignee(name string, email string) error
	EditAssignee(index int, name string, email string) error
	DeleteAssignee(index int) error
	SeedAssignees()

	//******************** these are related to the TotoService
	ShowAllTodo()
	AddTodo(title string, assigneeID int) error
	EditTodo(index int, title string, assigneeID int, status bool) error
	DeleteTodo(index int) error
	SeedTodos()
}
