import org.junit.jupiter.api.*;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.io.InputStream;

import static org.junit.jupiter.api.Assertions.*;

/**
 * Test for a refactored version of Calculator that separates business logic from I/O
 * This demonstrates how the code could be structured for better testability
 */
public class CalculatorRefactoredTest {
    
    // Refactored calculator logic without static dependencies
    static class TestableCalculator {
        private float total = 0;
        private boolean isFirstRun = true;
        
        public float getTotal() {
            return total;
        }
        
        public void setTotal(float total) {
            this.total = total;
        }
        
        public boolean isFirstRun() {
            return isFirstRun;
        }
        
        public void setFirstRun(boolean firstRun) {
            this.isFirstRun = firstRun;
        }
        
        public float performOperation(int operator, float operand) {
            switch (operator) {
                case 1: // Addition
                    total += operand;
                    break;
                case 2: // Subtraction
                    total -= operand;
                    break;
                case 3: // Multiplication
                    total *= operand;
                    break;
                case 4: // Division
                    if (operand != 0) {
                        total /= operand;
                    } else {
                        throw new ArithmeticException("Division by zero");
                    }
                    break;
                case 5: // Modulo
                    total %= operand;
                    break;
                case 6: // Clear Total
                    total = 0;
                    isFirstRun = true;
                    break;
                default:
                    throw new IllegalArgumentException("Invalid operator: " + operator);
            }
            return total;
        }
        
        public void handleFirstRun(float initialNumber) {
            total = initialNumber;
            isFirstRun = false;
        }
    }
    
    @Test
    void testPerformAddition() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(10);
        
        float result = calc.performOperation(1, 5); // 10 + 5
        
        assertEquals(15, result, 0.001);
        assertEquals(15, calc.getTotal(), 0.001);
    }
    
    @Test
    void testPerformSubtraction() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(20);
        
        float result = calc.performOperation(2, 7); // 20 - 7
        
        assertEquals(13, result, 0.001);
        assertEquals(13, calc.getTotal(), 0.001);
    }
    
    @Test
    void testPerformMultiplication() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(6);
        
        float result = calc.performOperation(3, 4); // 6 * 4
        
        assertEquals(24, result, 0.001);
        assertEquals(24, calc.getTotal(), 0.001);
    }
    
    @Test
    void testPerformDivision() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(15);
        
        float result = calc.performOperation(4, 3); // 15 / 3
        
        assertEquals(5, result, 0.001);
        assertEquals(5, calc.getTotal(), 0.001);
    }
    
    @Test
    void testPerformDivisionByZero() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(15);
        
        Exception exception = assertThrows(ArithmeticException.class, () -> {
            calc.performOperation(4, 0); // 15 / 0
        });
        
        assertEquals("Division by zero", exception.getMessage());
        assertEquals(15, calc.getTotal(), 0.001); // Total should remain unchanged
    }
    
    @Test
    void testPerformModulo() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(17);
        
        float result = calc.performOperation(5, 5); // 17 % 5
        
        assertEquals(2, result, 0.001);
        assertEquals(2, calc.getTotal(), 0.001);
    }
    
    @Test
    void testClearTotal() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(100);
        calc.setFirstRun(false);
        
        float result = calc.performOperation(6, 0); // Clear
        
        assertEquals(0, result, 0.001);
        assertEquals(0, calc.getTotal(), 0.001);
        assertTrue(calc.isFirstRun());
    }
    
    @Test
    void testInvalidOperator() {
        TestableCalculator calc = new TestableCalculator();
        calc.setTotal(10);
        
        Exception exception = assertThrows(IllegalArgumentException.class, () -> {
            calc.performOperation(99, 5); // Invalid operator
        });
        
        assertTrue(exception.getMessage().contains("Invalid operator"));
        assertEquals(10, calc.getTotal(), 0.001); // Total should remain unchanged
    }
    
    @Test
    void testHandleFirstRun() {
        TestableCalculator calc = new TestableCalculator();
        
        assertTrue(calc.isFirstRun());
        assertEquals(0, calc.getTotal(), 0.001);
        
        calc.handleFirstRun(25);
        
        assertFalse(calc.isFirstRun());
        assertEquals(25, calc.getTotal(), 0.001);
    }
    
    @Test
    void testComplexOperationSequence() {
        TestableCalculator calc = new TestableCalculator();
        
        // Start with 10
        calc.handleFirstRun(10);
        
        // 10 + 5 = 15
        calc.performOperation(1, 5);
        assertEquals(15, calc.getTotal(), 0.001);
        
        // 15 - 3 = 12
        calc.performOperation(2, 3);
        assertEquals(12, calc.getTotal(), 0.001);
        
        // 12 * 2 = 24
        calc.performOperation(3, 2);
        assertEquals(24, calc.getTotal(), 0.001);
        
        // 24 / 6 = 4
        calc.performOperation(4, 6);
        assertEquals(4, calc.getTotal(), 0.001);
        
        // 4 % 3 = 1
        calc.performOperation(5, 3);
        assertEquals(1, calc.getTotal(), 0.001);
        
        // Clear to 0
        calc.performOperation(6, 0);
        assertEquals(0, calc.getTotal(), 0.001);
        assertTrue(calc.isFirstRun());
    }
}
