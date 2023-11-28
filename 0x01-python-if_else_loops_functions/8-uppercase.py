#!/usr/bin/python3

def uppercase(s):
    for char in s:
        uppercase_char = chr(ord(char) - 32) if 'a' <= char <= 'z' else char
        print("{}".format(uppercase_char), end="")
    print()
