public class Fish implements Prey,Predator{
    public  void hunt() {
        System.out.println("this fish is hunting smaller fish");
    }

    @Override
    public void flee() {
        System.out.println("This fish is hunted by bigger fish");
    }
}
