import src.lists.MyLists;

public class Main {
    public static void main(String[] args) {
        MyLists mylist = new MyLists();
        mylist.add("astra");
        mylist.add("james");
        mylist.add("john");
        mylist.remove(2);
        mylist.print();
    }
}
