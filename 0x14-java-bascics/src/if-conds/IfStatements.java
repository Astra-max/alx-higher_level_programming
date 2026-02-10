public class IfStatements {
    String userName;
    String grade;
   IfStatements(String userName, String grade) {
        this.userName = userName;
        this.grade = grade;
   }

        if (this.grade.equals("A")) {
            grade = "A";
        } else if (this.grade.equals("B")) {
            grade = "B";
        } else if (this.grade.equals("C")) {
            grade = "C";
        } else if (this.grade.equals("D")) {
            grade = "D";    
        } else {
            grade = "F";
        }

        System.out.println("Your grade is: " + grade);
    }