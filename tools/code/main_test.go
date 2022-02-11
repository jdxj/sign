package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const (
	filename = "code.go"
	src      = `
package main

type code int

const (
	 // Err1 500 Reserve 
	Err1 code = iota + 100000
	// Err2 400 abc
	Err2
	_
	Err5
)

const Err3 int = 56
const Err4 int = 89

`
)

func TestBuild(t *testing.T) {
	fileSet := token.NewFileSet()

	astFile, err := parser.ParseFile(fileSet, filename, src, parser.ParseComments)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	ast.Print(fileSet, astFile)
}

func TestParseCodes(t *testing.T) {
	codes := ParseCodes()
	for _, v := range codes {
		fmt.Printf("%+v\n", v)
	}
}

func TestRenderCodes(t *testing.T) {
	codes := []Code{
		{
			Name: "abc",
			HTTP: 123,
			Desc: "cba",
		},
		{
			Name: "def",
			HTTP: 456,
			Desc: "fed",
		},
		{
			Name: "ghi",
			HTTP: 789,
			Desc: "ihg",
		},
	}
	_ = codes
	data := RenderCodes(nil)
	WriteToFile(data)
}
