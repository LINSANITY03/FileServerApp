package main

import (
	"flag"
)

// Main entrypoint functions of the file.
func main() {
	filepath, useHelp, useQuit := inputParams()

	flag.Parse()

	argHandler(filepath, useHelp, useQuit)

}

// responsible for setting variable for flag
func inputParams() (*string, *bool, *bool) {

	filepath := flag.String("file", " ", "Path of a file")
	useHelp := flag.Bool("help", false, "display available syntax")
	useQuit := flag.Bool("q", false, "Exit the program")
	// whenever we use flag it is always a pointer ("*")
	// *useHelp returns false
	// useHelp returns address 0xc00000a0c6
	return filepath, useHelp, useQuit
}
