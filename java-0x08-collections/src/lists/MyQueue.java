package src.lists;

import java.util.LinkedList;
import java.util.Queue;

public class MyQueue {
    private final Queue<String> q = new LinkedList<>();

    public void add(String name) {
        q.add(name);
    }
}
