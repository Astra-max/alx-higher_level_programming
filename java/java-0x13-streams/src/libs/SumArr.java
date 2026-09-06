import java.util.Arrays;

public class SumArr {
	public  void sum() {
		int[] numbers = {1,2,3,4,5};
		int sum = Arrays.stream(numbers)
			.reduce(0, (a, b) -> a + b);

				            System.out.println(sum);
					        }
}
