package com;

public class GenericClass<T> {
    private final T userName;
    public GenericClass(T user) {
        this.userName = user;
    }
    public T getName() {
        return userName;
    }
}