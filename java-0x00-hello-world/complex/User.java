package complex;

public class User {
    static String name;

    public User(String name) {
        User.name = name;
    }

    public static String getName() {
        return User.name;
    }
    
    public static void setName(String newName) { 
        User.name = newName;
    }
}