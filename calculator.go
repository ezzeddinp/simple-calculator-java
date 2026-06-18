package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Calculator struct {
	scanner *bufio.Scanner
	total   float64
	fr      bool
}

func NewCalculator() *Calculator {
	return &Calculator{
		scanner: bufio.NewScanner(os.Stdin),
		total:   0,
		fr:      true,
	}
}

func (c *Calculator) readFloat() float64 {
	c.scanner.Scan()
	input := c.scanner.Text()
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid number, please try again.")
		return c.readFloat()
	}
	return val
}

func (c *Calculator) readInt() int {
	c.scanner.Scan()
	input := c.scanner.Text()
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid integer, please try again.")
		return c.readInt()
	}
	return val
}

func (c *Calculator) run() {
	var operator int
	for operator != 7 {
		fmt.Println("Calculator")
		fmt.Println("==========")
		fmt.Printf("Total: %v\n", c.total)
		if c.fr {
			fmt.Print("Input number:")
			a := c.readFloat()
			c.total += a
			c.fr = false
			operator = c.op()
		} else {
			operator = c.op()
		}
	}
}

func (c *Calculator) op() int {
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Modulo")
	fmt.Println("6. Clear Total")
	fmt.Println("7. Exit")
	fmt.Print("Operator [1...7]: ")
	operator := c.readInt()

	switch operator {
	case 1: // Addition
		fmt.Print("Input number :")
		a := c.readFloat()
		c.total += a
	case 2: // Subtraction
		fmt.Print("Input number :")
		a := c.readFloat()
		c.total -= a
	case 3: // Multiplication
		fmt.Print("Input number :")
		a := c.readFloat()
		c.total *= a
	case 4: // Division
		fmt.Print("Input number:")
		a := c.readFloat()
		if a != 0 {
			c.total /= a
		} else {
			fmt.Println("no")
		}
	case 5: // Modulo
		fmt.Print("Input number:")
		a := c.readFloat()
		c.total = float64(int(c.total) % int(a))
	case 6: // Clear Total
		c.total = 0
		c.fr = true
	case 7: // Exit
		fmt.Println("Exiting the calculator. Goodbye!")
	default:
		fmt.Println("Invalid operator. Please select a valid option.")
	}

	fmt.Println()
	return operator
}

func main() {
	calc := NewCalculator()
	calc.run()
}
