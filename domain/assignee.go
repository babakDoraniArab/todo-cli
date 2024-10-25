package domain

import (
	"errors"
	"strings"
	"time"
)

type Assignee struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type AssigneeList []Assignee

// var assignees AssigneeList

func ValidateAssigneesIndex(index int, assignees AssigneeList) (bool, error) {
	var err error
	
	if index < 0 || index > len(assignees) {
		err = errors.New("the assignee ID is incorrect")

	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func ValidateAssignee(title string, email string) (bool, error) {

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
