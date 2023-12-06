#!/usr/bin/python3

def common_elements(set_1, set_2):
    # returns common element in sets

    common_item = set_1 & set_2
    if common_item:
        return list(common_item)
    else:
        return False


if __name__ == '__main__':
    set_1 = {"Python", "C", "Javascript"}
    set_2 = {"Bash", "C", "Ruby", "Perl"}
    c_set = common_elements(set_1, set_2)
    print(sorted(list(c_set)))
