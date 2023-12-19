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
	// whenever we use flag it is always a pointer ("*")
	// *useHelp returns false
	// useHelp returns address 0xc00000a0c6

	flag.Parse()

	if flag.NArg() == 0 || flag.Arg(1) != "fsa" {
		fmt.Println("correct command required")
		os.Exit(1)
	}

	if *useHelp {
		// logic if user enter "-help or --help arguments"
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

		// logic if user enter "-q or --q arguments"
		os.Exit(0)
	}

}
