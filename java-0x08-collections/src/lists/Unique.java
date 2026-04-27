package src.lists;

import java.util.HashSet;
import java.util.Set;

public class Unique {
    private final Set<String> st = new HashSet<>();

    public void add(String name) {
        boolean exists = st.add(name);

        if (!exists) {
            System.out.println("Added user succefully");
            return;
        }
        System.out.println("oops user exist");
    }

    public boolean remove(String name) {
        boolean removed = st.remove(name);
        if (removed) {
            System.out.printf("User %s removed succefully", name);
            return true;
        }
        return false;
    }
    public boolean exist(String name) {
        boolean ok = st.contains(name);
        if (ok) {
            System.out.println("User exist");
            return true;
        }
        return false;
    }

    public int size() {
        return st.size();
    }

    public void union() {
        Set<String> users = new HashSet<>();
        st.addAll(users);
    }
}
