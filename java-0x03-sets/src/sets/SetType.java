package sets;

import java.util.*;

public class SetType {
    private Set<Integer> st = new HashSet<>();
    public void basicSet() {
        st.add(9);
        System.out.println(st);
    }

    public void remove(int value) {
        boolean removed = st.remove(value);

        if (removed) {
            System.out.println("rmoved the value"+ value);
        } else {
            System.out.println("value not found");
        }
    }
    public void exist(int value) {
        boolean ok = st.contains(value);

        if (ok) {
            System.out.println("value exist on set");
        } else {
            System.out.println("value not found");
        }
    }
    public boolean duplicates(int[] arr) {
        for (int val : arr) {
            if (!st.add(val)) {
                System.out.println("found duplicate");
                return true;
            }
        }
        return false;
    }
}