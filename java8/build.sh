#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package java8 -listener -Dlanguage=Go Java8Lexer.g4 Java8Parser.g4
