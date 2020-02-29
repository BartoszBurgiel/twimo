package base

import "math"

// Base represents a base to which a decimal number will be translated to/from
type Base []byte

// ToBase converts a decimal number to the base
func (b Base) ToBase(u uint64, sol string) string {

	convertedN := uint64(len(b))

	// If beyond limit
	if u >= convertedN {

		// Calculate the value and pass the reminder
		sol += b.ToBase(u/convertedN, b.inBase(u%convertedN))
	} else {
		sol += b.inBase(u)
	}
	return sol
}

// FromBase converts a number in the base to a decimal number
func (b Base) FromBase(str string) (out uint64) {
	out = 0
	// Iterate over string
	for i, s := range str {
		out += uint64(b.getIndex(byte(s))) * uint64(math.Pow(float64(len(b)), float64(i)))
	}
	return out
}

// inBase is a helper function to
func (b Base) inBase(n uint64) string {
	return string(b[n])
}

// getIndex returns the index of the byte in the base character slice
func (b Base) getIndex(by byte) int {
	for i, bb := range b {
		if bb == by {
			return i
		}
	}
	return -1
}
