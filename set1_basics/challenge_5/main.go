package main

import (
	"encoding/hex"
	"fmt"
)

// Repeating-key XOR
func chal() string {
	// Opening stanza and our repeating key
	line := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")

	i := 0
	var encryptedLine []byte
	for _, b := range line {
		if i > 2 {
			i = 0
		}
		encryptedLine = append(encryptedLine, byte(b^key[i]))
		i++
	}

	return hex.EncodeToString(encryptedLine)
}

func main() {
	fmt.Println(chal())
}
