#!/usr/bin/python3

for i in range(10):
    for j in range(i + 1, 10):
        if i != j:
            end_str = ", " if i * 10 + j != 89 else "\n"
            print("{:02d}".format(i * 10 + j), end=end_str)
