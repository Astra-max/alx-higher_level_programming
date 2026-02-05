import objects.ModuleCode;

class Main {
    public static void main(String[] args) {
        ModuleCode me = new ModuleCode("astra", 78);
        System.out.println(me.getUser().toUpperCase()+ " is a " + me.getAge() + " years old.");
    }
}