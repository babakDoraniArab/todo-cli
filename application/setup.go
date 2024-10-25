package application

import (
	"fmt"
	"github.com/babakDoraniArab/todo-cli/infrastructure"
)


type ApplicationSetup struct {
	TodoService     *TodoService
	AssigneeService *AssigneeService
}

func Setup() (*ApplicationSetup, error) {

	todoStorage, err := infrastructure.InitTodoStorage()
	if err != nil {
		return nil, fmt.Errorf("error initializing TodoStorage: %v", err)
	}


	assigneeStorage, err := infrastructure.InitAssigneeStorage()
	if err != nil {
		return nil, fmt.Errorf("error initializing AssigneeStorage: %v", err)
	}

// I did it with constructor 
	todoService := NewTodoService(todoStorage,assigneeStorage)
	
// I did not create a constructor but initialized an instance manually
// the applicatioSetup is getting pointer because it will reserve a place in memory and always work on this address so if you want to have constructor you can go with todoService approach I coded in the top or you can create an instance manually but make sure you are ppasing the & of that 
	assigneeService := &AssigneeService{
		Storage: assigneeStorage,
	}

	return &ApplicationSetup{
		TodoService:     todoService,
		AssigneeService: assigneeService,
	}, nil
}
