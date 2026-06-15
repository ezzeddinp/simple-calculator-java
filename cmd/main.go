package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	total    float64
	isFirst  bool
	scanner  *bufio.Scanner
}

func NewCalculator() *Calculator {
	return &Calculator{
		total:    0,
		isFirst:  true,
		scanner:  bufio.NewScanner(os.Stdin),
	}
}

func (c *Calculator) Run() {
	for {
		c.displayHeader()
		
		if c.isFirst {
			c.promptAndSetInitialNumber()
			c.isFirst = false
		}
		
		operator := c.promptOperator()
		if operator == 7 {
			fmt.Println("Exiting the calculator. Goodbye!")
			return
		}
		
		c.executeOperation(operator)
		fmt.Println() // Empty line for readability
	}
}

func (c *Calculator) displayHeader() {
	fmt.Println("Calculator")
	fmt.Println("==========")
	fmt.Printf("Total: %.2f\n", c.total)
}

func (c *Calculator) promptAndSetInitialNumber() {
	fmt.Print("Input number: ")
	if c.scanner.Scan() {
		input := strings.TrimSpace(c.scanner.Text())
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			c.total = num
		}
	}
}

func (c *Calculator) promptOperator() int {
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
		if op, err := strconv.Atoi(input); err == nil && op >= 1 && op <= 7 {
			return op
		}
		fmt.Println("Invalid operator. Please select a valid option.")
	}
	return c.promptOperator() // Retry on invalid input
}

func (c *Calculator) executeOperation(operator int) {
	switch operator {
	case 1: // Addition
		c.performAddition()
	case 2: // Subtraction
		c.performSubtraction()
	case 3: // Multiplication
		c.performMultiplication()
	case 4: // Division
		c.performDivision()
	case 5: // Modulo
		c.performModulo()
	case 6: // Clear Total
		c.clearTotal()
	}
}

func (c *Calculator) performAddition() {
	if num := c.promptNumber(); num != 0 {
		c.total += num
	}
}

func (c *Calculator) performSubtraction() {
	if num := c.promptNumber(); num != 0 {
		c.total -= num
	}
}

func (c *Calculator) performMultiplication() {
	if num := c.promptNumber(); num != 0 {
		c.total *= num
	}
}

func (c *Calculator) performDivision() {
	if num := c.promptNumber(); num != 0 {
		if num != 0 {
			c.total /= num
		} else {
			fmt.Print("no")
		}
	}
}

func (c *Calculator) performModulo() {
	if num := c.promptNumber(); num != 0 {
		c.total = float64(int(c.total) % int(num))
	}
}

func (c *Calculator) clearTotal() {
	c.total = 0
	c.isFirst = true
}

func (c *Calculator) promptNumber() float64 {
	fmt.Print("Input number: ")
	if c.scanner.Scan() {
		input := strings.TrimSpace(c.scanner.Text())
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			return num
		}
	}
	return 0
}

func main() {
	calculator := NewCalculator()
	calculator.Run()
}
