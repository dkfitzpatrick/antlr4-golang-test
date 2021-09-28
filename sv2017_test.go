
package main

import (
	"fmt"
	"strconv"
	"testing"
	
	// "github.com/pkg/profile"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dkfitzpatrick/antlr4-golang-test/sv2017"
)

type SourceListener struct {
	*sv2017.BaseSV2017ParserListener
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
	lexer := sv2017.NewSV2017Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	prsr := sv2017.NewSV2017Parser(stream)

	prsr.BuildParseTrees = true
	tree := prsr.Source_text()
	listener := &SourceListener{
		FileName: file,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func BenchmarkSV2017Merlin(b *testing.B) {
	benchmarkSource(b, "sv2017/data/altera_merlin_burst_adapter_new.sv")
}

func BenchmarkSV2017Riscv_6stage(b *testing.B) {
	benchmarkSource(b, "sv2017/data/riscv_6stage.sv")
}

func BenchmarkSV2017Score4(b *testing.B) {
	benchmarkSource(b, "sv2017/data/score4_tb.sv")
}

func BenchmarkSV2017Sprites(b *testing.B) {
	benchmarkSource(b, "sv2017/data/sprites_memory.sv")
}



