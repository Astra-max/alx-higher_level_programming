package priorique.libs;

import java.util.Iterator;
import java.util.PriorityQueue;

public class Heap<T> {
    private PriorityQueue<T> pq;

    public Heap() {
        this.pq = new PriorityQueue<>();
    }

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
    public T getValue(T val) {
        if (pq.contains(val)) return val;
        return null;
    }
    public Heap<T> setValue(T from, T to) {
        if (!pq.contains((from))) return this;

        Heap<T> newHeap = new Heap<>();
        Iterator<T> it = pq.iterator();

        while (it.hasNext()) {
            T toChange = it.next();
            if (toChange.equals(from)) {
                newHeap.add(to);
            } else {
                newHeap.add(toChange);
            }
        }
        return newHeap;
    }
    public boolean removeVal(T value) {
        if (this.pq.contains(value)) return this.pq.remove(value);
        return false;
    }

    public boolean removeAll(T value) {
        if (this.pq.contains(value)) return this.pq.removeAll(this.pq);
        return false;
    }
}
