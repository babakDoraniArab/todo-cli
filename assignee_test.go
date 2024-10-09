package main

import (
	"slices"
	"testing"
)

// TODO test add func
func TestAdd(t *testing.T) {
	//01 we need to create a list of assingess
	var assignees AssigneeList

	// create a new assignee and test it is created correctly or not
	// TODO create a new assignee
	newAssignee := assignees.add("babak", "babak@gmail.com")
	// TODO check name
	if newAssignee.name != "babak" {
		t.Errorf("name must be babak but it's not correct we got %s", newAssignee.name)
	}
	//TODO check email
	if newAssignee.email != "babak@gmail.com" {
		t.Errorf("email must be babak@gmail.com but it's not correct we got %s", newAssignee.name)
	}

	//TODO check created_at and updated at are empty or not
	if newAssignee.created_at.IsZero() || newAssignee.updated_at.IsZero() {
		t.Error("created_at or updated_at is empty string, please check it ")
	}

	//TODO check assigne list has the new assignee or not
	if !slices.Contains(assignees, newAssignee) {
		t.Error("newAssignee is not added to the assignee slice")
	}

}

// TODO test validate func
