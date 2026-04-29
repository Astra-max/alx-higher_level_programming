import com.GenericClass;

public class Main {
    public static void main(String[] arg) {
        GenericClass<String> gcc = new GenericClass<String>("astra");
        System.out.println("user name: " + gcc.getName());
    }
}
