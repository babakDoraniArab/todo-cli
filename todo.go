package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	title      string
	assigneeID uint8
	status     bool
	created_at time.Time
	updated_at time.Time
}

type TodoList []Todo

var todos TodoList

// ***********************   Getter  **************

func (todos *TodoList) showAll() {
	for index, value := range *todos {
		fmt.Printf("id is = %d , todo title is %s and assignee Id is %d , status is %v and created at %s and updated at %s\n\n",
			index,
			value.title,
			value.assigneeID,
			value.status,
			value.created_at.Format(time.RFC3339),
			value.updated_at.Format(time.RFC1123),
		)
	}

}

// TODO show done

// TODO show undones

//TODO show assingnee's one

// ***********************   Setter  **************

func (todos *TodoList) add(title string, assigneeID uint8) Todo {

	newTodo := Todo{
		title:      title,
		assigneeID: assigneeID,
		status:     false,
		created_at: time.Now(),
		updated_at: time.Now(),
	}

	*todos = append(*todos, newTodo)

	return newTodo
}

// TODO edit Todo
func (todos *TodoList) editTodo(index uint8, title string, assigneeID uint8, status bool) Todo {
	validation, err := todos.validateTodo(index, title, assigneeID)

	if !validation {
		panic(err)
	}
	(*todos)[index].title = title
	(*todos)[index].assigneeID = assigneeID
	(*todos)[index].updated_at = time.Now()

	// 	//TODO save to file should call in here

	return (*todos)[index]
}

// // TODO delete Todo
func (todos *TodoList) deleteTodo(index int) {
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	todos.showAll()

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

func (todos *TodoList) seedtodos() {

	todos.add("task1", 0)
	todos.add("task2for person2 ", 1)
}
