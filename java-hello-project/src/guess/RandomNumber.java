package guess;

import java.util.Random;
import utilility.Inputs;

public class RandomNumber {

    private final Inputs inputs = new Inputs();
    private final Random random = new Random();

    private int chooseRange() {
        String range = inputs.getInput("Welcome to the Guessing Game.\nPlease pick a maximum number for the range (the minimum is 1): ");
        try {
            int maxVal = Integer.parseInt(range);
            if (maxVal < 1) {
                System.out.println("Please enter a number greater than or equal to 1.");
                return chooseRange();
            }
            return maxVal;
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return chooseRange();
        }
    }

    private int getRandomNumber(int min, int max) {
        return random.nextInt(max - min + 1) + min;
    }

    private void guessNumber(int min, int max) {
        int randomNumber = getRandomNumber(min, max);

        while (true) {
            String userGuess = inputs.getInput("Guess a number between " + min + " and " + max + ": ");
            int numInput = getIntInput(userGuess);

            if (numInput < randomNumber) {
                System.out.println("Too low! Try again.");
            } else if (numInput > randomNumber) {
                System.out.println("Too high! Try again.");
            } else {
                System.out.println("Congratulations! You've guessed the number!");
                break;
            }
        }
    }

    private int getIntInput(String input) {
        try {
            return Integer.parseInt(input);
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return -1; // Return -1 to indicate an error
        }
    }

    public void play() {
        int max = chooseRange();
        while (true) {
            guessNumber(1, max);
            String playAgain = inputs.getInput("Do you want to play again? (yes/no): ");
            if (!playAgain.equalsIgnoreCase("yes")) {
                System.out.println("Thanks for playing!");
                break;
            }
        }
    }
}
