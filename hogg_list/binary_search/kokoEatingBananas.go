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
	//    Optimize: with binary search

	upper := getMax(piles)
	for k := 1; k <= upper; k++ {
		currentCost := 0
		for i := range piles {
			currentCost += (piles[i] + k - 1) / k
		}

		if currentCost <= h {
			return k
		}
	}
	return upper
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
