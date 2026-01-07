package main

func romanToInt(s string) int {
    res := 0
    // full num
    valueMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    lastValue := 0
    for i := 0; i < len(s); i++ {
        // count occurrences of Roman numeral chars
        currValue := valueMap[s[i]]
        if lastValue == 0 {
            res += currValue
        } else {
            if lastValue > currValue || lastValue == currValue {
               res += currValue
        } else {
           res += currValue - (2 * lastValue)
          }
        }
      
        lastValue = currValue
    }
    return res 
}