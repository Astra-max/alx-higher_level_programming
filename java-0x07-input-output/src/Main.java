import files.Basics;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        String filename = "astra.txt";
        Main app = new Main();
        app.control(filename);
    }

    public void control(String fileName) {
        Basics ioFileOper = new Basics(fileName);

        if (!ioFileOper.checkFile()) {
            ioFileOper.create();
            ioFileOper.write("My dummy data\n");
            ioFileOper.read();
        } else {
            ioFileOper.write("Another data added\n");
            ioFileOper.read();
        }
    }
}