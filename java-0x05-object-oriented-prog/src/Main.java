import car.Car;
import car.constructor.CarConstructor;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        Car car = new Car();
        CarConstructor c = new CarConstructor("Toyota", "green", 2024);
        c.getDetails();
        car.drive();

        Anonymous anon = new Anonymous();
        anon.print();
    }
}