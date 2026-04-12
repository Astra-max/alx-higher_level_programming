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
        control(userInput);
    }
    public static int input(String message) {
        System.out.print(message);
        Scanner scanner = new Scanner(System.in);
        return scanner.nextInt();
    }
    public static void control(int userInput) {
        Animal animal;
        switch (userInput) {
            case 1:
                animal = new Cat();
                animal.sound();
                break;
            case 2:
                animal = new Cow();
                animal.sound();
                break;
            default:
                animal = new Animal();
                animal.sound();
        }
    }
}