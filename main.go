package main

func main() {
	assignees.add("bbbb", "babak.dorani@gmail.com")
	assignees.add("babak2", "babak4.dorani@gmail.com")
	assignees.add("babak3", "babak4.dorani@gmail.com")
	assignees.editAssignee(0, "ali", "ali.test@gmail.com")
	assignees.showAll()

}
