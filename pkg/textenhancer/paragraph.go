package textenhancer

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gjf20/dnd-text-enhancer/pkg/thesaurus"
	"github.com/jdkato/prose/v2"
)

type Paragraph struct {
	currentWordInd      int
	words               []string
	replacedWords       map[int]string
	thesaurus           thesaurus.Thesaurus
	locationToEditIndex map[int]int
	editIndexToLocation map[int]int
}

func NewParagraph(tokens []prose.Token, thesaurus thesaurus.Thesaurus) Paragraph {
	words := []string{}
	for _, tok := range tokens {
		words = append(words, tok.Text)
	}

	return Paragraph{currentWordInd: 0, words: words, replacedWords: make(map[int]string), thesaurus: thesaurus, locationToEditIndex: make(map[int]int), editIndexToLocation: make(map[int]int)}
}

func (p *Paragraph) GetEditInfo() (string, int) {
	paragraphWithIndices := strings.Builder{}
	for i, word := range p.words {
		var suffix string
		if _, ok := p.replacedWords[i]; ok {
			suffix = fmt.Sprintf("[%v]", p.locationToEditIndex[i]) // insert id numbers to use as edit indices
		}
		var prefix string
		if !isPunctuation(word) {
			prefix = " "
		}

		paragraphWithIndices.WriteString(prefix + word + suffix)
	}
	return strings.Trim(paragraphWithIndices.String(), " "), len(p.replacedWords)
}

func (p *Paragraph) GetReadableText() string {
	paragraphString := strings.Builder{}
	for _, word := range p.words {
		var prefix string
		if !isPunctuation(word) {
			prefix = " "
		}

		paragraphString.WriteString(prefix + word)
	}
	return strings.Trim(paragraphString.String(), " ")
}

func (p *Paragraph) GetEditOptions(editableWordIndex int) []string {
	wordIndex := p.editIndexToLocation[editableWordIndex]
	word := p.replacedWords[wordIndex]
	synonyms, err := p.thesaurus.GetSynonyms(word)
	if err != nil {
		fmt.Printf("\nEncountered error getting synonyms for %v\n", p.words[wordIndex])
		return nil
	}
	options := append([]string{word}, synonyms...)
	return options
}

func (p *Paragraph) SwapWord(wordEditIndex int, newWord string) {
	wordIndex := p.editIndexToLocation[wordEditIndex]
	p.words[wordIndex] = newWord
}

func (p *Paragraph) replaceWithRandomSynonym(word string) {
	wordIndex := p.getIndex(word)
	if wordIndex == -1 {
		return //cannot replace a word that is not in the list
	}
	syns, err := p.thesaurus.GetSynonyms(word)
	if err != nil {
		fmt.Printf("Could not get synonyms for %v, %v will be retained in the enhanced text\n", word, word)
		return
	}

	randomIndex := rand.Intn(len(syns)) //no rand seed so enhancements are different across the same words and consistent across runs.
	randomSynonym := syns[randomIndex]

	p.words[wordIndex] = randomSynonym
	p.replacedWords[wordIndex] = word
	editIndex := len(p.locationToEditIndex) + 1
	p.locationToEditIndex[wordIndex] = editIndex
	p.editIndexToLocation[editIndex] = wordIndex
}

func (p *Paragraph) getIndex(word string) int {
	for i := p.currentWordInd; i < len(p.words); i++ {
		if word == p.words[i] {
			p.currentWordInd = i + 1
			return i
		}
	}
	return -1
}

func isPunctuation(text string) bool {
	textByte := []byte(text)[0]
	return (textByte >= 33 && textByte <= 47) || (textByte >= 58 && textByte <= 64)
}
