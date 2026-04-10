import arrayList.DynamicList;
import arrayList.User;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        DynamicList<User> dl = new DynamicList<>();
        dl.addDetails(new User("maxwel", 89));
        dl.showDetails();
    }

    public static void forLoopArrStr() {
        String[] userNames = {"James", "Adede", "Alphonse"};

        for (String user : userNames) {
            System.out.println(user + "is our member");
        }
    }
    public static void forLoopArrInt() {
        int[] userAges = {23,44,56};

        for (int i = 0; i<userAges.length; i++) {
            System.out.println(i + " user: "+ "age: "+ userAges[i]);
        }
    }
}