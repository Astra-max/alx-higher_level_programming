#!/usr/bin/python3

def remove_char_at(s, n):
    if n < 0 or n >= len(s):
        return s

    result = ""
    for i in range(len(s)):
        if i != n:
            result += s[i]

    return result


if __name__ == "__main__":
    print(remove_char_at("Best School", 3))
    print(remove_char_at("Chicago", 2))
    print(remove_char_at("C is fun!", 0))
    print(remove_char_at("School", 10))
    print(remove_char_at("Python", -2))
