package main

import "testing"

// Tests our implementation of the challenge
func TestChal(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	outputExpected := "Cooking MC's like a pound of bacon"

	outputActual := chal(input)
	if outputActual != outputExpected {
		t.Errorf("Challenge 3 failed! \n\tGot: %s \n\twant: %s.", outputActual, outputExpected)
	}
}
