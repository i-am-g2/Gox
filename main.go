package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/i-am-g2/Gox/gox"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage gox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
	if gox.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	for {
		fmt.Print("> ")
		scanner := bufio.NewReader(os.Stdin)
		line, _, _ := scanner.ReadLine()
		run(string(line))
		gox.HadError = false
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(64)
	}
	run(string(bytes))
	if gox.HadError {
		os.Exit(65)
	}
}

func run(source string) {
	scanner := gox.NewScanner(source)
	tokens := scanner.ScanTokens()

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}
}
