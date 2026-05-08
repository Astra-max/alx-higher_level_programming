users = [1,2,3,4,5]
# users.reverse()

def reverse_list(users):
    results = []

    start, end = 0, len(users)-1
    while end >= start:
        results.append(users[end])
        end -= 1

    print(results)


reverse_list(users)
