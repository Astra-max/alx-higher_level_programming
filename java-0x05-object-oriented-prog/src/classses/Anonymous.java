/**
 * anonymous class - are classes that are defined and executed immediately
 * used to override methods from a class
 */

public class Anonymous {
    public String getUser() {
        return "astra6";
    }
    public Abstraction shadow = new Abstraction() {
        @Override
        public String getUser() {
            return "astra7";
        }
    };
    public void print() {
        System.out.println(shadow.getUser());
    }
}