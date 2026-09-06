from remove import nums

def insert(arr, idx, val):
    results = []

    if idx > len(arr):
        print("Index out of bound")
        return

    startIdx = 0

    while startIdx < len(arr):
        if startIdx == idx:
            results.append(val)
            results.append(arr[startIdx])
        else:
            results.append(arr[startIdx])

        startIdx += 1

    print(results)

insert(nums,2,89)
