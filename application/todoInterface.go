package application

import "github.com/babakDoraniArab/todo-cli/domain"

type TodoStorage interface {
	Save(todo domain.TodoList) error
	Load()(domain.TodoList,error)
}
