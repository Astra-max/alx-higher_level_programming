package priorique.libs;

import java.util.PriorityQueue;

public class Heap {
    private PriorityQueue<String> pq;
    Heap(PriorityQueue<String> pq) {
        this.pq = pq;
    }
    public void add(String user) {
        //check if user exist in the list
        if (!pq.contains(user)) pq.add(user);
    }
}
