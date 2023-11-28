#!/usr/bin/python3

def numbers(i):
    for number in range(i):
        if number < 99:
            print("{:02d},".format(number), end=" ")
        else:
            print("{:02d}".format(number))


numbers(100)
