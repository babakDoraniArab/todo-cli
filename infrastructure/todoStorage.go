package infrastructure

import (
	"encoding/json"

	"os"
	

	"github.com/babakDoraniArab/todo-cli/domain"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)

const (
	TodoFileName = "todoDB.json"
)

type TodoStorage struct {
	filepath string
}

func  InitTodoStorage() (*TodoStorage,error) {
	storageAddr, err :=  Initialization(TodoFileName, []domain.Todo{})
	helper.CheckErr(err,"could not initialize Todo storage")
	
	return &TodoStorage{filepath: storageAddr} ,nil
}

func (todoStorage *TodoStorage) Save(todoList domain.TodoList) error {

	file, _ := json.MarshalIndent(todoList, "", " ")

	err := os.WriteFile(todoStorage.filepath, file, 0755)
	helper.CheckErr(err, "could not save the todos")

	return nil

}

func (todoStorage *TodoStorage) Load() (domain.TodoList,error) {
	fileContent, _ := os.ReadFile(todoStorage.filepath)
	var todo domain.TodoList

	err := json.Unmarshal(fileContent, &todo)
	helper.CheckErr(err, "could not load the todo storage")

	return todo, nil

}
