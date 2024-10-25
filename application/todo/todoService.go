package todo

import (
	"fmt"
	"github.com/babakDoraniArab/todo-cli/application/assignee"
	"time"

	"github.com/babakDoraniArab/todo-cli/domain"
	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)

type TodoService struct {
	Storage         TodoStorage
	AssigneeStorage assignee.AssigneeStorage
}

//************************ constructor************

func NewTodoService(storage TodoStorage, assigneeStorage assignee.AssigneeStorage) *TodoService {
	return &TodoService{Storage: storage, AssigneeStorage: assigneeStorage}
}

// ***********************   Getter  **************

func (t *TodoService) ShowAll() {
	todos, err := t.Storage.Load()
	helper.CheckErr(err, "could not load the todos from file")
	for index, value := range todos {
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

// ***********************   Setter  **************

func (t *TodoService) Add(title string, assigneeID int) domain.Todo {

	// helper.CheckErr(AssigneeDbCheck())
	todos, err := t.Storage.Load()
	helper.CheckErr(err, "could not load the todos from file")

	helper.CheckErr(t.ValidateTodo(title, assigneeID))
	newTodo := domain.Todo{
		Title:      title,
		AssigneeID: assigneeID,
		Status:     false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	todos = append(todos, newTodo)
	err = t.Storage.Save(todos)
	helper.CheckErr(err, "todos saving has error when I try to add")

	return newTodo
}

func (t *TodoService) EditTodo(index int, title string, assigneeID int, status bool) error {

	todos, err := t.Storage.Load()
	helper.CheckErr(err, "could not load the todos from file")

	// helper.CheckErr(AssigneeDbCheck())
	helper.CheckErr(t.ValidateTodo(title, assigneeID, index))

	(todos)[index].Title = title
	(todos)[index].AssigneeID = assigneeID
	(todos)[index].UpdatedAt = time.Now()

	helper.CheckErr(t.Storage.Save(todos), "todos saving has error when I try to edit")

	return nil
}

func (t *TodoService) DeleteTodo(index int) error {
	todos, err := t.Storage.Load()
	helper.CheckErr(err, "could not load the todos from file")

	todos = append((todos)[:index], (todos)[index+1:]...)
	helper.CheckErr(t.Storage.Save(todos), "todos saving has error")
	return nil

}

func (t *TodoService) Seedtodos() {

	t.Add("task1", 0)
	t.Add("task2 for person2 ", 1)

}

//************************ validator ***************

func (t *TodoService) ValidateTodo(title string, assigneeID int, index ...int) error {
	assigneeList, err := t.AssigneeStorage.Load()
	helper.CheckErr(err, "could not load the todos from file")

	if len(title) < 5 {

		err = helper.NewError("title length is too short")

	}

	if assigneeID < 0 || assigneeList[assigneeID].Name == "" || assigneeList[assigneeID].Email == "" {

		err = helper.NewError("assignee ID is not correct")
	}

	if len(title) < 2 {

		err = helper.NewError("name should be at least 2 character")
	}

	if len(index) > 0 && index[0] < 0 {
		err = helper.NewError("the todo list id you have selected is not correct ")
	}

	if err != nil {
		return err
	}
	return nil
}
