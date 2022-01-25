package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserInputLine(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	fmt.Println(prompt)
	if scanner.Scan() {
		line = scanner.Text()
	}
	return line
}

func GetWordsToNotChange() map[string]bool {
	line := GetUserInputLine("\nEnter the words that should not be changed (ex. word1 word2 ...): ")

	words := strings.Split(line, " ")
	wordsToRetain := make(map[string]bool, len(words))
	for _, word := range words {
		wordsToRetain[word] = true
	}

	return wordsToRetain
}

func PromptUserForEdit() bool {
	line := GetUserInputLine("Would you like to edit the enhanced text? (y/N): ")
	if len(line) == 0 {
		return false
	}
	firstChar := line[0]

	return strings.EqualFold(string(firstChar), "y")
}

func GetChoiceIndex(choiceCount int) int {
	maxIndex := choiceCount + 1 // add 1 for the "Cancel" choice
	for {
		line := GetUserInputLine("Enter choice index: ")
		chosenIndex64, err := strconv.ParseInt(line, 0, 0)
		chosenIndex := int(chosenIndex64)
		if err != nil {
			fmt.Println("Could not parse number from user input, please try again")
			continue
		}
		if 1 <= chosenIndex && chosenIndex <= maxIndex {
			return chosenIndex
		} else {
			fmt.Printf("Index was invalid, please enter number from 1 to %v\n", maxIndex)
		}
	}
}

func GetEditChoice(choices []string) int {
	const choiceFormat = "\t%v. %v\n" // 1. foo
	for i, choice := range choices {
		fmt.Printf(choiceFormat, i+1, choice)
	}
	fmt.Printf(choiceFormat, len(choices)+1, "Cancel this edit")

	return GetChoiceIndex(len(choices))
}

func UserRequestedExit(userChoice, realChoiceCount int) bool {
	return userChoice == realChoiceCount+1 // exit choice is always last, indexes have been adjusted back to 0 indexing
}
