package main

import (
	"fmt"
	"os"
	"unicode"
	"strconv"
)

func main() {
	var has_error = false
	var line = 1

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
		EQUAL rune = '='
		BANG rune = '!'
		LESS rune = '<'
		GREATER rune = '>'
		SLASH rune = '/'
		QUOTE rune = '"'
	)

	const (	
		NUMBER rune = '#'
		DOLLAR rune = '$'
		ATSIGN rune = '@'
		CARET rune = '^'
		PERCENT rune = '%'
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
		indx := 0
		for indx < len(fileContents) {
    		ch := rune(fileContents[indx])
			switch rune(ch) {

			case ' ', '\t':  
				indx++

			case '\n':  
				line++ // Newline increments line count.
				indx++

			case RIGHT_PAREN:
				fmt.Println("RIGHT_PAREN ) null")
				indx++

			case LEFT_PAREN:
				fmt.Println("LEFT_PAREN ( null")
				indx++

			case RIGHT_BRACE:
				fmt.Println("RIGHT_BRACE } null")
				indx++
			
			case LEFT_BRACE:
				fmt.Println("LEFT_BRACE { null")
				indx++
			
			case COMMA:
				fmt.Println("COMMA , null")
				indx++

			case DOT:
				fmt.Println("DOT . null")
				indx++

			case MINUS:
				fmt.Println("MINUS - null")
				indx++

			case PLUS:
				fmt.Println("PLUS + null")
				indx++
			
			case SEMICOLON:
				fmt.Println("SEMICOLON ; null")
				indx++

			case STAR:
				fmt.Println("STAR * null")
				indx++
			
			
			case QUOTE:
				var result []rune
				indx++
				for indx < len(fileContents) && fileContents[indx] != '"' {
					if fileContents[indx] == '\n' {
						line++
					}
					result = append(result, rune(fileContents[indx]))
					indx++
				}

				if indx >= len(fileContents) {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
					has_error = true
				} else {
					indx++
					fmt.Printf("STRING \"%s\" %s\n", string(result), string(result))
				}
				
			
			case EQUAL:
				if (indx + 1 < len(fileContents) && fileContents[indx+1] == '='){
					fmt.Println("EQUAL_EQUAL == null")
					indx += 2
				} else{
					fmt.Println("EQUAL = null")
					indx++
				}
			
			case BANG:
				if (indx + 1 < len(fileContents) && fileContents[indx+1] == '='){
					fmt.Println("BANG_EQUAL != null")
					indx += 2
				} else{
					fmt.Println("BANG ! null")
					indx++
				}
			
			case LESS:
				if (indx + 1 < len(fileContents) && fileContents[indx+1] == '='){
					fmt.Println("LESS_EQUAL <= null")
					indx += 2
				} else{
					fmt.Println("LESS < null")
					indx++
				}
			
			case GREATER:
				if (indx + 1 < len(fileContents) && fileContents[indx+1] == '='){
					fmt.Println("GREATER_EQUAL >= null")
					indx += 2
				} else{
					fmt.Println("GREATER > null")
					indx++
				}
				
			case SLASH:
				if (indx + 1 < len(fileContents) && fileContents[indx+1] == '/'){
					indx += 2
					for indx < len(fileContents) && fileContents[indx] != '\n'{
						indx++
					}
					
				} else{
					fmt.Println("SLASH / null")
					indx++
				}


			case NUMBER, DOLLAR, ATSIGN, CARET, PERCENT:
				
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, ch)
				has_error = true
				indx++
			default:
				if (unicode.IsDigit(ch)){
					var result []rune
					for indx < len(fileContents) && fileContents[indx] != ' ' {
						if (rune(fileContents[indx]) == '.' || unicode.IsDigit(rune(fileContents[indx]))){
							result = append(result, rune(fileContents[indx]))
							indx++
						} else {
							break
						}	
					}
					floatVal, err := strconv.ParseFloat(string(result), 64)
					if err!=nil {
						fmt.Fprintf(os.Stderr, "[line %d] Error: Failed to parse number", line)
						has_error = true
					} else{
						if floatVal == float64(int(floatVal)) {
							// It's an integer, print without decimals
							fmt.Printf("NUMBER %s %.1f\n", string(result), floatVal)
						} else {
							// It's a float, print with decimals
							fmt.Printf("NUMBER %s %g\n", string(result), floatVal)
						}
					}
					
				} else if (unicode.IsLetter(ch) || ch == '_'){
					var result []rune
					for indx < len(fileContents) && fileContents[indx] != ' ' {
						if (rune(fileContents[indx]) == '_' || unicode.IsDigit(rune(fileContents[indx])) || unicode.IsLetter(rune(fileContents[indx]))){
							result = append(result, rune(fileContents[indx]))
							indx++
						} else {
							break
						}	
					} 
					fmt.Printf("IDENTIFIER %s null\n", string(result))
				} else{
					indx++
				}
				
			
			}
			}
		} 
		fmt.Println("EOF  null")
		if (has_error){
			os.Exit(65)
	}
}
