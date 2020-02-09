package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Convert hex to base64, pretty simple.
func chal(input string) string {
	// Per rules, convert to bytes
	bytes, _ := hex.DecodeString(input)
	output := base64.StdEncoding.EncodeToString(bytes)
	return output
}

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(chal(input))
}
