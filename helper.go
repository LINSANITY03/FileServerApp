package main

import (
	"file_transfer_protocol/file_components"
	"os"
)

// This functions handles logic behind argu or flag of cli commands.
func argHandler(filepath *string, useQuit *bool) {

	// if flag.NArg() == 0 || flag.Arg(0) != "fsa" {
	// 	fmt.Println("correct command required")
	// 	os.Exit(1)
	// }

	quitFlag(useQuit)
	file_components.ReadFileFromPath(filepath)
	// if *filename == " " {
	// 	fmt.Println("Provide a file path")
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("This is the file path:", *filename)
	// }
}

// this function when called exits the cli
func quitFlag(useQuit *bool) {
	if *useQuit {

		// logic if user enter "-q or --q arguments"
		os.Exit(0)
	}
}
