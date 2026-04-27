package src.lists;

import java.util.LinkedList;
import java.util.Queue;

public class MyQueue {
    private final Queue<String> q = new LinkedList<>();

    public void add(String name) {
        q.add(name);
    }

    public boolean remove(String name) {
        try {
            boolean done = q.remove(name);
            if (done) {
                return true;
            }
        } catch (Exception e) {
            System.out.println("Failed to removed user");
        }
        return false;
    }

    public String top() {
        try {
            String topVal = q.peek();

            if (!topVal.equals("")) {
                return topVal;
            }

        } catch (Exception e) {
            System.out.println("Error finding value at the top");
        }
        return "";
    }
}
