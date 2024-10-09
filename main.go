package main

import (
	"fmt"
	"time"
)

func main() {
	assignees.add("b", "babak.dorani@gmail.com")
	assignees.add("babak2", "babak4.dorani@gmail.com")
	assignees.add("babak3", "babak4.dorani@gmail.com")
	for _, value := range assignees {
		fmt.Printf("assigne name is %s and his/her email is %s and created at %s and updated at %s\n\n",
			value.name,
			value.email,
			value.created_at.Format(time.RFC3339),
			value.updated_at.Format(time.RFC1123),
		)
	}
}
