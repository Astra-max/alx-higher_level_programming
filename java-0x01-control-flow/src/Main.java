//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        System.out.println("Hello and welcome!");
        conditionIf();
    }
    public static void conditionIf() {
        int myAge = 36;

        if (myAge >= 40) {
            System.out.println("You are an adult");
        } else if (myAge <= 18) {
            System.out.println("You are a teenage");
        } else {
            System.out.println("You are a teenager");
        }
    }
}