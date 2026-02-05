import objects.ModuleCode;

class Main {
    public static void main(String[] args) {
        ModuleCode me = new ModuleCode("astra", 78);
        ModuleCode me2 = me;
        me2.User = "max";
        System.out.println(me.getUser().toUpperCase()+ " is a " + me2.getAge() + " years old.");
    }
}