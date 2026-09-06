
from remove import nums

def count(arr, val):
    occurence = 0

    for n in arr:
        if n == val:
            occurence += 1

    print(occurence)


count(nums, 3)
