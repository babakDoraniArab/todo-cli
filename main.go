package main

func main() {
	assignees.seedAssignees()
	assignees.editAssignee(0, "ali", "ali.test@gmail.com")
	assignees.showAll()

}
