/**
 * Cat - inherits from animal
 * sound - overrides parents sound
 * return - void
 */

public class Cat extends Animal {
    @Override
    public void sound() {
        System.out.println("cat meow!!");
    }
}
