package arrayList;

import java.util.ArrayList;

public class DynamicList<T> {
    private final ArrayList<T> userDetails;

    public DynamicList() {
        this.userDetails = new ArrayList<>();
    }
    public void addDetails(T user) {
        this.userDetails.add(user);
    }
    public void showDetails() {
        for (T userData : userDetails) {
            System.out.println(userData.toString());
        }
    }
}
