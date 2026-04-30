import com.SimpleThread;

public class Main {
    public static void main(String[] arg) {
        System.out.println("main starts");
        Thread st = new SimpleThread();
        st.start();
        System.out.println("main ends");
    }
}