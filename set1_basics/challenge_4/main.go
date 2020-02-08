package main

import (
	"encoding/hex"
	"strconv"
	"sort"
	"fmt"
)

// Detect single-character XOR.
func Chal(input string) string {
	// Per rules, convert to bytes
	bytes, _ := hex.DecodeString(input)

	// Map of strings, indexed by the ASCII XOR value
	outputs := make(map[int]string)

	// Keep track of the highest E, A, I, Space count for our Character Freq Analysis
	eai_freqs := make(map[int]int) 

	// ASCII 0-127 -- XOR against each ASCII value
	for i := 0; i < 128; i++ {
		var result[]byte
		for x := 0; x < len(bytes); x++ {
			result = append(result, bytes[x] ^ byte(i))
		}

		// Store result in map as a string
		outputs[i] = string(result)

		// Char freq analysis - e, a, i and space are most common in English sentences
		// Ref: https://www.rosettacode.org/wiki/Letter_frequency#Go
		eai := 0
		for x := 0; x < len(result); x++ {
			if(string(result[x]) == "e" || string(result[x]) == "E" ||
			string(result[x]) == "a" || string(result[x]) == "A" ||
			string(result[x]) == "i" || string(result[x]) == "I" ||
			string(result[x]) == " ") {
				eai++
			}
		}

		// Store eai count
		eai_freqs[i] = eai
	}

	// Sort the map by the eai count, we we can find the max eai
	sorted_eai_freqs := map[int]int{}
	eai_keys := []int{}
	for key, val := range eai_freqs {
		sorted_eai_freqs[val] = key
		eai_keys = append(eai_keys, val)
	}
	sort.Ints(eai_keys)

	// Iterate through our sorted map and locate our max eai
	i := 0
	for _, val := range eai_keys {
		if(i == len(eai_keys) - 1) {
			fmt.Println("XOR'd with: \"" + string(sorted_eai_freqs[val]) + "\"" + " (" + strconv.Itoa(sorted_eai_freqs[val]) + "):")
			fmt.Println("\t" + outputs[sorted_eai_freqs[val]])
			fmt.Println("\t{eai Count: " + strconv.Itoa(val) + "}\n")
			return outputs[sorted_eai_freqs[val]]
		}
		i++
	}

	return "[!] No result found!"
}

func main() {
	//input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	//fmt.Println(Chal(input))

	//TODO: Need to read each line from the input file...
}