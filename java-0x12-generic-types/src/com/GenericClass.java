package com;

/**
 * GenericClass - is generic class
 * enable code reusability
 */

public class GenericClass<T> implements GenericInterface<T> {
    private  T value;
    public GenericClass(T value) {
        this.value = value;
    }

    @Override
    public T getValue() {
        return value;
    }
    @Override
    public void setValue(T value) {
        this.value = value;
    }
}