package main

import (
	"bytes"
	"testing"
)

const testDirTreeWithFiles = `├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│		└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
└───zzfile.txt (empty)
`

func TestDirTreeWithFiles(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", true)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}

	result := out.String()
	if result != testDirTreeWithFiles {
		t.Errorf("got:\n%s\nexpected:\n%s\n", result, testDirTreeWithFiles)
	}
}

const testDirTreeWithoutFiles = `├───project
├───static
│	├───a_lorem
│	│	└───ipsum
│	├───css
│	├───html
│	├───js
│	└───z_lorem
│		└───ipsum
└───zline
	└───lorem
		└───ipsum
`

func TestDirTreeWithoutFiles(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", false)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testDirTreeWithoutFiles {
		t.Errorf("got:\n%s\nexpected:\n%s\n", result, testDirTreeWithoutFiles)
	}
}

func TestDirTreeWithUncorrectedDir(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "uncorrected directory", false)
	if err == nil {
		t.Errorf("test for OK Failed - error")
	}
}
