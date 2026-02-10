public class ForILoop {
    static void main(String[] args) {
        AllElements();
    }

    void AllElements() {
        String[] name = {"astra", "maxwel", "june", "jane"};

        for (int i =0; i<name.length; i++) {
            System.out.println(name[i]);
        }
    }
}