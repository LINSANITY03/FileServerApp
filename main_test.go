package main

import (
	"bytes"
	"file_transfer_protocol/file_components"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	inputFile    = "./testdata/example.txt"
	markdownFile = "./testdata/example.txt.html"
)

func TestFileFlag(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, "/testdata/example.txt")
	res := file_components.ReadFileFromPath(&cmdPath)

	if res != cmdPath {
		t.Error("Results were different than expected.")
	} else {
		fmt.Println("Expected outcome acheived.")

	}
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := file_components.ParseContent(input)
	expected, err := os.ReadFile(markdownFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, result) {
		t.Logf("markdown:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match markdown file")
	}

}
