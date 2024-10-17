package helper

import (
	"log"
	"os"
)



func CheckErr(err error, explanation ...string){
	if err!= nil {
		if len(explanation)> 0 {
		log.Printf("%s: %v\n",explanation[0],err)
	}else{
		log.Println("Error:", err)
	}
		os.Exit(0)
		
	}
}