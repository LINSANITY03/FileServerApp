package file_components

import (
	"fmt"
	"os"
)

// read the file from the given filepath
func ReadFileFromPath(filepath *string) {
	if file, err := os.Open(*filepath); err == nil {
		fmt.Printf("file %v exist on path %v.", file, *filepath)
	} else {
		fmt.Println("Path does not contains any file.")
		os.Exit(1)
	}
}
