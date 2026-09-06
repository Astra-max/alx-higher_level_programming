from remove import nums


def extend(list1,list2):
    for n in list2:
        list1.append(n)

    print(list1)

extend(nums, [2,3,4,5])
