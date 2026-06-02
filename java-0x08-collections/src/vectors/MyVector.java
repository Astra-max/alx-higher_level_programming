package vectors;


public class MyVector {
	private final Vector<String> friends = new Vector<>();
	public int size() {
		return friends.size();
	}

	public boolean add(String friend) {
		friends.add(friend);
		return true;
	}

	public boolean remove(String fakeFriends) {
		return friends.remove(fakeFriends);
	}
}
