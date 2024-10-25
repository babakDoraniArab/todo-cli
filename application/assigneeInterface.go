package application

import "github.com/babakDoraniArab/todo-cli/domain"

type AssigneeStorage interface {
	Save(assignee domain.AssigneeList) error
	Load()(domain.AssigneeList,error)
}
