package main

import (
	"encoding/hex"
	"fmt"
)

// Fixed XOR.
func chal(input string) string {
	// We will XOR our input against this, after hex decoding
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	// Per rules, convert to bytes
	bytes, _ := hex.DecodeString(input)

	var output []byte
	for x := 0; x < len(bytes); x++ {
		output = append(output, bytes[x]^b[x])
	}
	return hex.EncodeToString(output)
}

func main() {
	input := "1c0111001f010100061a024b53535009181c"
	fmt.Println(chal(input))
}
