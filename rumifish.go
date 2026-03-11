package main

import (
	"fmt"
	"os"
)

func eval(code string, accumulator *int) {
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case 'i':
			*accumulator += 1
		case 'd':
			*accumulator -= 1
		case 's':
			*accumulator *= *accumulator
		case 'o':
			fmt.Println(*accumulator)
		case 'h':
			return
		default:
			fmt.Println()
		}

		if *accumulator == 256 || *accumulator == -1 {
			*accumulator = 0
		}
	}
}

func main() {
	var accumulator int = 0

	if len(os.Args[1:]) > 0 {
		file, err := os.ReadFile(os.Args[1:][0])

		if err != nil {
			panic("file not found")
		} else {
			eval(string(file), &accumulator)
			return
		}
	}

	fmt.Println("rumifish || made by snarkb0t on github.")
	fmt.Println("enter 'q', 'quit' or 'exit' to exit.")
	fmt.Println("")

	for {
		var line string

		fmt.Print(">>> ")
		fmt.Scan(&line)

		if line == "q" || line == "quit" || line == "exit" {
			return
		}

		eval(line, &accumulator)
	}
}
