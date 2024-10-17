package main

import (

	"fmt"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
	"time"
)

type Todo struct {
	Title      string    `json:"title"`
	AssigneeID int     `json:"assigneeId"`
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

func (todos *TodoList) Add(title string, assigneeID int) Todo {

	
	helper.CheckErr(AssigneeDbCheck())
	helper.CheckErr(todos.validateTodo(title, assigneeID))
	newTodo := Todo{
		Title:      title,
		AssigneeID: assigneeID,
		Status:     false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	*todos = append(*todos, newTodo)
	helper.CheckErr(SaveTasks(*todos),"todos saving has error when I try to add")

	return newTodo
}


func (todos *TodoList) EditTodo(index int, title string, assigneeID int, status bool) error {

	helper.CheckErr(AssigneeDbCheck())
	helper.CheckErr(todos.validateTodo( title, assigneeID, index))
	
	
	(*todos)[index].Title = title
	(*todos)[index].AssigneeID = assigneeID
	(*todos)[index].UpdatedAt = time.Now()

	

	helper.CheckErr(SaveTasks(*todos),"todos saving has error when I try to edit")

	return nil
}


func (todos *TodoList) DeleteTodo(index int) {
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	helper.CheckErr(SaveTasks(*todos),"todos saving has error")
	todos.ShowAll()

}

func (todos *TodoList) validateTodo( title string, assigneeID int, index ...int) error {
	assigneeList := assignees.getAll()
	var err error

	if len(title) < 5 {
	
		err = helper.NewError("title length is too short")

	}
	if assigneeID <0 || assigneeList[assigneeID].Name == "" || assigneeList[assigneeID].Email == "" {
		
		err = helper.NewError("assignee ID is not correct")
	}

	if len(title) < 2 {
		
		err = helper.NewError("name should be at least 2 character")
	}

	if len(index)>0 &&  index[0] <0 {
		err = helper.NewError("the todo list id you have selected is not correct ")
	}



	if err != nil {
		return err
	}
	return nil
}

func (todos *TodoList) Seedtodos() {

	todos.Add("task1", 0)
	todos.Add("task2 for person2 ", 1)
	helper.CheckErr(SaveTasks(*todos),"todos seeding has error")

}

func (todos *TodoList) LoadFromFile() error {

	loaded_from_file, err := LoadTasks()
	helper.CheckErr(err,"could not load the files from database")
	*todos = loaded_from_file
	return nil
}
