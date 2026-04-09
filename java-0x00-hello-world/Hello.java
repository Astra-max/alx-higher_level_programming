import complex.Referenced;
import complex.User;
import primitive.PrimitiveVar;

public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello, world!");
        PrimitiveVar.primVariables();
        String newName = Referenced.strings(new User("james"));
        System.out.println("My new name is: " + newName);
    }
}
