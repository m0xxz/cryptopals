package main

import (
	"bufio"
	//"encoding/hex"
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
func xorAnalysis() map[int]byte {

}

// Detect single-character XOR.
func Chal() string {
	// Read input file
	data := readData()

	/* for i := range data {
		fmt.Println(data[i])
	} */

	// XOR Analysis

	return ""
}

func main() {
	fmt.Println(Chal())
}
