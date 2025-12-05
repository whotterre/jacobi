package main
import "math/bits"


// Reverse bits of a given 32 bits signed integer.
func ReverseBitsLazy(n int) int {
	uintN := uint32(n)
	return int(bits.Reverse32(uintN))
}

// Reverse bits of a given 32 bits signed integer.
func ReverseBitsIterative(n int) int {
	res := 0

	for i := 0; i < 32; i++ {
		res <<= 1
		// get the last bit
		lastBit := n & 1

		// add it to result 
		res = res | lastBit

		// shift by 1 to free the space at the end
		n >>= 1
	}

	return res
}