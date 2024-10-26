package main

import (
	"github.com/babakDoraniArab/todo-cli/application"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)

func main() {

	app, err := application.Setup()

	helper.CheckErr(err, "could not setup application")

	app.AssigneeService.SeedAssignees()
	app.AssigneeService.ShowAll()
	app.TodoService.Seedtodos()
	app.TodoService.ShowAll()

}
