package objects;

class ModuleCode {
    String User;
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


public class GradingSys extends  ModuleCode{
    char Grade;

    public GradingSys(String User, int Age, char grade) {
        super(User, Age);
        this.Grade = grade;
    }

    public String Message() {
        return switch (this.Grade) {
            case 'A' -> "Excellent !";
            default -> "You failed";
        }; 
    }
}