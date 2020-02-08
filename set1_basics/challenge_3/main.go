package main

import (
	"encoding/hex"
	"strconv"
	"fmt"
)

// Convert hex to base64, pretty simple.
func Chal(input string) string {
	// Per rules, convert to bytes
	bytes, _ := hex.DecodeString(input)

	// ASCII 0-127
	for i := 0; i < 128; i++ {
		var output []byte
		for x := 0; x < len(bytes); x++ {
			output = append(output, bytes[x] ^ byte(i))
		}

		// Char freq analysis - e, a, i and space are most common in English sentences
		// Ref: https://www.rosettacode.org/wiki/Letter_frequency#Go
		eai_count := 0
		for x := 0; x < len(output); x++ {
			//fmt.Println(string(output[x]))
			if string(output[x]) == "e" || string(output[x]) == "E" ||
			string(output[x]) == "a" || string(output[x]) == "A" ||
			string(output[x]) == "i" || string(output[x]) == "I" ||
			string(output[x]) == " " {
				eai_count++
			}
		}

		// I know there is only 1 result that will match this condition
		// Yeah, it could be better, but for now this is fine.
		if eai_count > 10 {
			fmt.Println("XOR'd with: \"" + string(i) + "\"" + " (" + strconv.Itoa(i) + "):")
			fmt.Println("\t" + string(output))
			fmt.Println("\t{eai Count: " + strconv.Itoa(eai_count) + "}\n")

			return string(output)
		}
	}

	return "[!] No result found!"
}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println(Chal(input))
}