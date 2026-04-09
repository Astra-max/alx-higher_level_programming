package complex;

public class Referenced {
    public static String strings(User user) {
        User.setName("Astra");
        return User.getName();
    }
}
