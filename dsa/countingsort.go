package main

/*
Counting sort from the Data Structures
and Algorithmic Thinking Book by Narasimha
Karumachi

@param A []int - the array to be sorted
@param k int - the maximum
*/
func countingSort(A []int, k int) []int {
	bucketLen := k + 1
	C := make([]int, bucketLen)

	sortedIndex := 0
	length := len(A)

	for i := 0; i < length; i++ {
		C[A[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for C[j] > 0 {
			A[sortedIndex] = j
			sortedIndex += 1
			C[j] -= 1
		}
	}
	return A
}
