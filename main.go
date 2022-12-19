package main

/*
   Question 2: Given two non-empty arrays of integers, write a function that determines if the second
   array is a subarray of the first one.
*/
/*
Solution for findng if a sequence is a subarray of an array
*/
func isSubarray[T any](A, B []int) bool {
	aLen := len(A)
	bLen := len(B)
	i, j := 0, 0

	for i < aLen && j < bLen {
		if A[i] == B[j] {
			i += 1
			j += 1

			if j == bLen {
				return true
			}
		} else {
			i = i - j + 1
			j = 0
		}
	}
	return false
}

/*
   Question 2: Given two non-empty arrays of integers, write a function that determines if the second
   array is a subsequence of the first one.
*/
/*
Solution for findng if a sequence is a subsequence of an array
*/
func isSubsequence[T any](A, B []int) bool {
	aLen := len(A)
	bLen := len(B)
	i, j := 0, 0

	for i < aLen && j < bLen {
		if A[i] == B[j] {
			i += 1
			j += 1

			if j == bLen {
				return true
			}
		} else {
			i += 1
		}
	}
	return false
}

func main() {

}
