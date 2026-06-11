package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	scanner  *bufio.Scanner
	a        float64
	fr       bool
	total    float64
	operator int
}

func NewCalculator() *Calculator {
	return &Calculator{
		scanner: bufio.NewScanner(os.Stdin),
		fr:      true,
		total:   0,
	}
}

func (c *Calculator) Run() {
	for {
		fmt.Println("Calculator")
		fmt.Println("==========")
		fmt.Printf("Total: %.2f\n", c.total)
		
		if c.fr {
			fmt.Print("Input number: ")
			if c.scanner.Scan() {
				input := strings.TrimSpace(c.scanner.Text())
				num, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("Invalid number. Please try again.")
					continue
				}
				c.a = num
				c.total += c.a
				c.fr = false
			}
			c.op()
		} else {
			c.op()
		}
		
		if c.operator == 7 {
			break
		}
	}
}

func (c *Calculator) op() {
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Modulo")
	fmt.Println("6. Clear Total")
	fmt.Println("7. Exit")
	fmt.Print("Operator [1...7]: ")
	
	if c.scanner.Scan() {
		input := strings.TrimSpace(c.scanner.Text())
		op, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid operator. Please select a valid option.")
			return
		}
		c.operator = op
	}

	switch c.operator {
	case 1: // Addition
		fmt.Print("Input number: ")
		if c.scanner.Scan() {
			input := strings.TrimSpace(c.scanner.Text())
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				break
			}
			c.a = num
			c.total += c.a
		}
	case 2: // Subtraction
		fmt.Print("Input number: ")
		if c.scanner.Scan() {
			input := strings.TrimSpace(c.scanner.Text())
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				break
			}
			c.a = num
			c.total -= c.a
		}
	case 3: // Multiplication
		fmt.Print("Input number: ")
		if c.scanner.Scan() {
			input := strings.TrimSpace(c.scanner.Text())
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				break
			}
			c.a = num
			c.total *= c.a
		}
	case 4: // Division
		fmt.Print("Input number: ")
		if c.scanner.Scan() {
			input := strings.TrimSpace(c.scanner.Text())
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				break
			}
			c.a = num
			if c.a != 0 {
				c.total /= c.a
			} else {
				fmt.Println("no")
			}
		}
	case 5: // Modulo
		fmt.Print("Input number: ")
		if c.scanner.Scan() {
			input := strings.TrimSpace(c.scanner.Text())
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
				break
			}
			c.a = num
			c.total = float64(int(c.total) % int(c.a))
		}
	case 6: // Clear Total
		c.total = 0
		c.fr = true // reset
	case 7: // Exit
		fmt.Println("Exiting the calculator. Goodbye!")
		return
	default:
		fmt.Println("Invalid operator. Please select a valid option.")
	}

	fmt.Println() // Print an empty line for better readability
}

func main() {
	calculator := NewCalculator()
	calculator.Run()
}
