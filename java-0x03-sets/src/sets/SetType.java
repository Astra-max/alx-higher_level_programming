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
}