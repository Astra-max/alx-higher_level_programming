#!/bin/python3


def iterative_fact(num: int):
    fact = num
    for i in range(1, num):
        fact = fact * i


    return fact

print(iterative_fact(5))
