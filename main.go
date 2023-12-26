package main

import (
	"flag"
)

// Main entrypoint functions of the file.
func main() {
	filepath, useQuit, useViewFile := inputParams()

	flag.Parse()

	argHandler(filepath, useQuit, useViewFile)

}

// responsible for setting variable for flag
func inputParams() (*string, *bool, *string) {

	filepath := flag.String("file", "", "Path of a file")
	useViewFile := flag.String("vf", "", "View the content of a file")
	useQuit := flag.Bool("q", false, "Exit the program")
	// whenever we use flag it is always a pointer ("*")
	// *useHelp returns false
	// useHelp returns address 0xc00000a0c6
	return filepath, useQuit, useViewFile
}
