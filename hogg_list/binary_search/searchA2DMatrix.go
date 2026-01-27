package main
// Slightly optimal brute force method 
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false 
    }
    for i := 0; i < len(matrix); i++ {
        if binSearchHelper(matrix[i], target) == false {
            continue
        } else {
            return true
        }
    }
    return false
}

func binSearchHelper(arr []int, target int) bool {
    n := len(arr)
    low := 0
    high := n - 1

    for low <= high {
        mid := low + (high - low) / 2

        if arr[mid] == target {
            return true
        }
        if arr[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
      return false
}