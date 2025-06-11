package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LOWERCASE_MIN    = rune(97)
	LOWERCASE_MAX    = rune(122)
	UPPERCASE_MIN    = rune(65)
	UPPERCASE_MAX    = rune(90)
	FULLRANGE_LENGTH = rune(26)
)

func isLowercase(input rune) (bool, error) {
	if input >= UPPERCASE_MIN && input <= UPPERCASE_MAX {
		return false, nil
	} else if input >= LOWERCASE_MIN && input <= LOWERCASE_MAX {
		return true, nil
	} else {
		return false, fmt.Errorf("character %d is neither lowercase or uppercase in unicode table for english alphabet", input)
	}
}

func getEdgesByLetter(input rune) (rune, rune) {
	isLowercase, err := isLowercase(input)
	if err != nil {
		panic(err)
	}

	if isLowercase {
		return LOWERCASE_MIN, LOWERCASE_MAX
	} else {
		return UPPERCASE_MIN, UPPERCASE_MAX
	}
}

func getShiftForwardBy(shiftBy rune, letter rune) string {
	min, max := getEdgesByLetter(letter)
	diffToEdge := max - letter

	if shiftBy <= diffToEdge {
		return string(letter + shiftBy)
	} else if shiftBy > diffToEdge && shiftBy <= (FULLRANGE_LENGTH+diffToEdge) {
		remainder := shiftBy - diffToEdge - 1
		return string(min + remainder)
	} else {
		withoutMinimal := shiftBy - diffToEdge - 1
		remainder := withoutMinimal % FULLRANGE_LENGTH
		return string(min + remainder)
	}
}

func getShiftBackwardsBy(shiftBy rune, letter rune) string {
	min, max := getEdgesByLetter(letter)
	diffToEdge := letter - min

	if shiftBy <= diffToEdge {
		return string(letter - shiftBy)
	} else if shiftBy > diffToEdge && shiftBy <= (FULLRANGE_LENGTH+diffToEdge) {
		remainder := shiftBy - diffToEdge
		return string(max - remainder + 1)
	} else {
		withoutMinimal := shiftBy - diffToEdge - 1
		remainder := withoutMinimal % FULLRANGE_LENGTH
		return string(max - remainder)
	}
}

func cipher(input string, shiftN int, isEncrypt bool) string {
	var shiftBy rune
	if shiftN > int(UPPERCASE_MAX) {
		shiftBy = rune(shiftN % int(FULLRANGE_LENGTH))
	} else {
		shiftBy = rune(shiftN)
	}
	splitBySpace := strings.Fields(input)
	var result string

	for i := range splitBySpace {
		part := []rune(splitBySpace[i])

		for k := range part {
			var shiftedChar string
			if isEncrypt {
				shiftedChar = getShiftForwardBy(shiftBy, part[k])
			} else {
				shiftedChar = getShiftBackwardsBy(shiftBy, part[k])
			}
			result = result + shiftedChar
		}
		result = result + " "
	}
	return result
}

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("Usage: <program> <input string> <number to shift by>")
	}
	input := args[1]

	shiftN, err := strconv.Atoi(args[2])
	if err != nil {
		panic("2nd argument must be of type number")
	}

	fmt.Println("Input: ", input)
	fmt.Println("Shift by: ", shiftN)

	fmt.Println()

	encrypted := cipher(input, shiftN, true)
	fmt.Printf("Encrypted: %s\n", encrypted)

	fmt.Println()

	decrypted := cipher(encrypted, shiftN, false)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
