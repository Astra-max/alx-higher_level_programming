def remove(nums, val):
    results = []

    for n in nums:
        if n == val:
            continue
        else:
            results.append(n)

    print(results)

nums = [1,2,3,3,4,5,3]
remove(nums, 2)
