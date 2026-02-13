package main

import "strconv"
func isBalanced(num string) bool {
    evenSum, oddSum := 0, 0
    for i := 0; i < len(num); i++ {
        if i % 2 == 0 {
            res, _ := strconv.Atoi(string(num[i]))
            evenSum += res
        } else {
            res, _ := strconv.Atoi(string(num[i]))
            oddSum += res
        }
    }
    return evenSum == oddSum
}