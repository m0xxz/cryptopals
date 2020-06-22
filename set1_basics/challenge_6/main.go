package main

import (
	"fmt"
)

// Compute the hamming distance between two byte arrays
func hamming(a, b []byte) int {
	if len(a) != len(b) {
		fmt.Println("[!] Hamming: a,b must be same length!")
	}

	dist := 0
	for i := range a {
		// 8 bits in 1 byte
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (a[i] & mask) != (b[i] & mask) {
				dist++
			}
		}
	}
	return dist
}

// Break repeating-key XOR
func chal() string {
	fmt.Println(hamming([]byte("this is a test"), []byte("wokka wokka!!!")))
	return ""
}

func main() {
	fmt.Println(chal())
}
