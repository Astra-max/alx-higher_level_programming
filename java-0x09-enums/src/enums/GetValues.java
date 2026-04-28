package enums;

public class GetValues {
    public void getAll() {
        System.out.println("Please select from available colors");
        for (MyEnums en : MyEnums.values()) {
            System.out.println("color --> " + en.getValue());
        }
    }
}
