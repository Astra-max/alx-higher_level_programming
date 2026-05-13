package priorique;

import java.util.PriorityQueue;
import priorique.libs.Heap;

public class Main {
 public static void main(String[] args) {
    PriorityQueue<String> pq = new PriorityQueue<>();
    Heap<String> heap = new Heap<>(pq);
    heap.add("max");
    heap.add("max");
    heap.print();
 }   
}
