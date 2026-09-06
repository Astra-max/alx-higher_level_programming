public class ToTittle {
	private final String name = "astra";

	public static void main(String[] args) {}
	public String totitle() {
		char[] charArr = name.toCharArray();
		if (charArr.length == 0) return "imvalid";

		char first = charArr[0];

		if (first >= 'a' && first <= 'z') first -= 32;
		char[] newArr = new char[]{};
		newArr.push(first);

		for (int i = 1; i<charArr.length-1; i++) {
			newArr.push(charArr[i]);
		}
		return new String(newArr);
	}
}
