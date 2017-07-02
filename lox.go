package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"lox/lox"
)

func main() {
	exitCode := 0

	switch len(os.Args) - 1 {
	case 0:
		exitCode = runPrompt()
	case 1:
		exitCode = runFile(os.Args[1])
	default:
		fmt.Fprintln(os.Stderr, "Usage: golox [script]")
		exitCode = -1
	}

	os.Exit(exitCode)
}

func runPrompt() int {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		interpreter = lox.NewInterpreter()
	)
	
	for {
		fmt.Print("> ")
		
		scanner.Scan()
		line := scanner.Text()
		
		if line == "exit" {
			break
		}
		
		tokens := lox.Lex(line)
		
		parser := lox.NewParser(tokens)
		
		stmts, hadError := parser.Parse()
		if hadError {
			continue
		}
		
		err := interpreter.Interpret(stmts)
		if err != nil {
			fmt.Println("Runtime error:", err.Error())
		}
	}

	return 0
}

func runFile(filename string) int {
	data, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Printf("Error reading %q: %v", filename, err)
		return 1
	}
	
	tokens := lox.Lex(string(data))
	parser := lox.NewParser(tokens)
	
	stmts, hadError := parser.Parse()
	if hadError {
		fmt.Println("Parsing errors occured")
		return 2
	}
	
	interpreter := lox.NewInterpreter()
	
	err = interpreter.Interpret(stmts)
	if err != nil {
		fmt.Println("Runtime error:", err.Error())
		return 3
	}
	
	return 0
}
