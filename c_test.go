
package main

import (
	"fmt"
	"strconv"
	"testing"
	
	// "github.com/pkg/profile"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dkfitzpatrick/antlr4-golang-test/c"
)

type SourceListener struct {
	*c.BaseCListener
	FileName       string
}

type SourceErrorListener struct {
	antlr.ErrorListener
}

func (d *SourceErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, 
	line, column int, msg string, e antlr.RecognitionException) {
	panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}


func benchmarkSource(b *testing.B, file string) {
        // defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
        // defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	fmt.Printf("loading:  %s\n", file)
	input, _ := antlr.NewFileStream(file)
	lexer := c.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	prsr := c.NewCParser(stream)

	prsr.BuildParseTrees = true
	tree := prsr.CompilationUnit()
	listener := &SourceListener{
		FileName: file,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func BenchmarkCTest1(b *testing.B) {
	benchmarkSource(b, "c/data/test1.c")
}

