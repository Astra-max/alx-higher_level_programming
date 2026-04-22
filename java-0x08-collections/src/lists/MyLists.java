package src.lists;

import java.util.ArrayList;
import java.util.List;

public class MyLists {
    List<String> myList = new ArrayList<>();
    public void add(String user) {
        myList.add(user);
    }
    public void print() {
        for (int i = 0; i< myList.size(); i++) {
            System.out.printf("UserName: %s\n", myList.get(i));
        }
    }
}
