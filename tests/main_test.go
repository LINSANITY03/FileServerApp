package main

import (
	"file_transfer_protocol/file_components"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFileFlag(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, "../testdata/example.txt")
	fmt.Println(cmdPath)
	res := file_components.ReadFileFromPath(&cmdPath)

	if res != cmdPath {
		t.Error("Results were different than expected.")
	} else {
		fmt.Println("Expected outcome acheived.")

	}
}
