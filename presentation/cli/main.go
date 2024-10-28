package cli

import (
	"flag"
	"fmt"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
	"github.com/babakDoraniArab/todo-cli/presentation/http"
	"os"
)

func Execute(cliService *CliService) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(helper.Red + "you need to add a subcommand" + helper.Reset)
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		AddRun(addCmd, cliService)

	case "edit":
		editCmd.Parse(os.Args[2:])
		EditRun(editCmd, cliService)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		DeleteRun(deleteCmd, cliService)
	case "serve":
		serveCmd.Parse(os.Args[2:])
		http.Run(serveCmd)

	default:
		flag.Usage()
	}

}
