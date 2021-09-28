
package main

import (
	"fmt"
	"strconv"
	"testing"
	
	// "github.com/pkg/profile"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dkfitzpatrick/antlr4-golang-test/earth"
)

type SourceListener struct {
	*earth.BaseEarthParserListener
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
	lexer := earth.NewEarthLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	prsr := earth.NewEarthParser(stream)

	prsr.BuildParseTrees = true
	tree := prsr.EarthFile()
	listener := &SourceListener{
		FileName: file,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func BenchmarkEarthFile1(b *testing.B) {
	benchmarkSource(b, "earth/data/earthfile1")
}

func BenchmarkEarthFile2(b *testing.B) {
	benchmarkSource(b, "earth/data/earthfile2")
}

