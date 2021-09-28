
package main

import (
	"fmt"
	"strconv"
	"testing"
	
	// "github.com/pkg/profile"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dkfitzpatrick/antlr4-golang-test/verilog"
)

type SourceListener struct {
	*verilog.BaseVerilogParserListener
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
	lexer := verilog.NewVerilogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	prsr := verilog.NewVerilogParser(stream)

	prsr.BuildParseTrees = true
	tree := prsr.Source_text()
	listener := &SourceListener{
		FileName: file,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func BenchmarkVerilogClock(b *testing.B) {
	benchmarkSource(b, "verilog/data/CLOCK_SOURCE.v")
}

func BenchmarkVerilogCNN(b *testing.B) {
	benchmarkSource(b, "verilog/data/CNN_top.v")
}

func BenchmarkVerilogCore(b *testing.B) {
	benchmarkSource(b, "verilog/data/mkFBox_Core.v")
}



