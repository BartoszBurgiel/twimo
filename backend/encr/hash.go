package encr

import (
	"strings"
	"twimo/backend/encr/base"
)

// Hash a string and return the hashed code in given base
// this is only 1-way algorithm -> no unhashing
func Hash(str string, salt string, base base.Base) (string, error) {
	hashedStr := ""
	firstHashedStr := ""

	// Iterate over the string backwards
	for i := 0; i < len(str); i++ {

		// Append hashed byte value in given base
		// Formula for the fitstHashedStr = str[i] * len(str) - str[len-i-1]
		firstHashedStr += base.ToBase(uint64(int(str[i])*len(str)-int(str[len(str)-i-1])), "")
		hashedStr += base.ToBase(uint64(int(str[i])), "")

	}

	// Append hahsed password in given base
	key, err := StringToKey(salt)
	if err != nil {
		return "", err
	}

	// Hash the password
	hashedPass, err := key.Lock([]byte(hashedStr), base)
	if err != nil {
		return "", err
	}

	// Divide the hashed pass by 2
	hashedPass = hashedPass[:int(len(hashedPass)/2)]

	// Remove the seperators
	hashedPass = strings.ReplaceAll(hashedPass, "-", "")

	// Append hashed password to the hashed string
	// by shuffling the hashed string into the hashed password
	hashedPassword := []byte(hashedPass)
	positions := int(len(hashedPassword) / len(hashedStr))
	for i := 0; i < len(hashedStr); i++ {
		hashedPassword[i*positions] = hashedStr[i]
	}

	// Shuffle basedSalt into the hashed strings
	tempFistHS := []byte(firstHashedStr)
	tempHS := []byte(hashedStr)

	// How often should the shuffling algorithm perform -> the shorter slice
	iterations := 0
	if len(tempFistHS) > len(tempHS) {
		iterations = len(tempHS)
	} else {
		iterations = len(tempFistHS)
	}

	for i := 0; i < iterations; i += 2 {

		// If the byte from salt exists
		if i < len(salt) {
			tempFistHS[i] = salt[i]
			tempHS[i] = salt[len(salt)-1-i]
			hashedPassword[i] = salt[i]
			hashedPassword[len(hashedPassword)-i-1] = salt[i]
		}
	}

	// Reswap
	firstHashedStr = string(tempFistHS)
	hashedStr = string(tempHS)

	return firstHashedStr + "-" + string(hashedPassword) + "-" + hashedStr, nil
}
