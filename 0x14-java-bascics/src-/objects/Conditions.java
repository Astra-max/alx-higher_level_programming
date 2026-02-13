package objects;

public class  Conditions {
    int Num;

    public String Msgs() {
        String msgs = this.Num < 78 ? "less value" : "more or equal";
        return msgs;
    }
}