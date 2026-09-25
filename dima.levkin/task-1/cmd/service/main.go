package main

import "fmt"

func main() {
	var (
		operation     string
		firstOperand  int
		secondOperand int
	)
	_, err := fmt.Scan(&firstOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&secondOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	if secondOperand == 0 && operation == "/" {
		fmt.Println("Division by zero")
		return
	}

	var result = firstOperand

	switch operation {
	case "+":
		result += secondOperand
	case "-":
		result -= secondOperand
	case "*":
		result *= secondOperand
	case "/":
		result /= secondOperand
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
