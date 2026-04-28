package enums;

public enum MyEnums {
    BLUE("blue"), RED("red"), GREEN("green");
    private String Value;

    MyEnums(String value) {
        this.Value = value;
    }

    public String getValue() {
        return this.Value;
    }
}
