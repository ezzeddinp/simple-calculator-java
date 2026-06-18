package main

import (
	"bytes"
	"strings"
	"testing"

	"calculator/internal/scanner"
)

func TestCalculator_Addition(t *testing.T) {
	input := strings.NewReader("5\n1\n3\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	// Simulate the Run loop manually for testing
	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 8.00") {
		t.Errorf("Addition failed: expected Total: 8.00, got output:\n%s", output.String())
	}
}

func TestCalculator_Subtraction(t *testing.T) {
	input := strings.NewReader("10\n2\n4\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 6.00") {
		t.Errorf("Subtraction failed: expected Total: 6.00, got output:\n%s", output.String())
	}
}

func TestCalculator_Multiplication(t *testing.T) {
	input := strings.NewReader("2\n3\n5\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 10.00") {
		t.Errorf("Multiplication failed: expected Total: 10.00, got output:\n%s", output.String())
	}
}

func TestCalculator_Division(t *testing.T) {
	input := strings.NewReader("20\n4\n4\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 5.00") {
		t.Errorf("Division failed: expected Total: 5.00, got output:\n%s", output.String())
	}
}

func TestCalculator_DivisionByZero(t *testing.T) {
	input := strings.NewReader("10\n4\n0\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "no") {
		t.Errorf("Division by zero failed: expected 'no' in output, got:\n%s", output.String())
	}
}

func TestCalculator_Modulo(t *testing.T) {
	input := strings.NewReader("10\n5\n3\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 1.00") {
		t.Errorf("Modulo failed: expected Total: 1.00, got output:\n%s", output.String())
	}
}

func TestCalculator_ClearTotal(t *testing.T) {
	input := strings.NewReader("5\n6\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Total: 0.00") {
		t.Errorf("Clear total failed: expected Total: 0.00, got output:\n%s", output.String())
	}
}

func TestCalculator_InvalidOperator(t *testing.T) {
	input := strings.NewReader("5\n99\n7\n")
	output := &bytes.Buffer{}
	calc := NewCalculator()
	sc := scanner.NewScanner(input, output)

	calc.fr = true
	var operator int
	for {
		output.WriteString("Calculator\n")
		output.WriteString("==========\n")
		output.WriteString("Total: " + formatFloat(calc.total) + "\n")

		if calc.fr {
			output.WriteString("Input number: ")
			a, err := sc.ReadFloat()
			if err != nil {
				output.WriteString("Invalid input\n")
				continue
			}
			calc.total += a
			calc.fr = false
			operator = simulateOp(calc, sc, output)
		} else {
			operator = simulateOp(calc, sc, output)
		}

		if operator == 7 {
			output.WriteString("Exiting the calculator. Goodbye!\n")
			break
		}
	}

	if !strings.Contains(output.String(), "Invalid operator") {
		t.Errorf("Invalid operator handling failed: expected 'Invalid operator' in output, got:\n%s", output.String())
	}
}

// Helper functions
func formatFloat(f float64) string {
	// Simple formatting to match the calculator's output
	return strings.TrimRight(strings.TrimRight(format.Sprintf("%.2f", f), "0"), ".")
}

func simulateOp(c *Calculator, s *scanner.Scanner, output *bytes.Buffer) int {
	output.WriteString("1. Addition\n")
	output.WriteString("2. Subtraction\n")
	output.WriteString("3. Multiplication\n")
	output.WriteString("4. Division\n")
	output.WriteString("5. Modulo\n")
	output.WriteString("6. Clear Total\n")
	output.WriteString("7. Exit\n")
	output.WriteString("Operator [1...7]: ")

	operator, err := s.ReadInt()
	if err != nil {
		output.WriteString("Invalid operator\n")
		return 0
	}

	switch operator {
	case 1:
		output.WriteString("Input number: ")
		a, err := s.ReadFloat()
		if err != nil {
			output.WriteString("Invalid input\n")
			break
		}
		c.total += a
	case 2:
		output.WriteString("Input number: ")
		a, err := s.ReadFloat()
		if err != nil {
			output.WriteString("Invalid input\n")
			break
		}
		c.total -= a
	case 3:
		output.WriteString("Input number: ")
		a, err := s.ReadFloat()
		if err != nil {
			output.WriteString("Invalid input\n")
			break
		}
		c.total *= a
	case 4:
		output.WriteString("Input number: ")
		a, err := s.ReadFloat()
		if err != nil {
			output.WriteString("Invalid input\n")
			break
		}
		if a != 0 {
			c.total /= a
		} else {
			output.WriteString("no\n")
		}
	case 5:
		output.WriteString("Input number: ")
		a, err := s.ReadFloat()
		if err != nil {
			output.WriteString("Invalid input\n")
			break
		}
		c.total = float64(int(c.total) % int(a))
	case 6:
		c.total = 0
		c.fr = true
	case 7:
		return 7
	default:
		output.WriteString("Invalid operator. Please select a valid option.\n")
	}

	output.WriteString("\n")
	return operator
}
