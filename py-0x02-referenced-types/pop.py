from remove import nums


def pop(arr):
    if len(arr) == 0:
        return 0
    last = arr[len(arr)-1]
    arr = arr[:-1]
    print(arr)
    return last

print(pop(nums))
