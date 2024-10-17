package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title      string    `json:"title"`
	AssigneeID uint8     `json:"assigneeId"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TodoList []Todo

var todos TodoList

// ***********************   Getter  **************

func (todos *TodoList) ShowAll() {
	for index, value := range *todos {
		fmt.Printf("id is = %d , todo title is %s and assignee Id is %d , status is %v and created at %s and updated at %s\n\n",
			index,
			value.Title,
			value.AssigneeID,
			value.Status,
			value.CreatedAt.Format(time.RFC3339),
			value.UpdatedAt.Format(time.RFC1123),
		)
	}

}

// TODO show done

// TODO show undones

//TODO show assingnee's one

// ***********************   Setter  **************

func (todos *TodoList) Add(title string, assigneeID uint8) Todo {

//TODO check the assigneeDB is not empty 


	newTodo := Todo{
		Title:      title,
		AssigneeID: assigneeID,
		Status:     false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	*todos = append(*todos, newTodo)
	err := SaveTasks(*todos)
	if err != nil {
		fmt.Errorf("todos seeding has error : %v", err)
	}

	return newTodo
}

// TODO edit Todo needs to check assigneeDB first to ensure it's not empty 
func (todos *TodoList) EditTodo(index uint8, title string, assigneeID uint8, status bool) Todo {
	validation, err := todos.validateTodo(index, title, assigneeID)

	if !validation {
		panic(err)
	}
	(*todos)[index].Title = title
	(*todos)[index].AssigneeID = assigneeID
	(*todos)[index].UpdatedAt = time.Now()

	// 	//TODO save to file should call in here
	err = SaveTasks(*todos)
	if err != nil {
		fmt.Errorf("todos seeding has error : %v", err)
	}

	return (*todos)[index]
}

// // TODO delete Todo
func (todos *TodoList) DeleteTodo(index int) {
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	err := SaveTasks(*todos)
	if err != nil {
		fmt.Errorf("todos seeding has error : %v", err)
	}
	todos.ShowAll()

}

func (todos *TodoList) validateTodo(index uint8, title string, assigneeID uint8) (bool, error) {
	assigneeList := assignees.getAll()
	var err error

	if len(title) < 5 {
		err = errors.New("title length is too short")

	}
	if assigneeList[index].Name == "" || assigneeList[index].Email == "" {
		err = errors.New("assignee is not correct")
	}

	if len(title) < 2 {
		err = errors.New("name should be at least 2 character")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (todos *TodoList) Seedtodos() {

	todos.Add("task1", 0)
	todos.Add("task2 for person2 ", 1)
	err := SaveTasks(*todos)
	if err != nil {
		fmt.Errorf("todos seeding has error : %v", err)
	}
}

func (todos *TodoList) LoadFromFile() error {

	loaded_from_file, err := LoadTasks()
	if err != nil {
		return err
	}
	*todos = loaded_from_file
	return nil
}
