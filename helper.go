package main

import (
	"file_transfer_protocol/file_components"
	"fmt"
	"os"
)

// This functions handles logic behind argu or flag of cli commands.
func argHandler(filepath *string, useHelp *bool, useQuit *bool) {

	// if flag.NArg() == 0 || flag.Arg(0) != "fsa" {
	// 	fmt.Println("correct command required")
	// 	os.Exit(1)
	// }

	helpFlag(useHelp)
	quitFlag(useQuit)
	file_components.ReadFileFromPath(filepath)
	// if *filename == " " {
	// 	fmt.Println("Provide a file path")
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("This is the file path:", *filename)
	// }
}

// logic when user enter -help or --help
func helpFlag(useHelp *bool) {
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
}

// this function when called exits the cli
func quitFlag(useQuit *bool) {
	if *useQuit {

		// logic if user enter "-q or --q arguments"
		os.Exit(0)
	}
}
