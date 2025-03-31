package main

import (
	"fmt"
	"os"
)

func main() {


	const (
		LEFT_PAREN rune = '('
		RIGHT_PAREN rune = ')'
		LEFT_BRACE rune = '{'
		RIGHT_BRACE rune = '}'
		COMMA rune = ','
		DOT rune = '.'
		MINUS rune = '-'
		PLUS rune = '+'
		SEMICOLON rune = ';'
		STAR rune = '*'
	)

	const (	
		NUMBER rune = '#'
		DOLLAR rune = '$'
		ATSIGN rune = '@'
		CARET rune = '^'
	)

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	
	
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	if len(fileContents) < 0 {
		panic("File empty")
	} else {
		for _, ch := range fileContents{
			switch rune(ch) {
			case RIGHT_PAREN:
				fmt.Println("RIGHT_PAREN ) null")

			case LEFT_PAREN:
				fmt.Println("LEFT_PAREN ( null")

			case RIGHT_BRACE:
				fmt.Println("RIGHT_BRACE } null")
			
			case LEFT_BRACE:
				fmt.Println("LEFT_BRACE { null")
			
			case COMMA:
				fmt.Println("COMMA , null")

			case DOT:
				fmt.Println("DOT . null")

			case MINUS:
				fmt.Println("MINUS - null")

			case PLUS:
				fmt.Println("PLUS + null")
			
			case SEMICOLON:
				fmt.Println("SEMICOLON ; null")

			case STAR:
				fmt.Println("STAR * null")

			case NUMBER, DOLLAR, ATSIGN, CARET:
				
				fmt.Fprintf(os.Stderr, "[line %1] Error: Unexpected character: %c\n", ch)
				os.Exit(65)
			}
		}
		fmt.Println("EOF  null")
	}
}
