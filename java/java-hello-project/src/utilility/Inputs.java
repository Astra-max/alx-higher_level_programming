package utilility;

import java.util.Scanner;


public class Inputs { 
    private static final Scanner scanner = new Scanner(System.in);
    
    public String getInput(String prompt) {
        System.out.print(prompt);
        String userInput;

        try {
            userInput = scanner.nextLine();
        } catch (Exception e) {
            System.out.println("An error occurred while reading input: " + e.getMessage());
            userInput = "";
        }
        return userInput;
    }
}
