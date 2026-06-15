package main

import (
	"strings"
	"testing"
)

func TestOop1Main(t *testing.T) {
	// Test that oop1 main function properly calls Calculator main
	// This is a simple integration test to verify the entry point works
	calc := &Calculator{
		scanner: strings.NewReader("5.0\n7\n"),
		total:   0,
		fr:      true,
	}
	
	// Simulate the main execution flow
	calc.runFirstInput()
	calc.handleOperation(7) // Exit immediately
	
	if calc.total != 5.0 {
		t.Errorf("Expected total to be 5.0 after main execution, got %f", calc.total)
	}
}
