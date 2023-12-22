package file_components

import (
	"fmt"
	"os"
)

// read the file from the given filepath
func ReadFileFromPath(filepath *string) {
	file, err := os.Open(*filepath)
	defer file.Close()
	if err == nil {
		fmt.Printf("file %v exist on path %v.", file, *filepath)
	} else {
		fmt.Println("Path does not contains any file.")
		os.Exit(1)
	}
}
