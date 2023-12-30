package main

import (
	"flag"
)

// Main entrypoint functions of the file.
func main() {
	filepath, useQuit, useViewFile, usePreview := inputParams()

	flag.Parse()

	argHandler(filepath, useQuit, useViewFile, usePreview)

}

// responsible for setting variable for flag
func inputParams() (*string, *bool, *string, *bool) {

	filepath := flag.String("file", "", "Path of a file")
	useViewFile := flag.String("vf", "", "View the content of a file")
	useQuit := flag.Bool("q", false, "Exit the program")
	usePreview := flag.Bool("s", false, "Display the markdown file")
	// whenever we use flag it is always a pointer ("*")
	// *useHelp returns false
	// useHelp returns address 0xc00000a0c6
	return filepath, useQuit, useViewFile, usePreview
}
