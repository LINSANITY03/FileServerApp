package file_components

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>Markdown Preview Tool</title>
	</head>
	<body>
	`
	footer = `
	</body>
	</html>
	`
)

// This function reads the content of a file and write to a html file.
// the created file name will be in a temp file.
// Prints the location of the created file.
func MarkDownRun(filename string, skipPrev bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(filename)
	if err != nil {
		flag.Usage()
		return err
	}
	htmlData := ParseContent(input)
	// Create temporary file and check for errors
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		flag.Usage()
		return err
	}
	if err := temp.Close(); err != nil {
		flag.Usage()
		return err
	}
	outName := temp.Name()
	fmt.Println(outName)
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	defer os.Remove(outName)
	if skipPrev {
		return nil
	}
	return preview(outName)

}

// Input file content are added to a buffer and returned.
func ParseContent(input []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	//  Create a buffer of bytes to write to file
	var buffer bytes.Buffer
	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)
	return buffer.Bytes()
}

// write the data into a .html file extension
func saveHTML(outFname string, data []byte) error {
	// Write the bytes to the file
	return os.WriteFile(outFname, data, 0644)
}

// run the generated html file in browser
func preview(fname string) error {
	return exec.Command("cmd", "/c", "start", "", fname).Run()
}
