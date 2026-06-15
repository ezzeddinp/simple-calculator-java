package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCalculatorInitialization(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader(""),
		total:   0,
		fr:      true,
	}
	
	if calc.total != 0 {
		t.Errorf("Expected initial total to be 0, got %f", calc.total)
	}
	if !calc.fr {
		t.Error("Expected fr to be true on initialization")
	}
}

func TestCalculatorAddition(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("1\n5.5\n7\n"),
		total:   10.0,
		fr:      false,
	}
	
	// Test addition operation
	calc.handleOperation(1)
	if calc.total != 15.5 {
		t.Errorf("Expected total after addition to be 15.5, got %f", calc.total)
	}
}

func TestCalculatorSubtraction(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("2\n3.5\n7\n"),
		total:   10.0,
		fr:      false,
	}
	
	calc.handleOperation(2)
	if calc.total != 6.5 {
		t.Errorf("Expected total after subtraction to be 6.5, got %f", calc.total)
	}
}

func TestCalculatorMultiplication(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("3\n2.5\n7\n"),
		total:   4.0,
		fr:      false,
	}
	
	calc.handleOperation(3)
	if calc.total != 10.0 {
		t.Errorf("Expected total after multiplication to be 10.0, got %f", calc.total)
	}
}

func TestCalculatorDivision(t *testing.T) {
	t.Run("ValidDivision", func(t *testing.T) {
		calc := &Calculator{
			scanner: strings.NewReader("4\n2.0\n7\n"),
			total:   10.0,
			fr:      false,
		}
		
		calc.handleOperation(4)
		if calc.total != 5.0 {
			t.Errorf("Expected total after division to be 5.0, got %f", calc.total)
		}
	})
	
	t.Run("DivisionByZero", func(t *testing.T) {
		var buf bytes.Buffer
		calc := &Calculator{
			scanner: strings.NewReader("4\n0\n7\n"),
			total:   10.0,
			fr:      false,
			output:  &buf,
		}
		
		calc.handleOperation(4)
		if calc.total != 10.0 {
			t.Errorf("Expected total to remain 10.0 after division by zero, got %f", calc.total)
		}
		if !strings.Contains(buf.String(), "no") {
			t.Error("Expected 'no' message for division by zero")
		}
	})
}

func TestCalculatorModulo(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("5\n3.0\n7\n"),
		total:   10.0,
		fr:      false,
	}
	
	calc.handleOperation(5)
	if calc.total != 1.0 {
		t.Errorf("Expected total after modulo to be 1.0, got %f", calc.total)
	}
}

func TestCalculatorClear(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("6\n7\n"),
		total:   15.0,
		fr:      false,
	}
	
	calc.handleOperation(6)
	if calc.total != 0 {
		t.Errorf("Expected total after clear to be 0, got %f", calc.total)
	}
	if !calc.fr {
		t.Error("Expected fr to be true after clear operation")
	}
}

func TestCalculatorExit(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader(""),
		total:   10.0,
		fr:      false,
	}
	
	calc.handleOperation(7)
	// Exit operation should not modify total
	if calc.total != 10.0 {
		t.Errorf("Expected total to remain 10.0 after exit, got %f", calc.total)
	}
}

func TestCalculatorInvalidOperator(t *testing.T) {
	var buf bytes.Buffer
	calc := &Calculator{
		scanner: strings.NewReader(""),
		total:   10.0,
		fr:      false,
		output:  &buf,
	}
	
	calc.handleOperation(8) // Invalid operator
	if !strings.Contains(buf.String(), "Invalid operator") {
		t.Error("Expected invalid operator error message")
	}
}

func TestCalculatorFirstRun(t *testing.T) {
	calc := &Calculator{
		scanner: strings.NewReader("5.5\n1\n2.5\n7\n"),
		total:   0,
		fr:      true,
	}
	
	// Simulate first run behavior
	calc.runFirstInput()
	if calc.total != 5.5 {
		t.Errorf("Expected total after first input to be 5.5, got %f", calc.total)
	}
	if calc.fr {
		t.Error("Expected fr to be false after first run")
	}
}

func TestCalculatorIntegration(t *testing.T) {
	input := "10.0\n1\n5.0\n3\n2.0\n6\n7\n"
	calc := &Calculator{
		scanner: strings.NewReader(input),
		total:   0,
		fr:      true,
	}
	
	// Simulate complete workflow
	calc.runFirstInput() // First number: 10.0
	
	calc.handleOperation(1) // Add 5.0 -> total = 15.0
	if calc.total != 15.0 {
		t.Errorf("After addition: expected 15.0, got %f", calc.total)
	}
	
	calc.handleOperation(3) // Multiply by 2.0 -> total = 30.0
	if calc.total != 30.0 {
		t.Errorf("After multiplication: expected 30.0, got %f", calc.total)
	}
	
	calc.handleOperation(6) // Clear -> total = 0
	if calc.total != 0 {
		t.Errorf("After clear: expected 0, got %f", calc.total)
	}
	
	calc.handleOperation(7) // Exit
	// Should not affect total
	if calc.total != 0 {
		t.Errorf("After exit: expected 0, got %f", calc.total)
	}
}
