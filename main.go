package main

func main() {
	assignees.seedAssignees()
	assignees.editAssignee(0, "ali", "ali.test@gmail.com")
	assignees.showAll()
	todos.add("task1", 0)
	todos.add("task2for person2 ", 1)
	todos.showAll()

}
