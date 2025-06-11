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

func getShiftedPos(input rune, shiftAsNum rune, isLowercase bool) rune {
	var min, max rune
	if isLowercase {
		min = LOWERCASE_MIN
		max = LOWERCASE_MAX
	} else {
		min = UPPERCASE_MIN
		max = UPPERCASE_MAX
	}

	rangeVals := max - min + 1

	minimalDifference := max - input

	if shiftAsNum <= minimalDifference {
		return input + shiftAsNum
	} else if shiftAsNum <= (minimalDifference + rangeVals) {
		diff := shiftAsNum - minimalDifference
		return min + diff
	} else {
		remainder := shiftAsNum % rangeVals
		left := remainder - minimalDifference

		if left == 0 {
			return max
		} else if left > 0 {
			return min + left - 1 // -1, because going from max -> min counts as 1 shift
		} else {
			return input + remainder
		}
	}
}

func getShiftedChar(shiftN int, input rune) string {
	shiftNAsNum := rune(shiftN)

	isLowercase, err := isLowercase(input)
	if err != nil {
		panic(err)
	}

	result := getShiftedPos(input, shiftNAsNum, isLowercase)
	return string(result)
}

func shiftWholeSentence(inputString string, shiftN int) string {
	splitBySpace := strings.Fields(inputString)
	var encrypted string

	for j := range splitBySpace {
		part := []rune(splitBySpace[j])

		for i := range part {
			shifted := getShiftedChar(shiftN, part[i])
			encrypted = encrypted + shifted
		}
		encrypted = encrypted + " "
	}
	return encrypted
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

func getShiftBackwardsBy(shiftBy rune, letter rune) string {
	min, max := getEdgesByLetter(letter)
	diffToEdge := letter - min

	if shiftBy <= diffToEdge {
		return string(letter - shiftBy)
	} else if shiftBy > diffToEdge && shiftBy <= (FULLRANGE_LENGTH+diffToEdge) {
		remainder := shiftBy - diffToEdge
		return string(max - remainder)
	} else {
		withoutMinimal := shiftBy - diffToEdge - 1
		remainder := withoutMinimal % FULLRANGE_LENGTH
		return string(max - remainder)
	}
}

func decrypt(input string, shiftN int) string {
	shiftBy := rune(shiftN)
	splitBySpace := strings.Fields(input)
	var decrypted string

	for i := range splitBySpace {
		part := []rune(splitBySpace[i])

		for k := range part {
			shiftedBackwardsChar := getShiftBackwardsBy(shiftBy, part[k])
			decrypted = decrypted + shiftedBackwardsChar
		}
		decrypted = decrypted + " "
	}
	return decrypted
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

	encrypted := shiftWholeSentence(input, shiftN)
	fmt.Printf("Encrypted: %s\n", encrypted)

	fmt.Println()

	decrypted := decrypt(encrypted, shiftN)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
