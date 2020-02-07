package main

import "testing"

// Tests our implementation of the challenge
func TestChal(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output_expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    output_actual := Chal(input)
    if output_actual != output_expected {
       t.Errorf("Challenge 1 failed! \n\tGot: %s \n\twant: %s.", output_actual, output_expected)
    }
}