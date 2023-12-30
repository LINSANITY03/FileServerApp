package main

import (
	"file_transfer_protocol/file_components"
	"flag"
	"os"
)

// This functions handles logic behind args or flag of cli commands.
func argHandler(filepath *string, useQuit *bool, useViewFile *string, usePreview *bool) {

	// if flag.NArg() == 0 || flag.Arg(0) != "fsa" {
	// 	fmt.Println("correct command required")
	// 	os.Exit(1)
	// }
	if *useQuit {

		quitFlag(useQuit)
	}

	if *filepath != "" {

		file_components.ReadFileFromPath(filepath)
	}

	if *useViewFile == "" {
		flag.Usage()
		os.Exit(1)

	} else {
		file_components.MarkDownRun(*useViewFile, *usePreview)
	}

}

// this function when called exits the cli
func quitFlag(useQuit *bool) {
	if *useQuit {

		// logic if user enter "-q or --q arguments"
		os.Exit(0)
	}
}
