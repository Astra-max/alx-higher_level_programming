public class StrComp {
	/**
	 * strinhs are referenced types in java
	 */
	private final String user1 = "John";
	private final String user2 = "john";
	public static void main(String[] args) {
		StrComp str = new StrComp();
		System.out.println("strict equality: " + str.equal());
		System.out.println("loose equality check: "+ str.ignoreCase());
	}

	public boolean equal() {
		return user1.equals(user2);
	}

	public boolean ignoreCase() {
		return user1.equalsIgnoreCase(user2);
	}
}
