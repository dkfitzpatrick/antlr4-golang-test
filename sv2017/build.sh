#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package sv2017 -listener -Dlanguage=Go SV2017Lexer.g4 SV2017Parser.g4
