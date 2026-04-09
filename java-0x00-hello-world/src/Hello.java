import complex.Referenced;
import complex.User;
import primitive.PrimitiveVar;

public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello, world!");
        PrimitiveVar.primVariables();

        Referenced ref = new Referenced();

        String newName = ref.strings(new User("Astra6"));
        System.out.println("My new name is: " + newName);
    }
}