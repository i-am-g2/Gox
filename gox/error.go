package gox

import "fmt"

// HadError boolean denoting whether error was encountered during scanning/parsing,etc or not
var HadError bool = false

func error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where, msg string) {
	fmt.Println("[line", line, "] Error:"+where+": "+msg)
}
