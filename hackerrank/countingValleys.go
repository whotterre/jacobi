package main

func countingValleys(steps int32, path string) int32 {
    altitude := 0
    valleyCount := 0
    
    for i := 0; i < len(path); i++ {
        if(path[i] == 'U'){
            altitude += 1       // Ascend if we have a "U"
        } else if (path[i] == 'D'){
            altitude -= 1      // Descend if we get a "D"
        }
        // We have a valley when we're at ground level and have an elevation ahead of us
        if altitude == 0 && path[i] == 'U' {
            valleyCount += 1
        }
    }
    return int32(valleyCount)
}
