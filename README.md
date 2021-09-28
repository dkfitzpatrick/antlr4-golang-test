

# Grammars for performance testing

A variety of grammars setup for testing performance of golang target for Antlr4.   Base
structure is:

```text
   grammar_test.go
   grammar/*.g4, build.sh, data/
```

The test files only include Benchmark targets.   Run via:

```text
  % go test -run=None -bench=. <grammar>_test.go
```


As of 9/28/2021 - the following are useable:

```text
   c_test.go
   verilog_test.go
   sv2017_test.go
```

The following show fails wrt to the golang generated code (primarily via function overloading
which is not supported in golang):

```text
   java8_test.go
   cpp14_test.go
```


## Details on Antlr Golang Target Fails

## Java8 Grammar (./java8)

java8/java8_parser.go:20936:6: NewEmptyStatementContext redeclared in this block
	previous declaration at parser/java8_parser.go:20275:34

## CPP14 Grammar (./cpp14)

cpp14/cpp14_parser.go:8788:6: NewExpressionContext redeclared in this block
	previous declaration at parser/cpp14_parser.go:5218:6
cpp14/cpp14_parser.go:12221:6: NewEmptyDeclarationContext redeclared in this block
	previous declaration at parser/cpp14_parser.go:11278:36
cpp14/cpp14_parser.go:18851:6: NewDeclaratorContext redeclared in this block
	previous declaration at parser/cpp14_parser.go:5651:6
cpp14/cpp14_parser.go:22313:6: NewInitializerContext redeclared in this block
	previous declaration at parser/cpp14_parser.go:6005:6
cpp14/cpp14parser_base_listener.go:150:58: *NewExpressionContext is not a type
cpp14/cpp14parser_base_listener.go:153:57: *NewExpressionContext is not a type
cpp14/cpp14parser_base_listener.go:168:58: *NewDeclaratorContext is not a type
cpp14/cpp14parser_base_listener.go:171:57: *NewDeclaratorContext is not a type
cpp14/cpp14parser_base_listener.go:180:59: *NewInitializerContext is not a type
cpp14/cpp14parser_base_listener.go:183:58: *NewInitializerContext is not a type
cpp14/cpp14parser_base_listener.go:183:58: too many errors


