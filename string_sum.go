package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	inputRune := []rune(input)
	if len(input) == 0 {
		err := fmt.Errorf("%s", errorEmptyInput)
		return "", err
	}

	var firstValue, secondValue strings.Builder

	i := 0
	for ; i != len(inputRune) && (i == 0 || (inputRune[i] != '+' && inputRune[i] != '-')); i++ {
		if inputRune[i] == '+' {
			continue
		}
		firstValue.WriteRune(inputRune[i])
	}
	if i == len(inputRune) {
		err := fmt.Errorf("%s", errorNotTwoOperands)
		return "", err
	}

	j := i
	for ; j != len(inputRune) && (j == i || (inputRune[j] != '+' && inputRune[j] != '-')); j++ {
		if inputRune[j] == '+' {
			continue
		}
		secondValue.WriteRune(inputRune[j])
	}
	if j != len(inputRune) {
		err := fmt.Errorf("%s", errorNotTwoOperands)
		return "", err
	}

	firstInt, err := strconv.Atoi(firstValue.String())
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	secondInt, err := strconv.Atoi(secondValue.String())
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	output = strconv.Itoa(firstInt + secondInt)

	return output, nil
}
