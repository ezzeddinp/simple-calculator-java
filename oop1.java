import java.util.*;

class Calculator {
    static Scanner scanner = new Scanner(System.in);
    static float a;
    static boolean fr = true;
    static float total = 0;
    static int operator;

    public static void main(String[] args) {

        do {
            System.out.println("Calculator");
            System.out.println("==========");
            System.out.println("Total: " + total);
            if (fr) {
                System.out.print("Input number:");
                a = scanner.nextFloat();
                total += a;
                fr = false;
                op();
            } else { // run ke 2/n
                op();
            }
        } while (operator != 7);

    }

    static void op() {
        System.out.println("1. Addition");
        System.out.println("2. Subtraction");
        System.out.println("3. Multiplication");
        System.out.println("4. Division");
        System.out.println("5. Modulo");
        System.out.println("6. Clear Total");
        System.out.println("7. Exit");
        System.out.print("Operator [1...7]: ");
        operator = scanner.nextInt();
        scanner.nextLine();

        switch (operator) {
            case 1: // Addition
                System.out.print("Input number :");
                a = scanner.nextFloat();
                total += a;
                break;
            case 2: // Subtraction
                System.out.print("Input number :");
                a = scanner.nextFloat();
                total -= a;
                break;
            case 3: // Multiplication
                System.out.print("Input number :");
                a = scanner.nextFloat();
                total *= a;
                break;
            case 4: // Division
                System.out.print("Input number:");
                a = scanner.nextFloat();
                if (a != 0) {
                    total /= a;
                } else {
                    System.out.print("no");
                }
                break;
            case 5: // Modulo
                System.out.print("Input number:");
                a = scanner.nextFloat();
                total %= a;
                break;
            case 6: // Clear Total
                total = 0;
                fr = true; // reset
                break;
            case 7: // Exit
                System.out.println("Exiting the calculator. Goodbye!");
                return;
            default:
                System.out.println("Invalid operator. Please select a valid option.");
                scanner.close();
        }

        System.out.println(); // Print an empty line for better readability
    }
}

public class oop1 {
    void main(String[] args) {
        Calculator.main(args);
    }
}
