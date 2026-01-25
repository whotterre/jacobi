package main

import (
    "sort"
    "math"
)
func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)

    if len(nums) == 1 {
        return 0
    }   
    minDiff := math.MaxInt32
    // Set a pointer at 0
    // set a pointer k - 1
   for i := 0; i + k - 1 < len(nums); i++ {
        j := i + k - 1
        // take max 
        maxEl := nums[j]
        minEl := nums[i]
        diff := maxEl - minEl
        if diff < minDiff {
           minDiff = diff
        }
        j += 1
    }
    return minDiff 
}
