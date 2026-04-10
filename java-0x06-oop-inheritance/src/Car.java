public class Car extends Vehicle {
    String name;
    String color;
    Car(String name, String color) {
        this.name = name;
        this.color = color;
    }
    public void move() {
        System.out.println("the car is moving");
    }
}
