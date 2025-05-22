package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	// Go through the array
	for i := 0; i < len(flowerbed); i++ {
    // Check if we find an empty flowerbed
		if flowerbed[i] == 0 {
      // Check it's adjacent elements - left and right
			leftEmpty := (i == 0) || (flowerbed[i - 1] == 0)
      rightEmpty := (i == len(flowerbed) - 1) || (flowerbed[i + 1] == 0)
      // Place a flower if only the right and left are empty
			if leftEmpty && rightEmpty {
				flowerbed[i] = 1
				count += 1
			}
		}

	}
  // We can only place flowers if we have more or equal free spaces
	return count >= n
}


