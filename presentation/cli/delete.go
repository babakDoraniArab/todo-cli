package cli

import (
	"flag"
	"fmt"
)

func DeleteRun(flagset *flag.FlagSet, cliservice *CliService) {
	fmt.Printf("DeleteRun Section ,%v\n", flagset)
}
