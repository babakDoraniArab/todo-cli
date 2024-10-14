package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Assignee struct {
	name       string
	email      string
	created_at time.Time
	updated_at time.Time
}

type AssigneeList []Assignee

var assignees AssigneeList

func (assignees AssigneeList) showAll() {
	for index, value := range assignees {
		fmt.Printf("id is = %d , assigne name is %s and his/her email is %s and created at %s and updated at %s\n\n",
			index,
			value.name,
			value.email,
			value.created_at.Format(time.RFC3339),
			value.updated_at.Format(time.RFC1123),
		)
	}

}

func (assignees *AssigneeList) add(title string, email string) Assignee {

	validation, err := assignees.validateAssignee(title, email)

	if !validation {
		panic(err)
	}
	newAssignee := Assignee{
		name:       title,
		email:      email,
		created_at: time.Now(),
		updated_at: time.Now(),
	}

	*assignees = append(*assignees, newAssignee)

	return newAssignee
}

// TODO edit assignee
func (assignees *AssigneeList) editAssignee(index int, name string, email string) Assignee {
	validation, err := assignees.validateAssigneesIndex(index)
	if !validation {
		panic(err)
	}
	(*assignees)[index].name = name
	(*assignees)[index].email = email

	//TODO save to file should call in here

	return (*assignees)[index]
}

// TODO delete assignee
func (assignees *AssigneeList) deleteAssignee(index int) {
	*assignees = append((*assignees)[:index], (*assignees)[index+1:]...)
	assignees.showAll()

}
func (assignees *AssigneeList) validateAssigneesIndex(index int) (bool, error) {
	var err error
	if index < 0 || index > len(*assignees) {
		err = errors.New("the assignee ID is incorrect")

	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (assignees *AssigneeList) validateAssignee(title string, email string) (bool, error) {

	var err error
	if len(email) < 6 {
		err = errors.New("email length is too short")

	}
	if !strings.Contains(email, "@") {
		err = errors.New("email is not valid")
	}
	if len(title) < 2 {
		err = errors.New("name should be at least 2 character")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
