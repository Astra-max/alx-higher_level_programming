package priorique.libs;

import java.util.Iterator;
import java.util.PriorityQueue;

public class Heap {
    private PriorityQueue<String> pq;
    public Heap(PriorityQueue<String> pq) {
        this.pq = pq;
    }
    public void add(String user) {
        //check if user exist in the list
        if (!pq.contains(user)) pq.add(user);
    }
    public void print() {
        Iterator<String> it = pq.iterator();

        while (it.hasNext()) {
            String user = it.next();
            System.out.printf("user %s: ", user);
        }
    }
}
