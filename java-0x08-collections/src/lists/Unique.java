package src.lists;

import java.util.HashSet;
import java.util.Set;

public class Unique {
    private final Set<String> st = new HashSet<>();

    public void add(String name) {
        st.add(name);
    }
    public boolean remove(String name) {
        boolean removed = st.remove(name);
        if (removed) {
            System.out.printf("User %s removed succefully", name);
            return true;
        }
        return false;
    }
}
