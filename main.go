package main

import (
	"fmt"
	"os"
)

func main() {
	argHandler()

}

func argHandler() []string {
	all_args := os.Args

	if len(os.Args) < 2 || os.Args[1] != "filesev" {
		fmt.Println("Command not recognized")
		os.Exit(0)
	}

	// switch all_args[1] {
	// case condition:

	// }

	return all_args
}
