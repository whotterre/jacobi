package main

func findClosestNumber(nums []int) int {
    closest := nums[0]

    for _, num := range nums {
        if abs(num) < abs(closest) {
            closest = num
        }

        // tie breaker event 
        if abs(num) == abs(closest) {
            closest = max(num, closest)
        }
    }
    return closest
}

func abs(a int) int {
    if a < 0 {
        return -a 
    } 
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}