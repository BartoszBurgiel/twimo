package encr

import (
	"strconv"
	"strings"
	"twimo/backend/encr/base"
)

// Key represents the key the data will be locked with
type Key struct {
	Value uint64

	// Div represents the key after division
	Div struct {
		// First half of the key
		First uint64

		// Second half of the key
		Sec uint64

		// Factor
		Fac uint64
	}
}

// New key constructor
func New(code string) (Key, error) {
	out := Key{}

	// Verify if code is valid
	err := isCodeValid(code)
	if err != nil {
		return out, err
	}

	out.Value = byteSliceToInt(code)
	out.Div = splitKeyCode(code)

	return out, nil
}

// Lock the given data as a slice of bytes with the key
// and return the encoded string
func (k Key) Lock(data []byte, base base.Base) (string, error) {
	out := ""

	// Encoding function
	encode := k.deriveLockFunction()

	// Iterate over data
	for i, b := range data {

		// encode byte
		encodedByte := encode(b, uint64(i))

		// convert to base
		var enc string

		enc = base.ToBase(encodedByte, "")

		// append to hex
		out += enc

		// dividor
		if i < len(data)-1 {
			out += "-"
		}
	}
	return out, nil
}

// Unlock the hexadecimal string with the key
func (k Key) Unlock(data string, base base.Base) ([]byte, error) {
	out := []byte{}

	// decoding function
	decode := k.deriveUnlockFunction()

	// Divide encoded string
	encodedBytes := strings.Split(data, "-")

	// Iterate over encoded bytes
	for i, eB := range encodedBytes {

		// Check base and convert to base
		var encodedInt uint64
		encodedInt = base.FromBase(eB)
		out = append(out, decode(encodedInt, uint64(i)))
	}

	return out, nil
}

// StringToKey generates a key from the bytes of a given string
func StringToKey(str string) (key Key, err error) {

	processedString := 0

	// Iterate over string's bytes
	for i, b := range str {
		processedString += (i + int(b))
	}

	// Multiply processed string by the length of the
	processedString *= len(str)

	// Turn processed string (int) to string
	keyCode := strconv.Itoa(processedString)

	// Create key
	key, err = New(keyCode)

	return key, err
}
