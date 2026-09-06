/*
 * Cow - inherits fro animal class
 * sound - overrides parent class sound
 * return - void
 */

public class Cow  extends Animal {
    @Override
    public void sound() {
        System.out.println("cow moows");
    }
}
