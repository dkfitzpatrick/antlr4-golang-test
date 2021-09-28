

# Performance Tests


# Antlr Golang Target Fails

Golang target generation issues seem to be focused on relying on overloading functions.

## Java8 Grammar (fails/java8)

java8/java8_parser.go:20936:6: NewEmptyStatementContext redeclared in this block
	previous declaration at parser/java8_parser.go:20275:34

## CPP14 Grammar (fails/cpp14)

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


the following utilities are used:
   * github.com/rakyll/pprof-merge

   % pprof-merge <file1> <file2> ....
   % ls
    merged.data

	








# antlr4-golang-test
For testing antlr4 runtime performance

------------
Java8 client won't compile w/Go

[dan@sfp antlr4-golang-test]$ go build main.go
# github.com/dkfitzpatrick/antlr4-golang-test/java8
java8/java8_parser.go:20936:6: NewEmptyStatementContext redeclared in this block
	previous declaration at java8/java8_parser.go:20275:34
[dan@sfp antlr4-golang-test]$ grep NewEmptyStatementContext java8/*
java8/java8_parser.go:func NewEmptyStatementContext() *StatementContext {
java8/java8_parser.go:func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
java8/java8_parser.go:	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())



