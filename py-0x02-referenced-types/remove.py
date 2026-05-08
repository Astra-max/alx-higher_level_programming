def remove(nums, val):
    results = []

    for n in nums:
        if n == val:
            continue
        else:
            results.append(n)

    print(results)

nums = [3]
remove(nums, 2)
