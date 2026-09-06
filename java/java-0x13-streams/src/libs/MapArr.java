import java.util.Arrays;

public class MapArr {

    public void byTwo() {

        int[] numbers = {1,2,3,4,5};

        int[] doubled = Arrays.stream(numbers)
                .map(n -> n * 2)
                .toArray();

        System.out.println(Arrays.toString(doubled));
    }
}
