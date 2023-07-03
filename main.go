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

func generateWordCombinations(words []string, outputFile *os.File) {
	generatePermutations(words, len(words), outputFile)
}

func generatePermutations(words []string, n int, outputFile *os.File) {
	if n == 1 {
		combination := strings.Join(words, "")
		fmt.Fprintln(outputFile, combination) // Write the combination to the output file
	} else {
		for i := 0; i < n-1; i++ {
			generatePermutations(words, n-1, outputFile)
			if n%2 == 0 {
				words[i], words[n-1] = words[n-1], words[i]
			} else {
				words[0], words[n-1] = words[n-1], words[0]
			}
		}
		generatePermutations(words, n-1, outputFile)
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
func estimateFileSize(totalCombinations int, words []string) int64 {
	averageCharSize := 1 // Adjust this value based on your average character size

	totalSize := int64(0)

	for _, word := range words {
		wordSize := int64(len(word) * averageCharSize)
		totalSize += wordSize
	}

	totalSize *= int64(totalCombinations)

	return totalSize
}

func askConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please confirm that you have sufficient space on disk and want to generarte wordlist. 'yes' or 'no': ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "y":
			return true
		case "n":
			return false
		case "yes":
			return true
		case "no":
			return false
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func getInputFilenameWithExtension() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a filename: ")
	input, _ := reader.ReadString('\n')
	filename := strings.TrimSpace(input)

	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}

	return filename
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

	// Calculate the estimated file size
	size := estimateFileSize(count, inputs)
	fmt.Println("Total number of combinations:", count)
	fmt.Printf("Estimated size consumed by text file: %d bytes\n", size)

	confirmed := askConfirmation()
	filename := getInputFilenameWithExtension()

	if confirmed {
		fmt.Println("You confirmed with 'yes'.")
		// Open the output file
		file, err := os.Create("combinations.txt")
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()
		// Generate and store combinations directly to the output file
		generateWordCombinations(inputs, file)
		fmt.Println("Output filename:", filename)
	} else {
		fmt.Println("Process Canceled")
	}

}
