package main

import (
	"fmt"
	"os"
)

type Calculator struct {
	total float64
	fr    bool
}

func NewCalculator() *Calculator {
	return &Calculator{
		fr: true,
	}
}

func (c *Calculator) Run() {
	scanner := NewScanner(os.Stdin, os.Stdout)
	var operator int

	for {
		fmt.Println("Calculator")
		fmt.Println("==========")
		fmt.Printf("Total: %.2f\n", c.total)

		if c.fr {
			fmt.Print("Input number: ")
			a, err := scanner.ReadFloat()
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			c.total += a
			c.fr = false
			operator = c.op(scanner)
		} else {
			operator = c.op(scanner)
		}

		if operator == 7 {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}
	}
}

func (c *Calculator) op(scanner *Scanner) int {
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Modulo")
	fmt.Println("6. Clear Total")
	fmt.Println("7. Exit")
	fmt.Print("Operator [1...7]: ")

	operator, err := scanner.ReadInt()
	if err != nil {
		fmt.Println("Invalid operator")
		return 0
	}

	switch operator {
	case 1: // Addition
		fmt.Print("Input number: ")
		a, err := scanner.ReadFloat()
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		c.total += a
	case 2: // Subtraction
		fmt.Print("Input number: ")
		a, err := scanner.ReadFloat()
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		c.total -= a
	case 3: // Multiplication
		fmt.Print("Input number: ")
		a, err := scanner.ReadFloat()
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		c.total *= a
	case 4: // Division
		fmt.Print("Input number: ")
		a, err := scanner.ReadFloat()
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		if a != 0 {
			c.total /= a
		} else {
			fmt.Println("no")
		}
	case 5: // Modulo
		fmt.Print("Input number: ")
		a, err := scanner.ReadFloat()
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		c.total = float64(int(c.total) % int(a))
	case 6: // Clear Total
		c.total = 0
		c.fr = true
	case 7: // Exit
		return 7
	default:
		fmt.Println("Invalid operator. Please select a valid option.")
	}

	fmt.Println()
	return operator
}

func main() {
	calc := NewCalculator()
	calc.Run()
}
