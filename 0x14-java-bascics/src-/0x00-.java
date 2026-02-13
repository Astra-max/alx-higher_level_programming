import objects.GradingSys;

class Main {
    public static void main(String[] args) {
        GradingSys me = new GradingSys("astra", 78, 'E');
        System.out.println(me.getUser().toUpperCase()+ " you " + me.Message());
    }
}