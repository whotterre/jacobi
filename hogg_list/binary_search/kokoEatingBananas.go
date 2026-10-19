package main

// Brute force approach
func minEatingSpeed(piles []int, h int) int {
	//    initialize a var for cost, at 0
	//    iterate through the interval, [1, max(p)], with var k,
	//    iterate through the piles array, trying k in piles[i],
	//    currentCost += ceil(piles[i] / k)
	//    at the end of each iter, cost = min(cost, currentCost)
	//    return cost
	//    O(max(p) * len(p)) = infeasible for these constraints
	//    Optimize: with binary search over the interval [1, max(p)] reducing the TC to O(log(p) * len(p))

	upper := getMax(piles)
	l, r := 1, upper
	ans := upper 
	for l <= r {
		k := l + (r - l) / 2
		currentCost := 0
		for i := range piles {
			currentCost += ceil(piles[i], k)
		}

		if currentCost <= h {
			ans = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return ans
}

func getMax(arr []int) int {
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// For cleaner code, and understandability, this computes ceiling divisor of a by b
func ceil(a, b int) int {
    return (a + b - 1) / b
}