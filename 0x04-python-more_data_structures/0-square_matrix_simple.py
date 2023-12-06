#!/usr/bin/python3
def square_matrix_simple(matrix=[]):
    # Create a new matrix to store the squared values
    new_matrix = []

    # Iterate through each row in the input matrix
    for row in matrix:
        # Create a new row for the squared values
        new_row = []

        # Iterate each element in the row and
        # append its square to the new row
        for element in row:
            new_row.append(element ** 2)

        # Append the new row to the new matrix
        new_matrix.append(new_row)

    return new_matrix


if __name__ == "__main__":
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]

    new_matrix = square_matrix_simple(matrix)
    print(new_matrix)
    print(matrix)
