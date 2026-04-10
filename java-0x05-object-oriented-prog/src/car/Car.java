package car;

public class Car {
    String name = "mercedes";
    String model;
    int manufacture;
    String color = "white";

    public void drive() {
        System.out.printf("driving %s %s\n", this.color, this.name);
    }
    public void stop() {
        System.out.printf("stopping %s %s\n", this.color, this.name);
    }
}
