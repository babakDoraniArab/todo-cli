package main

import "time"


type todo struct {
	 title string
	 assignee int
	 status  bool
	 created_at time.Time
	 updated_at time.Time

}

var todos todo

// TODO add todo 
func (t *todo) add (title string , assignee int ){

}

// TODO delete todo 


// TODO edit todo 


// TODO toggle todo 


// TODO validate todo index 
