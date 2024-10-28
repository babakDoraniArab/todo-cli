package assignee

import (
	"fmt"
	"time"

	"github.com/babakDoraniArab/todo-cli/domain"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)

type AssigneeService struct {
	Storage AssigneeStorage
}

// ***********************   Getter  **************

func (s *AssigneeService) ShowAllAssignee() error {
	assignees, err := s.Storage.Load()
	helper.CheckErr(err, "could not load assignee")
	for index, value := range assignees {
		fmt.Printf("id is = %d , assigne name is %s and his/her email is %s and created at %s and updated at %s\n\n",
			index,
			value.Name,
			value.Email,
			value.CreatedAt.Format(time.RFC3339),
			value.UpdatedAt.Format(time.RFC1123),
		)
	}
	return nil

}
func (s *AssigneeService) GetAllAssginee() (domain.AssigneeList, error) {
	assignees, err := s.Storage.Load()
	helper.CheckErr(err, "could not load assignee")

	return assignees, nil
}

// ***********************   Setter  **************

func (s *AssigneeService) AddAssignee(name string, email string) error {

	_, err := domain.ValidateAssignee(name, email)
	helper.CheckErr(err, "validation error: ")

	assignees, err := s.Storage.Load()
	helper.CheckErr(err, "could not load assignee")

	newAssignee := domain.Assignee{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assignees = append(assignees, newAssignee)
	err = s.Storage.Save(assignees)
	helper.CheckErr(err, "could not save assignee")

	return nil
}

func (s *AssigneeService) EditAssignee(index int, name string, email string) error {

	assignees, err := s.Storage.Load()
	helper.CheckErr(err, "could not load assignee")

	_, err = domain.ValidateAssigneesIndex(index, assignees)
	helper.CheckErr(err, "validation  Index error: ")

	(assignees)[index].Name = name
	(assignees)[index].Email = email

	err = s.Storage.Save(assignees)
	helper.CheckErr(err, "could not save assignee")

	return nil
}

func (s *AssigneeService) DeleteAssignee(index int) error {

	assignees, err := s.Storage.Load()
	helper.CheckErr(err, "could not load assignee")

	assignees = append((assignees)[:index], (assignees)[index+1:]...)

	err = s.Storage.Save(assignees)
	helper.CheckErr(err, "could not save assignee")
	return nil

}

// ********************** helper ***********************

func (s *AssigneeService) SeedAssignees() {

	s.AddAssignee("babak dorani", "babak.dorani@gmail.com")
	s.AddAssignee("babak do2000", "babak.do2000@gmail.com")
	s.AddAssignee("babak outlook", "babak.dorani@outlook.com")

}
