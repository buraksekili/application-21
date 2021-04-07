package task1

import (
	"strings"
)

// Task1 is the solution for the question 01
func Task1(input1 string, input2 [][]string) []string {

	// If input1 is empty, return empty
	if len(input1) == 0 {
		return []string{}
	}

	input1Arr := strings.Split(input1, " ")
	lenInput2 := len(input2)
	result := []string{}
	for i, word := range input1Arr {
		if i >= lenInput2 {
			break
		}
		simWords := input2[i]
		for _, sim := range simWords {
			input1Arr[i] = sim
			result = append(result, strings.Join(input1Arr, " "))
		}
		input1Arr[i] = word
	}

	return result

}
