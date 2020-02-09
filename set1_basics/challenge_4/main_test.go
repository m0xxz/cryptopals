package main

import "testing"

// Tests our implementation of the challenge
func TestChal(t *testing.T) {
	outputExpected := "Now that the party is jumping"

	outputActual := chal()
	if outputActual != outputExpected {
		t.Errorf("Challenge 4 failed! \n\tGot: %s \n\twant: %s.", outputActual, outputExpected)
	}
}
