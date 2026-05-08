from remove import nums

def insert_sort(arr):
    first = 0
    second = 1
    for n in arr:
        while second >= 0 and first < len(arr) and second < len(arr):
            if arr[first] > arr[second]:
                arr[first],arr[second] = arr[second],arr[first]
                second -= 1
            else:
                second += 1
        first += 1
        

    return arr

print(insert_sort(nums))
