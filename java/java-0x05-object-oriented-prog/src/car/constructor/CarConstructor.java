package car.constructor;

public class CarConstructor {
    String name;
    String color;
    int manufactureDate;

   public  CarConstructor(String name, String color, int date) {
        this.color = color;
        this.name = name;
        this.manufactureDate = date;
    }
    public void getDetails() {
        System.out.printf("Brand new %s %s\n",this.color, this.name);
    }
}
