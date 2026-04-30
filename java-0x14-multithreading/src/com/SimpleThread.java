package com;

public class SimpleThread extends Thread {
    @Override

    public void run() {
        for (int i = 0; i<9; i++) {
            System.out.println("Inside Simple thread" + i);
        }
    }
}
