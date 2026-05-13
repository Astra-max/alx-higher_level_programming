package priorique.libs;

import java.util.Iterator;
import java.util.PriorityQueue;

public class Heap<T> {
    private PriorityQueue<T> pq;
    public Heap(PriorityQueue<T> pq) {
        this.pq = pq;
    }
    public void add(T user) {
        //check if user exist in the list
        if (!pq.contains(user)) pq.add(user);
        else {
            System.out.printf("User %s exists\n", user);
        }
    }
    public void print() {
        Iterator<T> it = pq.iterator();

        while (it.hasNext()) {
            T user = it.next();
            System.out.printf("user :%s \n", user);
        }
    }
}
