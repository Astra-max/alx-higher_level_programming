package hashMap;


public class Maps {
    private final HashMap<String, Integer> map = new HashMap<>();
    public void add(int value) {
        map.put("name", value)
    }
    public void print() {
        for (String key: map) {
            System.out.println("Got: ", map.get(key))
        }
    }
}