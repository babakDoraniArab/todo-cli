package main

import (
	"github.com/babakDoraniArab/todo-cli/application"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
	"github.com/babakDoraniArab/todo-cli/presentation/cli"
)

func main() {

	app, err := application.Setup()

	helper.CheckErr(err, "could not setup application")

	cliService := cli.NewCliService(app.TodoService, app.AssigneeService)
	cli.Execute(cliService)
}
