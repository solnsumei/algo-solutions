"""
    Question 1: Given two non-empty arrays of integers, write a function that determines if the second
    array is a subarray of the first one.
"""
"""
Solution for findng if a sequence is a subarray of an array
"""
def is_subarray(A, B):
    n = len(A)
    m = len(B)
    i = j = 0

    while i < n and j < m:
        if A[i] == B[j]:
            i += 1
            j += 1

            if j == m:
                return True
        else:
            i = i - j + 1
            j = 0

    return False


"""
    Question 2: Given two non-empty arrays of integers, write a function that determines if the second
    array is a subsequence of the first one.
"""
"""
Solution for findng if a sequence is a subsequence of an array
"""
def is_subsequence(A, B):
    n = len(A)
    m = len(B)
    i = j = 0

    while i < n and j < m:
        if A[i] == B[j]:
            i += 1
            j += 1

            if j == m:
                return True
        else:
            i += 1

    return False



if __name__ == "__main__":
    a = [1, 3, -6, 3, 10, 12]
    b = [3, 3]

    print(is_subarray(a, b))
    print(is_subsequence(a, b))
