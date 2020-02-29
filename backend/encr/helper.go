package encr

import (
	"fmt"
	"math"
	"strconv"
)

/*
HELPER FUNCTIONS
*/

// split the key code into halfs and the factor
func splitKeyCode(code string) struct {
	First uint64
	Sec   uint64
	Fac   uint64
} {

	// calculate the halfs
	firstHalf := code[:len(code)/2]
	secHalf := code[len(code)/2:]

	if firstHalf == secHalf {
		secHalf += "3"
	}

	// first and last letter combined
	factor := 10*(int(code[0])-48) + int(code[len(code)-1]) - 48

	return struct {
		First uint64
		Sec   uint64
		Fac   uint64
	}{
		First: spiceUpHalfs(firstHalf) + byteSliceToInt(code),
		Sec:   spiceUpHalfs(secHalf) + byteSliceToInt(code),
		Fac:   uint64(factor * len(code)),
	}
}

// check if a given byte is a number
func isNumber(b byte) bool {
	// if b is in the range of numbers
	if b > 47 && b < 58 {
		return true
	}
	return false
}

// verify if the given key is legal
func isCodeValid(code string) error {

	if len(code) < 2 {
		return fmt.Errorf("code is too short: %s", string(code))
	}

	// iterate over code and check if every digit is a number
	for _, c := range code {
		if !isNumber(byte(c)) {
			return fmt.Errorf("invalid code; illegal character found: %s", string(c))
		}
	}

	// check if number is valid
	if byteSliceToInt(code) < 0 {
		return fmt.Errorf("invalid code; too big")
	}

	return nil
}

// int to []byte
func intToByteSlice(code uint64) []byte {
	// convert code to string and than to []byte
	return []byte(strconv.FormatUint(code, 10))
}

// convert a byte slice into an int
func byteSliceToInt(code string) uint64 {
	var out uint64 = 0

	// Iterate over code
	for i, c := range code {

		// current number
		currNum := uint64(c) - 48

		// add currentNumber * 10^position
		out += currNum * uint64(math.Pow(10, float64(len(code)-1-i)))
	}

	return out
}

// Spice up the value of the halfs for the key
func spiceUpHalfs(half string) uint64 {
	val := byteSliceToInt(half)
	fac := uint64(math.Pow(10.0, float64(len(half)/2)))
	// value * 10^(length/2) + value
	return val*fac + val
}
