#!/usr/bin/python3

for number in range(100):
    if number < 10:
        print("0{},{}".format(number, " "), end="")
    elif number == 99:
        print("{}".format(number))
    else:
        print("{},{}".format(number, " "), end="")
