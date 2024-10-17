package helper

import (
	"log"
	"os"
)



func CheckErr(err error, explanation ...string){
	if err!= nil {
		if len(explanation)> 0 {
		log.Printf(Red+"%s: %v\n"+Reset,explanation[0],err)
	}else{
		log.Printf(Red+"Error: %s"+Reset , err)
	}
		os.Exit(0)
		
	}
}