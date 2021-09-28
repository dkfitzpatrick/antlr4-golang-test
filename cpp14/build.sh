#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package cpp14 -listener -Dlanguage=Go CPP14Lexer.g4 CPP14Parser.g4
