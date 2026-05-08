#!/bin/python3

friends = list()

friends.append("astra")
friends.append("judy")
friends.insert(1,"max")
counts = friends.count("judy")


for friend in friends:
    print(friend)


print(f"judy exists {counts} times")
