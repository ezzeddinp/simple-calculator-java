import org.junit.jupiter.api.*;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.io.InputStream;

import static org.junit.jupiter.api.Assertions.*;

public class CalculatorTest {
    private Calculator calculator;
    private final ByteArrayOutputStream outputStream = new ByteArrayOutputStream();
    private final PrintStream originalOut = System.out;
    private InputStream originalIn;
    
    @BeforeEach
    void setUp() {
        // Redirect System.out to capture output
        System.setOut(new PrintStream(outputStream));
        originalIn = System.in;
        
        // Reset static variables before each test
        Calculator.total = 0;
        Calculator.fr = true;
        Calculator.operator = 0;
    }
    
    @AfterEach
    void tearDown() {
        // Restore original System.out and System.in
        System.setOut(originalOut);
        System.setIn(originalIn);
    }
    
    @Test
    void testInitialState() {
        assertEquals(0, Calculator.total, "Total should start at 0");
        assertTrue(Calculator.fr, "First run flag should be true initially");
    }
    
    @Test
    void testAdditionOperation() {
        // Simulate input: number 5, operator 1 (addition), number 3
        String input = "5\n1\n3\n7\n";
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        
        Calculator.main(new String[]{});
        
        String output = outputStream.toString();
        assertTrue(output.contains("Total: 8"), "Should show total after addition");
    }
    
    @Test
    void testSubtractionOperation() {
        // Set initial total to 10, then subtract 3
        Calculator.total = 10;
        Calculator.fr = false;
        
        String input = "2\n3\n7\n"; // operator 2 (subtraction), number 3
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        
        // We need to test the op() method directly
        System.setIn(new ByteArrayInputStream("2\n3\n".getBytes()));
        Calculator.scanner = new Scanner(System.in);
        Calculator.op();
        
        assertEquals(7, Calculator.total, 0.001, "10 - 3 should equal 7");
    }
    
    @Test
    void testMultiplicationOperation() {
        Calculator.total = 5;
        Calculator.fr = false;
        
        String input = "3\n4\n"; // operator 3 (multiplication), number 4
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        assertEquals(20, Calculator.total, 0.001, "5 * 4 should equal 20");
    }
    
    @Test
    void testDivisionOperation() {
        Calculator.total = 20;
        Calculator.fr = false;
        
        String input = "4\n5\n"; // operator 4 (division), number 5
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        assertEquals(4, Calculator.total, 0.001, "20 / 5 should equal 4");
    }
    
    @Test
    void testDivisionByZero() {
        Calculator.total = 20;
        Calculator.fr = false;
        
        String input = "4\n0\n"; // operator 4 (division), number 0
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        String output = outputStream.toString();
        assertTrue(output.contains("no"), "Should display 'no' for division by zero");
        assertEquals(20, Calculator.total, 0.001, "Total should remain unchanged after division by zero");
    }
    
    @Test
    void testModuloOperation() {
        Calculator.total = 10;
        Calculator.fr = false;
        
        String input = "5\n3\n"; // operator 5 (modulo), number 3
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        assertEquals(1, Calculator.total, 0.001, "10 % 3 should equal 1");
    }
    
    @Test
    void testClearTotalOperation() {
        Calculator.total = 100;
        Calculator.fr = false;
        
        String input = "6\n"; // operator 6 (clear)
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        assertEquals(0, Calculator.total, "Total should be 0 after clear");
        assertTrue(Calculator.fr, "First run flag should be true after clear");
    }
    
    @Test
    void testExitOperation() {
        Calculator.fr = false;
        
        String input = "7\n"; // operator 7 (exit)
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        String output = outputStream.toString();
        assertTrue(output.contains("Exiting the calculator"), "Should display exit message");
    }
    
    @Test
    void testInvalidOperator() {
        Calculator.fr = false;
        
        String input = "99\n"; // invalid operator
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        Calculator.scanner = new Scanner(System.in);
        
        Calculator.op();
        
        String output = outputStream.toString();
        assertTrue(output.contains("Invalid operator"), "Should display invalid operator message");
    }
    
    @Test
    void testFirstRunFlow() {
        // Test the complete first run: input number, then addition
        String input = "10\n1\n5\n7\n";
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        
        Calculator.main(new String[]{});
        
        String output = outputStream.toString();
        assertTrue(output.contains("Total: 15"), "First run should add initial number");
        assertTrue(output.contains("Calculator"), "Should display calculator header");
    }
    
    @Test
    void testMultipleOperationsSequence() {
        // Test sequence: 10 + 5 - 3 * 2
        String input = "10\n1\n5\n2\n3\n3\n2\n7\n";
        System.setIn(new ByteArrayInputStream(input.getBytes()));
        
        Calculator.main(new String[]{});
        
        // 10 + 5 = 15, 15 - 3 = 12, 12 * 2 = 24
        String output = outputStream.toString();
        assertTrue(output.contains("Total: 24") || output.contains("Total: 24.0"), 
                  "Should correctly calculate sequence of operations");
    }
}
