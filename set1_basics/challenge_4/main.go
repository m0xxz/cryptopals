package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	//"sort"
	//"strconv"
)

// Reads data from the 4.txt file provided
func readData() []string {
	data := make([]string, 0)

	file, err := os.Open("set1_basics/challenge_4/4.txt")
	if err != nil {
		log.Fatal(err)
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

// Detect single-character XOR.
func chal() string {
	// Read input file
	data := readData()

	/*for i := range data {
		fmt.Println(data[i])
	}*/

	// XOR Analysis
	xorMap := xorAnalysis(data)

	for i := range xorMap {
		fmt.Println(data[i], "XOR Analysis:")
		for k := range xorMap[i] {
			fmt.Println("\t XOR'd with ", string(k), "\n\t", string(xorMap[i][k]))
		}
	}

	// EAI Character Frequency Analysis

	return ""
}

func main() {
	fmt.Println(chal())
}
