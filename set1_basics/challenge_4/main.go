package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Reads data from the 4.txt file provided
func readData() []string {
	data := make([]string, 0)

	// Look for our file in two locations
	file, err := os.Open("set1_basics/challenge_4/4.txt")
	if err != nil {
		file, err = os.Open("4.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

// Preforms single character XOR on each data point, for all 0-127 ASCII
// Returns xorMap[index][XOR'd char as int]
func xorAnalysis(data []string) map[int]map[int][]byte {
	xorMap := make(map[int]map[int][]byte, 0)

	for i := range data {
		xorMap[i] = make(map[int][]byte)

		// Per rules, convert to bytes
		bytes, _ := hex.DecodeString(data[i])

		// ASCII 0-127 -- XOR against each ASCII value
		for ch := 0; ch < 128; ch++ {
			xorMap[i][ch] = make([]byte, 0)
			for x := 0; x < len(bytes); x++ {
				xorMap[i][ch] = append(xorMap[i][ch], bytes[x]^byte(ch))
			}
		}
	}

	return xorMap
}

// Preforms Character Frequency Analysis
// Returns eaiMap[index][XOR'd char as int][eai_count]
func eaiAnalysis(xorMap map[int]map[int][]byte) map[int]map[int]map[int][]byte {
	eaiMap := make(map[int]map[int]map[int][]byte, 0)

	for i := range xorMap {
		eaiMap[i] = make(map[int]map[int][]byte)

		for ch := range xorMap[i] {
			eaiMap[i][ch] = make(map[int][]byte)

			// Char freq analysis - e, a, i and space are most common in English sentences
			// Ref: https://www.rosettacode.org/wiki/Letter_frequency#Go
			eai := 0
			for x := 0; x < len(xorMap[i][ch]); x++ {
				if string(xorMap[i][ch][x]) == "e" || string(xorMap[i][ch][x]) == "E" ||
					string(xorMap[i][ch][x]) == "a" || string(xorMap[i][ch][x]) == "A" ||
					string(xorMap[i][ch][x]) == "i" || string(xorMap[i][ch][x]) == "I" ||
					string(xorMap[i][ch][x]) == " " {
					eai++
				}
			}

			eaiMap[i][ch][eai] = xorMap[i][ch]
		}
	}

	return eaiMap
}

// Detect single-character XOR.
func chal() string {
	// Read input file
	data := readData()

	// XOR Analysis
	xorMap := xorAnalysis(data)

	// EAI Character Frequency Analysis
	eaiMap := eaiAnalysis(xorMap)

	// Find the max EAI value
	maxI, maxK, maxJ := 0, 0, 0
	for i := range eaiMap {
		for k := range eaiMap[i] {
			for j := range eaiMap[i][k] {
				if j > maxJ {
					maxJ = j
					maxK = k
					maxI = i
				}
			}
		}
	}

	// Output our findings
	fmt.Println("[*] ", data[maxI])
	fmt.Println("\t XOR'd w/ '" + string(maxK) + "' (ASCII " + strconv.Itoa(maxK) + ")")
	fmt.Println("\t " + strings.TrimSuffix(string(xorMap[maxI][maxK]), "\n"))

	// Return our result, stripping new lines
	return strings.TrimSuffix(string(xorMap[maxI][maxK]), "\n")
}

func main() {
	chal()
}
