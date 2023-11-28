#!/usr/bin/python3

def reverse_hex():
    for i in range(122, 96, -1):
        print("{}".format(chr(i) if i % 2 == 0 else chr(i - 32)), end="")


reverse_hex()
