package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// TODO create a folder in ~/.todo address and create two empty json file called assignee.json and db.json
func Init() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}
	todoAppDir := filepath.Join(homeDir, ".todoApp")

	_ ,err = os.Stat(todoAppDir)
	if err!= nil && os.IsNotExist(err) {
		err:= os.Mkdir(todoAppDir,0755)
		if err!= nil {
			panic("error creating new folder for .todoAppDir")
		}
	}
	fmt.Printf(todoAppDir)
}
//TODO load db.json

//TODO save db.json
