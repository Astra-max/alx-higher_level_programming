package errors;

public class ThrowingErrors {
    public String party(int age) {
        if (age < 18) {
            throw new CustomeErrors("Sorrry you are a minor");
        }
        return null;
    }
}
