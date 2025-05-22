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
