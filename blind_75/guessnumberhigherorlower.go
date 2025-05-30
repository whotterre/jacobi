package main

// guess is a placeholder function for the actual guess API.
// In a real scenario, this function would be provided by the system.
// It returns:

func guessNumber(n int) int { 
    l, r := 1, n
    for l <= r {
        mid := l + (r-l)/2  
        res := guess(mid)   
        switch {
        case res == 0:
            return mid
        case res == -1:
            r = mid - 1
        case res == 1:
            l = mid + 1
        }
    }
    return -1  
}

func guess(mid int) int {
     // Placeholder for the guess API, just occupying the space
     // The real one is given by the system
	return mid - 42
}
