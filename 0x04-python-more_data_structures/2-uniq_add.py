#!/usr/bin/python3

def uniq_add(my_list=[]):
    unique_numb = set()

    for num in my_list:
        unique_numb.add(num)

    result = sum(unique_numb)

    return result


if __name__ == '__main__':
    my_list = [1, 2, 3, 1, 4, 2, 5]
    result = uniq_add(my_list)
    print("Result: {:d}".format(result))
