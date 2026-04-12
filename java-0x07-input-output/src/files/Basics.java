package files;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Scanner;

public class Basics {
    private final String myFile;
    private final File file;
    public Basics(String fileName) {
        this.myFile = fileName;
        this.file = new File(fileName);
    }

    public boolean checkFile() {
        return file.exists();
    }
    public File create() {
        try {
            if (file.createNewFile()) {
                return this.file;
            } else {
                System.out.printf("File %s exists", this.myFile);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        return this.file;
    }
    public void write(String data) {
        try {
            FileWriter writer = new FileWriter(this.file,true);
            writer.write(data);
            System.out.println("[#]: data saved successfully!!!");
            writer.close();
        } catch(IOException e) {
            System.out.printf("error: %s\n", e.getMessage());
        }
    }
    public void read() {
        try {
            Scanner reader = new Scanner(this.file);
            while (reader.hasNextLine()) {
                String data = reader.nextLine();
                System.out.println(data);
            }
            reader.close();
        } catch(FileNotFoundException e) {
            System.out.printf("file not found: %s\n", e.getMessage());
        }
    }
}
