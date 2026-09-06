def remove(nums, val):
    results = []

    for n in nums:
        if n == val:
            continue
        else:
            results.append(n)

    print(results)

nums = [1,6,2,3,3,11,9,4,5,6,6,7,9,10]
remove(nums, 2)
