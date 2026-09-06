from remove import nums

arr = [n*2 for n in nums]
ne = [n*m for n in nums for m in nums]
print(arr, ne)
