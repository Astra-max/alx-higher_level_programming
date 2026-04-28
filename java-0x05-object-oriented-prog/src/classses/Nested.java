/**
 * Nested class - has nested class Innerclass
 * nested class help group related code together
 * to execute inner class you create out class obj first
 * them inner class object
 */

public class Nested {
    private String name = "astra6";

    public String OutterName() {
        return this.name;
    }
    public class InnerClass {
        private String name = "james";

        public String InnerName() {
            return this.name;
        }
    }
}