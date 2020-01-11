package main

import (
	"fmt"
	"io/ioutil"
	dummygenerator "twimo/other/dummyGenerator"
)

func main() {
	// Call the gererator method here
	querries := dummygenerator.Generate()

	// Write querries off to a file
	err := ioutil.WriteFile("setup.sql", querries.Bytes(), 6440)

	// Check if no errors occured
	if err != nil {
		// Print error
		fmt.Println(err)
	}

	// Call the gererator method here
	querries = dummygenerator.GenerateEncoded()

	// Write querries off to a file
	err = ioutil.WriteFile("encodedSetup.sql", querries.Bytes(), 6440)

	// Check if no errors occured
	if err != nil {
		// Print error
		fmt.Println(err)
	}
}
