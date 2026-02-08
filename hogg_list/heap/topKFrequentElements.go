package main

import "sort"

func topKFrequent(nums []int, k int) []int {
    // freq map approach?

    freqMap := make(map[int]int)

    // count number of items in the array
    for _, num := range nums {
        freqMap[num]++
    }
    type Pair struct {
        Num int 
        Count int 
    }

    // make a slice 
    pairSlice := make([]Pair, 0, len(freqMap))
    for num, count := range freqMap {
        pairSlice = append(pairSlice, Pair{num, count})
    }

    sort.Slice(pairSlice, func (i, j int) bool {
        return pairSlice[i].Count > pairSlice[j].Count
    })

    result := make([]int, k)
    for i := range k {
        result[i] = pairSlice[i].Num
    }
    return result
}

