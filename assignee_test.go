package main

import (
	"slices"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {

	var assignees AssigneeList

	newAssignee := assignees.add("babak", "babak@gmail.com")

	if newAssignee.name != "babak" {
		t.Errorf("name must be babak but it's not correct we got %s", newAssignee.name)
	}
	if newAssignee.email != "babak@gmail.com" {
		t.Errorf("email must be babak@gmail.com but it's not correct we got %s", newAssignee.name)
	}
	if newAssignee.created_at.IsZero() || newAssignee.updated_at.IsZero() {
		t.Error("created_at or updated_at is empty string, please check it ")
	}
	if !slices.Contains(assignees, newAssignee) {
		t.Error("newAssignee is not added to the assignee slice")
	}

}

func TestValidateAssignee(t *testing.T) {
	email := "babak@gmail.com"
	title := "babak"

	if len(email) < 6 {
		t.Error("email is too short")
	}

	if !strings.Contains(email, "@") {
		t.Error("email is not valid")
	}
	if len(title) < 2 {
		t.Error("title is too short")
	}

}

//TODO testDeleteAssignee

//TODO testValidateAssigneesIndex


//TODO testEditAssignee

// TODO testShowAll