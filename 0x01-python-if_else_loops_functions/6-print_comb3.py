#!/usr/bin/python3

for i in range(10):
    for j in range(i + 1, 10):
        if i != j:
            print("{:02d}".format(i * 10 + j), end=", " if i * 10 + j != 89 else "\n")

