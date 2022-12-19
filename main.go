package main

import "fmt"

/*
   Question 2: Given two non-empty arrays of integers, write a function that determines if the second
   array is a subarray of the first one.
*/
/*
Solution for findng if a sequence is a subarray of an array
*/
func isSubarray(A, B []int) bool {
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
func isSubsequence(A, B []int) bool {
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
	arr1 := []int{1, -4, 3, 5, 10, -6}
	arr2 := []int{-4, 3, 5}
	arr3 := []int{-4, 1, 5}
	arr4 := []int{1, 5, 10}

	fmt.Println(isSubarray(arr1, arr2))
	fmt.Println(isSubarray(arr1, arr3))
	fmt.Println(isSubarray(arr1, arr4))

	fmt.Println(isSubsequence(arr1, arr2))
	fmt.Println(isSubsequence(arr1, arr4))
	fmt.Println(isSubsequence(arr1, arr3))
}
