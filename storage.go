package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// TODO create a folder in ~/.todo address and create two empty json file called assignee.json and db.json
var (
	homeDir          string
	AppDir           string = ".todoApp"
	TodoAppDir       string
	AssigneeFileAddr string
	TodoFileAddr     string
	err              error
)

func Init() {

	homeDir, err = os.UserHomeDir()

	TodoAppDir = filepath.Join(homeDir, AppDir)
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}

	_, err = os.Stat(TodoAppDir)
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir(TodoAppDir, 0755)
		if err != nil {
			panic("error creating new folder for .TodoAppDir")
		}
	}
	AssigneeFileAddr = filepath.Join(TodoAppDir, "assigneeDB.json")
	TodoFileAddr = filepath.Join(TodoAppDir, "todoDB.json")
	createTodoFile(TodoFileAddr)
	createAssigneeFile(AssigneeFileAddr)
}

func createTodoFile(filepath string) {
	//let's check the folder first
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		// so file is not here and we need to create it
		todoEmptyList := []Todo{}
		// we need to use MarshalIndent because this is an empty json file so for prettified JSON format with indentation if I use Marshal it will create the compact mode
		file, _ := json.MarshalIndent(todoEmptyList, "", " ")

		// let's create the file
		err := os.WriteFile(filepath, file, 0755)
		if err != nil {
			fmt.Printf("error creating file in %s address. %v\n", filepath, err)
		} else {
			fmt.Printf("file is created in this address %s\n", filepath)
		}

	}
}

func createAssigneeFile(filepath string) {
	// let's check the folder first
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		// so file is not here and we need to create it
		assigneeEmptyList := []Assignee{}
		// we need to use MarshalIndent because this is an empty json file so for prettified JSON format with indentation if I use Marshal it will create the compact mode
		file, _ := json.MarshalIndent(assigneeEmptyList, "", " ")

		// let's create the file
		err := os.WriteFile(filepath, file, 0755)
		if err != nil {
			fmt.Printf("error creating file in %s address. %v\n", filepath, err)
		} else {
			fmt.Printf("file is created in this address %s\n", filepath)
		}

	}
}

//TODO load db.json

func LoadTask() ([]Todo, error) {
	fileContent, err := os.ReadFile(TodoFileAddr)
	if err != nil {

		return nil, fmt.Errorf("file could not be opened %v", err)
	}

	var tasks []Todo
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		return nil, fmt.Errorf("unmarshal has problem: %v", err)
	}
	return tasks, nil

}

// TODO save db.json
func SaveTasks(tasks []Todo) error {

	fileContent, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return fmt.Errorf("save task marshal has problem : %v", err)
	}
	err = os.WriteFile(TodoFileAddr, fileContent, 0755)
	if err != nil {
		return fmt.Errorf("writing to file has problem: %v", err)
	}
	fmt.Println("tasks saved successfully")
	return nil

}
