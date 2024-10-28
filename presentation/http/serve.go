package http

import (
	"flag"
	"fmt"
)

//TODO I will add Http configuration in here

func Run(flagset *flag.FlagSet) {
	fmt.Printf("I will run echo in the near future,args are: %v", flagset)
}
