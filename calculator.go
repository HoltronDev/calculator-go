package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calc Program")
	fmt.Println("Supported operations:")
	fmt.Println("X + Y")
	fmt.Println("X - Y")
	fmt.Println("X * Y")
	fmt.Println("X / Y")

	input := ""
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "exit" {
			fmt.Println("Good bye.")
			return
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Printf("Not enough arguments.")
			return
		}

		first, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Printf("Error parsing first number: %v", err)
			return
		}

		second, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Printf("Error parsing first number: %v", err)
			return
		}

		switch op := parts[1]; op {
		case "+":
			fmt.Printf("Answer: %v\n", first+second)
		case "-":
			fmt.Printf("Answer: %v\n", first-second)
		case "*":
			fmt.Printf("Answer: %v\n", first*second)
		case "/":
			fmt.Printf("Answer: %v\n", first/second)
		default:
			fmt.Printf("Input %s is not valid.", op)
		}
	}
}
