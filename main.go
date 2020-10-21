package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError bool = false

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage gox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
	if hadError {
		os.Exit(65)
	}
}

func error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where, msg string) {
	fmt.Println("[line", line, "] Error:"+where+": "+msg)
	hadError = true
}

func runPrompt() {
	for {
		fmt.Print("> ")
		scanner := bufio.NewReader(os.Stdin)
		line, _, _ := scanner.ReadLine()
		run(string(line))
		hadError = false
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(64)
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
}

func run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}
}
