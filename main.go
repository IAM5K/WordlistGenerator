package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInputsFromUser() []string {
	reader := bufio.NewReader(os.Stdin)
	// Prompt the user to enter inputs
	fmt.Println("Enter the inputs (words, numbers, or special characters).\nYou can separate them with comma or space :")
	inputStr, _ := reader.ReadString('\n')
	// Remove trailing newline character and split the input by space or comma
	inputStr = strings.TrimSpace(inputStr)
	inputs := strings.FieldsFunc(inputStr, func(r rune) bool {
		return r == ' ' || r == ','
	})

	return inputs
}

func generateWordCombinations(words []string) []string {
	var combinations []string

	// Generate permutations and combinations of words
	generatePermutations(words, len(words), &combinations)

	return combinations
}

func generatePermutations(words []string, n int, combinations *[]string) {
	if n == 1 {
		*combinations = append(*combinations, strings.Join(words, ""))
	} else {
		for i := 0; i < n-1; i++ {
			generatePermutations(words, n-1, combinations)
			if n%2 == 0 {
				words[i], words[n-1] = words[n-1], words[i]
			} else {
				words[0], words[n-1] = words[n-1], words[0]
			}
		}
		generatePermutations(words, n-1, combinations)
	}
}

func calculateCombinationsCount(words []string) int {
	n := len(words)

	// Calculate the factorial of n
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

func main() {
	// Get inputs from the user
	inputs := getInputsFromUser()

	// Print each input object in a separate line
	for _, input := range inputs {
		fmt.Println(input)
	}
	// Predict size and count
	// Calculate the number of combinations
	count := calculateCombinationsCount(inputs)

	fmt.Println("Total number of combinations:", count)
	// fmt.Printf("Estimated size consumed by text file: %d bytes\n", size)
	// Generate word combinations
	combinations := generateWordCombinations(inputs)

	// Print the combinations
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
