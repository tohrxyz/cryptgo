package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LOWERCASE_MIN = rune(97)
	LOWERCASE_MAX = rune(122)
	UPPERCASE_MIN = rune(65)
	UPPERCASE_MAX = rune(90)
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

	encrypted := shiftWholeSentence(input, shiftN)
	fmt.Printf("Encrypted: %s\n", encrypted)
}
