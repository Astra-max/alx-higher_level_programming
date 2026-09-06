package encapsulation;

public class Encapsulation {
    private String name;
    private String color;

    Encapsulation(String name, String color) {
        this.name = name;
        this.color = color;
    }

    public String getName() {
        return this.name;
    }
    public String getColor() {
        return this.color;
    }
    public void setName(String name) {
        this.name = name;
    }
    public void setColor(String color) {
        this.color = color;
    }
}
