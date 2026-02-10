package objects;

public class ModuleCode {
    public String User;
    int Age;
    public ModuleCode(String User, int age) {
        this.User = User;
        this.Age = age;
    }

    public String getUser() {
        return this.User;
    }

    public int getAge() {
        return this.Age;
    }
}