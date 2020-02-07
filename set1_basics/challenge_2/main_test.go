package main

import "testing"

// Tests our implementation of the challenge
func TestChal(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	output_expected := "746865206b696420646f6e277420706c6179"

    output_actual := Chal(input)
    if output_actual != output_expected {
       t.Errorf("Challenge 2 failed! \n\tGot: %s \n\twant: %s.", output_actual, output_expected)
    }
}