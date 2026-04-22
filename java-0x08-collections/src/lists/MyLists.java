package src.lists;

import java.util.ArrayList;
import java.util.List;

public class MyLists {
    private final List<String> myList = new ArrayList<>();
    private String user;

    public void add(String user) {
        myList.add(user);
    }

    public void print() {
        for (int i = 0; i < myList.size(); i++) {
            System.out.printf("UserName: %s\n", myList.get(i));
        }
    }

    public void remove(int itemId) {

        for (int i = 0; i < myList.size(); i++) {
            if (i == itemId) {
                try {
                    user = myList.get(itemId);
                    myList.remove(itemId);
                } catch (IndexOutOfBoundsException e) {
                    System.out.printf("Failed to remove item: %s\n", e.getMessage());
                    return;
                }
                finally {
                    System.out.printf("removed %s succefully\n", user);
                }
            }
        }
    }
}
