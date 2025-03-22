package main

import "fmt"

func getNumber(prompt string) float64 {
    var num float64
    for {
        fmt.Print(prompt)
        _, err := fmt.Scanln(&num)
        if err == nil {
            break
        }
        fmt.Println("Invalid number. Please enter a valid number.")
        var discard string
        fmt.Scanln(&discard)
    }
    return num
}
func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	} else {
		return 0
	}
}
func main() {
	var num1, num2 float64
	var operator string
	var keepGoing = true

	num1 = getNumber("Enter first number: ")
	num2 = getNumber("Enter second number: ")

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)	
	for operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operator. Please enter a valid operator.")
		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)
	}


	for keepGoing {
		switch operator {
			case "+": fmt.Printf("%f + %f = %f\n", num1, num2, add(num1, num2))
			case "-": fmt.Printf("%f - %f = %f\n", num1, num2, subtract(num1, num2))
			case "*": fmt.Printf("%f * %f = %f\n", num1, num2, multiply(num1, num2))
			case "/": if num2 != 0 {
				fmt.Printf("%f / %f = %f\n", num1, num2, divide(num1, num2))
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
			num1 = getNumber("Enter first number: ")
			num2 = getNumber("Enter second number: ")

			fmt.Print("Enter operator (+, -, *, /): ")
			fmt.Scanln(&operator)	
			for operator != "+" && operator != "-" && operator != "*" && operator != "/" {
				fmt.Println("Invalid operator. Please enter a valid operator.")
				fmt.Print("Enter operator (+, -, *, /): ")
				fmt.Scanln(&operator)
			}
		}

	}
}