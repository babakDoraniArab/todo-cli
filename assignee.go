package main

import (
	"errors"
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

// TODO create assignee

var assignees AssigneeList

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

// TODO delete assignee

// TODO validate assignee index
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
