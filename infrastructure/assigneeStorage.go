package infrastructure

import (
	"encoding/json"
	"os"
	"github.com/babakDoraniArab/todo-cli/domain"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)

const (
	AssigneeFileName = "Assignee.json"
)

type AssigneeStorage struct {
	filepath string
}

func InitAssigneeStorage() (*AssigneeStorage, error) {
	fileAddr, err := Initialization(AssigneeFileName, []domain.Assignee{})
	
	helper.CheckErr(err,"could not initialize Assignee storage")
	return &AssigneeStorage{filepath: fileAddr}, nil
}

func (assigneeStorage *AssigneeStorage) Save(assigneeList domain.AssigneeList) error {

	file, _ := json.MarshalIndent(assigneeList, "", " ")

	err := os.WriteFile(assigneeStorage.filepath, file, 0755)
	helper.CheckErr(err, "could not save the assignee")

	return nil

}

func (assigneeStorage *AssigneeStorage) Load() (domain.AssigneeList,error ){
	fileContent, _ := os.ReadFile(assigneeStorage.filepath)
	var todo domain.AssigneeList

	err := json.Unmarshal(fileContent, &todo)
	helper.CheckErr(err, "could not load the assignee storage")

	return todo,nil

}
