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

	public boolean addAtIdx(int idx, String friend) {
		return friends.add(1, friend);
	}

	public boolean remove(String fakeFriends) {
		return friends.remove(fakeFriends);
	}
	public void print() {
		for (String friend: friends) {
			System.out.printf("%s is amazing friend of mine\n", friend);
		}
	}
	public boolean exists(String fakeFriend) {
		return friends.contains(fakeFriends);
	}
}
