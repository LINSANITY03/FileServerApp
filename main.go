package main

import (
	"fmt"
	"os"
)

func main() {
	commands := argHandler()

	for i := 0; i < len(commands); i++ {
		if commands[i] == "connect" {
			fmt.Println("connection detected")
		}
	}

}

func argHandler() []string {
	all_args := os.Args
	if len(os.Args) < 2 {
		fmt.Println("Command not recognized")
		os.Exit(0)
	}

	return all_args
}
