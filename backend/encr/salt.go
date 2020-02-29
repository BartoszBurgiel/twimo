package encr

import (
	"twimo/backend/encr/base"
)

// Salt generated from a given string
func Salt(str string, base base.Base) (uint64, string, error) {

	// if empty string
	if str == "" {
		return 0, "", nil
	}

	// Convert to base
	// Sum two bytes so that the hashed salt is shorter
	converted := ""
	for i := 0; i < len(str); i++ {
		converted += base.ToBase(uint64(str[i])*uint64(i+len(str)), "")
	}

	// Convert back
	convertedInt := uint64(0)
	convertedInt = base.FromBase(converted)

	// Derive the integer value of the salt from the string
	return convertedInt, converted, nil
}
