#!/bin/sh

java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Xexact-output-dir -package c -listener -Dlanguage=Go C.g4
