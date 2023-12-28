package file_components

import (
	"flag"
	"fmt"
	"os"
)

// read the file from the given filepath
func ReadFileFromPath(filepath *string) string {
	file, err := os.Open(*filepath)
	defer file.Close()
	if err != nil {
		fmt.Println("Path does not contains any file.")
		flag.Usage()
		os.Exit(1)
	} else {
		fmt.Printf("file %v exist on path %v.", file, *filepath)
	}
	return file.Name()
}
