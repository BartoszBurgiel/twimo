package encr

// derive function from the key for the calculation
func (k Key) deriveLockFunction() func(b byte, pos uint64) uint64 {
	n, m := k.getCoefficients()
	return func(b byte, pos uint64) uint64 {
		return (((n - m + 1) * (uint64(b) * (pos + 1))) * k.Div.Fac) + individualConstant(pos+1)
	}
}

// derive function to unlock (slove key) the locked byte
func (k Key) deriveUnlockFunction() func(num uint64, pos uint64) byte {
	n, m := k.getCoefficients()
	return func(num uint64, pos uint64) byte {
		return byte((((num - individualConstant(pos+1)) / (pos + 1)) / k.Div.Fac) / (n - m + 1))
	}
}

// Returns an individual constant depending on the position
func individualConstant(pos uint64) uint64 {
	if pos%2 == 0 {
		return -pos
	}
	return pos
}

// get coefficients n and m for the formula
func (k Key) getCoefficients() (n, m uint64) {
	// Define the coefficients -> n > m
	if k.Div.First > k.Div.Sec {
		n = k.Div.First
		m = k.Div.Sec
	} else {
		m = k.Div.First
		n = k.Div.Sec
	}

	return n, m
}
