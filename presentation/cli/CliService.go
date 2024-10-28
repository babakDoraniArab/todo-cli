package cli

import (
	"github.com/babakDoraniArab/todo-cli/application/assignee"
	"github.com/babakDoraniArab/todo-cli/application/todo"
	"github.com/babakDoraniArab/todo-cli/domain"
)

type CliService struct {
	TodoService     *todo.TodoService
	AssigneeService *assignee.AssigneeService
}

// I'm starting the service with this constructor
func NewCliService(todoService *todo.TodoService, assigneeService *assignee.AssigneeService) *CliService {
	return &CliService{
		TodoService:     todoService,
		AssigneeService: assigneeService,
	}
}

// //***************** AssigneeSection
func (c *CliService) GetAllAssignee() (domain.AssigneeList, error) {
	return c.AssigneeService.GetAllAssginee()
}
func (c *CliService) ShowAllAssignee() error {
	return c.AssigneeService.ShowAllAssignee()
}
func (c *CliService) AddAssignee(name string, email string) error {
	return c.AssigneeService.AddAssignee(name, email)
}
func (c *CliService) EditAssignee(index int, name string, email string) error {
	return c.AssigneeService.EditAssignee(index, name, email)
}
func (c *CliService) DeleteAssignee(index int) error {
	return c.AssigneeService.DeleteAssignee(index)
}
func (c *CliService) SeedAssignee() {
	c.AssigneeService.SeedAssignees()
}

// //***************** TodoSection
func (c *CliService) ShowAllTodo() {
	c.TodoService.ShowAllTodo()
}
func (c *CliService) AddTodo(title string, assigneeID int) error {
	return c.TodoService.AddTodo(title, assigneeID)
}
func (c *CliService) EditTodo(index int, title string, assigneeID int, status bool) error {
	return c.TodoService.EditTodo(index, title, assigneeID, status)
}
func (c *CliService) DeleteTodo(index int) error {
	return c.TodoService.DeleteTodo(index)
}
func (c *CliService) SeedTodos() {
	c.TodoService.SeedTodos()
}
