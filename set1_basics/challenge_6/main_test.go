package main

import "testing"

// Tests our implementation of the challenge
func TestChal(t *testing.T) {
	outputExpected := "blahblah"

	outputActual := chal()
	if outputActual != outputExpected {
		t.Errorf("Challenge 6 failed! \n\tGot: %s \n\twant: %s.", outputActual, outputExpected)
	}
}
