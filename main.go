package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	var keepGoing = true

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)


	for keepGoing {
		switch operator {
			case "+": fmt.Printf("%f + %f = %f\n", num1, num2, num1 + num2)
			case "-": fmt.Printf("%f - %f = %f\n", num1, num2, num1 - num2)
			case "*": fmt.Printf("%f * %f = %f\n", num1, num2, num1 * num2)
			case "/": if num2 != 0 {
				fmt.Printf("%f / %f = %f\n", num1, num2, num1 / num2)
			} else {
				fmt.Println("Cannot divide by zero")
			}
			default: fmt.Println("Invalid operator")
		}

		fmt.Print("Do you want to continue? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if choice == "n" {
			keepGoing = false
		} else {
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)

			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)

			fmt.Print("Enter operator (+, -, *, /): ")
			fmt.Scanln(&operator)
		}

	}

	

}