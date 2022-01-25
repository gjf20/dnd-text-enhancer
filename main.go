package main

import (
	"fmt"
	"os"

	"github.com/gjf20/dnd-text-enhancer/pkg/common"
	"github.com/gjf20/dnd-text-enhancer/pkg/textenhancer"
	"github.com/gjf20/dnd-text-enhancer/pkg/thesaurus"
)

const (
	customTextString = "Enhance custom text"
	apiKeyLength     = 36
)

func main() {
	argsWithoutProg := os.Args[1:]
	var apiKey string
	if len(argsWithoutProg) > 0 {
		if len(argsWithoutProg[0]) == apiKeyLength {
			apiKey = argsWithoutProg[0]
		} else {
			fmt.Printf("Merriam Webster Collegiate Thesaurus API Key was not the expected length: %v characters\n", apiKeyLength)
			return
		}
	} else {
		fmt.Printf("Please enter the Merriam Webster Collegiate Thesaurus API Key as a command line argument:\n `go run main.go <api-key>` or `<script> <api-key>`\n")
		return
	}

	mwThesaurus := thesaurus.MerriamWebsterThesaurus{ApiKey: apiKey, Client: &thesaurus.Client{}}
	thesaurusCache := thesaurus.ThesaurusCache{WordSynonyms: make(map[string][]string), OnlinePortal: &mwThesaurus}

	enhanceCustomText(&thesaurusCache)

}

func enhanceCustomText(thesaurus thesaurus.Thesaurus) {
	input := common.GetUserInputLine("\nEnter the text to enhance: ")

	unenhancedWords := common.GetWordsToNotChange()

	paragraph, err := textenhancer.EnhanceText(input, thesaurus, unenhancedWords)
	if err != nil {
		fmt.Printf("\nEncountered error enhancing text: %v\n", err)
	}

	printReadableText(paragraph)

	edit := common.PromptUserForEdit()
	if edit {
		editText(paragraph)
		printReadableText(paragraph)
	}

	return
}

func editText(paragraph textenhancer.Paragraph) {
	fmt.Printf("The enhanced text is printed below.  Words that have been replaced with synonyms will be annotated with '[#]'.\n\n")

	for {

		editText, numEditableWords := paragraph.GetEditInfo()
		fmt.Printf("%v\n", editText)

		fmt.Println("\nTo edit a word, enter its annotation number.")
		fmt.Printf("To cancel edit mode, enter \"%v\"\n", numEditableWords+1)

		wordEditIndex := common.GetChoiceIndex(numEditableWords)
		if common.UserRequestedExit(wordEditIndex, numEditableWords) {
			break
		}

		fmt.Println("\nSynonym Choices: ")
		editChoices := paragraph.GetEditOptions(wordEditIndex)
		userChoice := common.GetEditChoice(editChoices)
		if common.UserRequestedExit(userChoice, len(editChoices)) {
			continue
		}
		paragraph.SwapWord(wordEditIndex, editChoices[userChoice-1])

		fmt.Println("\n\nNew enhanced text:")
	}
}

func printReadableText(paragraph textenhancer.Paragraph) {
	fmt.Println("\nEnhanced Text:")
	fmt.Println(paragraph.GetReadableText())
}
