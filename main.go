package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func jumbleInputs(inputs []string) []string {
	var jumbledInputs []string

	// Jumble each input by permuting characters
	for _, input := range inputs {
		// Convert input to rune slice for character-level manipulation
		inputRunes := []rune(input)

		// Permute the input by swapping characters
		permute(inputRunes, 0, len(input)-1)

		// Append the jumbled input to the result array if its length is 8
		if len(inputRunes) == 8 {
			jumbledInputs = append(jumbledInputs, string(inputRunes))
		}
	}

	return jumbledInputs
}

func permute(input []rune, l, r int) {
	if l == r {
		return
	}

	for i := l; i <= r; i++ {
		input[l], input[i] = input[i], input[l]
		permute(input, l+1, r)
		input[l], input[i] = input[i], input[l]
	}
}

func calculateFileSize(possibilities int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	base := 1024.0

	if possibilities == 0 {
		return "0 B"
	}

	exponent := math.Floor(math.Log(float64(possibilities)) / math.Log(base))
	size := float64(possibilities) / math.Pow(base, exponent)

	sizeStr := strconv.FormatFloat(size, 'f', 2, 64)
	unit := units[int(exponent)]

	return sizeStr + " " + unit
}

func getInputInputsFromUser() []string {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter inputs separated by commas
	fmt.Println("Enter the inputs separated by commas:")
	inputStr, _ := reader.ReadString('\n')

	// Remove trailing newline character and split the input by commas
	inputStr = strings.TrimSpace(inputStr)
	inputs := strings.Split(inputStr, ",")

	return inputs
}

func getOutputFileNameFromUser() string {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the output file name
	fmt.Println("Enter the output file name:")
	fileName, _ := reader.ReadString('\n')

	// Remove trailing newline character
	fileName = strings.TrimSpace(fileName)

	return fileName
}

func generateWordList(inputs []string, fileName string) {
	// Jumble the inputs with a maximum length of 8 characters
	jumbledInputs := jumbleInputs(inputs)

	// Calculate the total number of possibilities
	totalPossibilities := int64(len(jumbledInputs))

	// Calculate the estimated size of the resulting file
	estimatedFileSize := calculateFileSize(totalPossibilities * 8) // Assuming 8 bytes per word

	// Ask the user for confirmation
	fmt.Printf("This will generate a word list with %d possibilities and an estimated file size of %s.\n", totalPossibilities, estimatedFileSize)
	fmt.Print("Do you want to proceed? (y/n): ")
	confirmation, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Process the user's confirmation
	confirmation = strings.TrimSpace(strings.ToLower(confirmation))
	if confirmation != "y" && confirmation != "yes" {
		fmt.Println("Operation aborted.")
		return
	}

	// Create and open the output file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the output file:", err)
		return
	}
	defer file.Close()

	// Append the jumbled inputs to the output file
	for _, input := range jumbledInputs {
		if _, err := file.WriteString(input + "\n"); err != nil {
			fmt.Println("Error writing to the output file:", err)
			return
		}
	}

	fmt.Println("Word list generation complete. Check", fileName, "for the results.")
}

func main() {
	// Get inputs from the user
	inputs := getInputInputsFromUser()

	// Get the output file name from the user
	fileName := getOutputFileNameFromUser()

	// Generate the word list
	generateWordList(inputs, fileName)
}
