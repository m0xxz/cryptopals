#!/bin/sh
echo "CryptoPals Challenges!"

# Set 1 Basics
echo "Running tests for Set 1 Basics..."
echo "Challenge 1:"
go test cryptopals/set1_basics/challenge_1
echo .

echo "Challenge 2:"
go test cryptopals/set1_basics/challenge_2
echo .

echo "Challenge 3:"
go test cryptopals/set1_basics/challenge_3
echo .

echo "Challenge 4:"
go test cryptopals/set1_basics/challenge_4
echo .

echo "Challenge 5:"
go test cryptopals/set1_basics/challenge_5
echo .

echo "Challenge 6:"
go test cryptopals/set1_basics/challenge_6
echo .

# More here once I complete those challenges

echo "Done!"