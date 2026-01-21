package main

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    res := make([]int, n)
    var stack []int

    for i := 0; i < n; i++ {
        for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {
            index := stack[len(stack) - 1]
            stack = stack[: len(stack) - 1]
            res[index] = i - index
        }
        stack = append(stack, i)
    }
    return res
}