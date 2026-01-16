package main

func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	low := 1
	high := n - 1
	badCandidate := n
	for low <= high {
		mid := low + (high - low) / 2

		if isBadVersion(mid){ // check the left portion if there are more
			badCandidate = mid
			high = mid - 1 
		} else {
			low = mid + 1
		}
	}
	return badCandidate
}