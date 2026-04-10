package complex;

public class User {
    String name;

    public User(String name) {
        this.name = name;
    }

    public String getName() {
        return this.name;
    }
    
    public void setName(String newName) { 
        this.name = newName;
    }
}