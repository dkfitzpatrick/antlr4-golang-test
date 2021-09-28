#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package verilog -listener -Dlanguage=Go VerilogLexer.g4 VerilogParser.g4
