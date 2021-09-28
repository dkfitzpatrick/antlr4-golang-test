#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package earth -listener -Dlanguage=Go EarthLexer.g4 EarthParser.g4
