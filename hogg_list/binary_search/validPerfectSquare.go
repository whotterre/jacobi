package main

func isPerfectSquare(num int) bool {

    if num < 2 {
        return true
    }

    low := 1
    high := num

    for low <= high {
        mid := low + (high-low)/2
        if mid * mid == num {
            return true
        }

        if mid * mid < num {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}