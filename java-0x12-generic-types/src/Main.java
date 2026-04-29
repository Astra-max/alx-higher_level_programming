import com.GenericClass;

public class Main {
    public static void main(String[] arg) {
        GenericClass<Integer> gcc = new GenericClass<>(89);
        System.out.println("user name: " + gcc.getValue());
    }
}
