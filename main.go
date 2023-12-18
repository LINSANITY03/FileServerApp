package main

import (
	"flag"
	"fmt"
	"os"
)

// Main entrypoint functions of the file.
func main() {
	argHandler()

}

// This functions handles logic behind argu or flag of cli commands.
func argHandler() {

	useHelp := flag.Bool("help", false, "display available syntax")
	useQuit := flag.Bool("q", false, "Exit the program")
	flag.Parse()
	// whenever we use flag it is always a pointer ("*")
	// *useHelp returns false
	// useHelp returns address 0xc00000a0c6

	if *useHelp {

		fmt.Println(
			`
			This is the default help message.
			Syntax:
			-help:
				Shows the available syntax for this cli.
			-q:
				Quits the cli.
			`)
	}

	if *useQuit {
		os.Exit(0)
	}

}
