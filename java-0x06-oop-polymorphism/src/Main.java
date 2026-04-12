import java.util.Scanner;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.

/**
 * Mian - program entry point
 * input - takes user input
 * control - executes user selected value
 * return - void
 */
public class Main {
    public static void main(String[] args) {
        int userInput = input("Select animal (1=cat, 2=cow): ");
        Animal animal = control(userInput);
        animal.sound();
    }
    public static int input(String message) {
        System.out.print(message);
        Scanner scanner = new Scanner(System.in);
        return scanner.nextInt();
    }
    public static Animal control(int userInput) {
        return switch (userInput) {
            case 1 -> new Cat();
            case 2 -> new Cow();
            default -> new Animal();
        };
    }
}