#!/usr/bin/python3

def print_last_digit(number):
    last_digit = abs(number) % 10
    print("{}".format(last_digit), end="")
    return last_digit


if __name__ == "__main__":
    print_last_digit(98)
    print_last_digit(0)
    r = print_last_digit(-1024)
    print(r)
