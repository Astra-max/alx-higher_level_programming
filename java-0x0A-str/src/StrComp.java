public class StrComp {
	/**
	 * strinhs are referenced types in java
	 */
	private final String user1 = "james";
	private final String user2 = "john";
	public static void main(String[] args) {
		StrComp str = new StrComp();
		System.out.println("loose equality: " + str.equal());
	}

	public boolean equal() {
		return user1.equals(user2);
	}
}
