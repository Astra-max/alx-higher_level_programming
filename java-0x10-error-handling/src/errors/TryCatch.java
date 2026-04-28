package errors;


import java.util.Scanner;


public class TryCatch {
    private final Scanner scan = new Scanner(System.in);
    int dividend;

    private int getNum(String message) {
        System.out.print(message);
        
        try {
            return scan.nextInt();
        } catch(Exception e) {
            System.out.println(e);
            return 0;
        }
    }
    public int divide() {
        try {
            int num1 = getNum("First value: ");
            int num2 = getNum("second value: ");
            dividend = num1/num2;
        } catch (ArithmeticException e) {
            System.out.printf("Got error %s\n", e.getMessage());
        }
        return dividend;
    }
}