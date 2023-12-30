package main

import (
	"bytes"
	"file_transfer_protocol/file_components"
	"os"
	"testing"
)

const (
	inputFile    = "../testdata/example.txt"
	markdownFile = "../testdata/example.txt.html"
	// asd = "C:\Users\dell\AppData\Local\Temp\mdp923273673.html"
	asd = "C:/Users/dell/AppData/Local/Temp/mdp1137100221.html"
)

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

// func TestMarkDownRun(t *testing.T) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cmd := exec.Command("go", "run", ".", "-vf", filepath.Join(dir, inputFile), "-s")
// 	cmd.Dir = filepath.Join(dir, "../")

// 	out, err := cmd.Output()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	resultFilePath := filepath.ToSlash(string(out))
// 	defer os.Remove(resultFilePath)

// 	result, err := os.ReadFile(resultFilePath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected, err := os.ReadFile(filepath.Join(dir, markdownFile))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !bytes.Equal(expected, result) {
// 		t.Logf("markdown:\n%s\n", expected)
// 		t.Logf("result:\n%s\n", result)
// 		t.Error("Result content does not match markdown file")
// 	}
// }
