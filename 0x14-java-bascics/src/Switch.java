public class Switch {
    public static void main(String[] args) {
        String grade = "B";
        String msg;

        switch (grade) {
            case "A":
                msg = "Excellent!";
                break;
            case "B":
                msg = "Good job!";
                break;
            case "C":
                msg = "You passed!";
                break;
            case "D":
                msg = "You barely passed!";
                break;
            default:
                msg = "Invalid grade";
        }

        System.out.println(msg + " Your grade is: " + grade);
    }
}