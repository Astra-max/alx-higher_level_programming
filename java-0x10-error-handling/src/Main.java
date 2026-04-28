import errors.ThrowingErrors;

public class Main {
    public static void main(String[] arg) {
        ThrowingErrors thrw = new ThrowingErrors();
        System.out.println("Answer == " + thrw.party(13));
    }
}