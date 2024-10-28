package cli

import (
	"flag"
	"fmt"
	"os"
)

func AddRun(flagset *flag.FlagSet, cliservice *CliService) {
	if len(flag.Args()) < 3 {
		fmt.Printf("you need to add more arguments\n")
		flag.Usage()
		return
	}

	// the second layer

	addTodoCmd := flag.NewFlagSet("todo", flag.ExitOnError)
	addAssigneeCmd := flag.NewFlagSet("assignee", flag.ExitOnError)

	switch os.Args[3] {
	case "todo":
		title := addTodoCmd.String("title", "", "title of todo")
		assigneId := addTodoCmd.Int("assgine_ID", -1, "please specify the assignee ID you can run show assignee to see all of them ")

		if *title == "" || *assigneId == -1 {
			fmt.Printf("you need to add more arguments\n")
			flag.Usage()
			return
		}
		err := cliservice.AddTodo(*title, *assigneId)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	case "assignee":
		name := addAssigneeCmd.String("name", "", "name of assignee")
		email := addAssigneeCmd.String("email", "", "email of assignee")
		if *name == "" || *email == "" {
			fmt.Printf("you need to add more arguments\n")
			flag.Usage()
			return
		}
		err := cliservice.AddAssignee(*name, *email)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	default:
		fmt.Printf("you need to add more arguments\n")
		flag.Usage()
		return

	}

}
