package main

func main() {
	// create database files in the ~/.todoApp/ filepath
	Init()
	// assignees.seedAssignees()
	assignees.seedAssignees()
	err := todos.LoadFromFile()
	if err != nil {
		panic(err)
	}
	
	// todos.Seedtodos()
	// todos.ShowAll()
	 todos.Add("1",1)
	// todos.DeleteTodo(0)
	todos.EditTodo(0, "task2 for person2 edited", 1, true)
	// todos.Seedtodos()
	// // assignees.showAll()

	todos.ShowAll()

	//TODO flags
}
