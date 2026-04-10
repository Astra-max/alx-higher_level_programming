package arrayList;

public class User {
    String name;
    int age;
    public User(String user, int age) {
        this.name = user;
        this.age = age;
    }
    public  String getName() {
        return this.name;
    }
    public int getAge() {
        return this.age;
    }
    @Override
    public String toString() {
        return String.format("Username: %s userAge: %d",name, age);
    }
}