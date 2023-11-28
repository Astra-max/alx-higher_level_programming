#!/usr/bin/python3

def fizzbuzz():
    for game in range(1, 101):
        if game % 3 == 0 and game % 5 == 0:
           print("FizzBuzz", end=" ")
        elif game % 3 == 0:
           print("Fizz")
        elif game % 5 == 0 and game % 5 == 0:
           print("Buzz", end=" ")
        else:
            print("{}".format(game), end=" ")


if __name__ == "__main__":
             fizzbuzz()
             print("")
