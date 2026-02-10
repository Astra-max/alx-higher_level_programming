public class While {
    public static void main(String[] args) {
        AllElements();
    }
    public AllElements() {
        String[] name = {"astra", "sam", "evans", "june"};
        int size = name.length()-1;
        while (size => 0) {
            System.out.println(name[size]);
            size--;
        }
    }
}