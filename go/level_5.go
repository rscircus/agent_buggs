// Objective: Find the correct passcode

package main

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// generate a random passcode in the format: [0..99]-[0..99]
var passcode = randomIntStr(100) + "-" + randomIntStr(100)

func main() {
	codes := readCodesFromKeypad()

	respCode, resp := validateCode(codes)
	if respCode == 0 {
		println("Access Denied:", resp)
	} else {
		println("Access Granted!")
	}
}

// readCodesFromKeypad - get codes from keypad input
func readCodesFromKeypad() []string {
	var passcode string
	var codes = make([]string, 0)

	// Agent Getter - bypass keypad input
	// codes := streamKeypad()

	// Agent Getter - try brute force login
	// TODO: not finished, someone's coming...
	// NOTE: itoa returns a string
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			passcode = strconv.Itoa(x) + "-" + strconv.Itoa(y)
			var tempcodes = make([]string, 0)
			tempcodes = append(tempcodes, passcode)
			// this is super wicked:
			if correct, _ := validateCode(tempcodes); correct == 1 {
				codes = append(codes, passcode)
				return codes
			}

		}
	}
	return codes
}

// validateCode checks if the correct passcode was found
func validateCode(codes []string) (int, string) {
	for i, c := range codes {
		// Epoch: brute-force guard
		if i > 3 {
			return 0, "3 Wrong Guesses - LOCKED!"
		}
		if c == passcode {
			return 1, c
		}
	}
	return 0, "Incorrect codes."
}

// randomIntStr generates a random int from [0..max] and converts it to a string
func randomIntStr(max int64) string {
	num, _ := rand.Int(rand.Reader, big.NewInt(max))
	return strconv.Itoa(int(num.Int64()))
}
