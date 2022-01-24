package textenhancer

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gjf20/dnd-text-enhancer/pkg/thesaurus"
	"github.com/jdkato/prose/v2"
)

type Paragraph struct {
	currentWordInd int
	words          []string
	replacedWords  map[int]string
	thesaurus      thesaurus.Thesaurus
}

func NewParagraph(tokens []prose.Token, thesaurus thesaurus.Thesaurus) Paragraph {
	words := []string{}
	for _, tok := range tokens {
		words = append(words, tok.Text)
	}

	return Paragraph{currentWordInd: 0, words: words, replacedWords: make(map[int]string), thesaurus: thesaurus} //what about punctuation?
}

func (p *Paragraph) swapWord(word string) {
	wordIndex := p.getIndex(word)
	if wordIndex == -1 {
		return //cannot replace a word that is not in the list
	}
	syns := p.thesaurus.GetSynonyms(word)

	randomIndex := rand.Intn(len(syns)) //no rand seed so enhancements are consistent between program runs
	randomSynonym := syns[randomIndex]

	p.words[wordIndex] = randomSynonym
	p.replacedWords[wordIndex] = word
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

func (p *Paragraph) getEditText() string {
	paragraphWithIndices := strings.Builder{}
	for i, word := range p.words {
		var suffix string
		if _, ok := p.replacedWords[i]; ok {
			suffix = fmt.Sprintf("[%v]", i) // insert id numbers to use as edit indices
		}
		var prefix string
		if !isPunctuation(word) {
			prefix = " "
		}

		paragraphWithIndices.WriteString(prefix + word + suffix)
	}
	return paragraphWithIndices.String()
}

func (p *Paragraph) getReadableText() string {
	paragraphWithIndices := strings.Builder{}
	for _, word := range p.words {
		var prefix string
		if !isPunctuation(word) {
			prefix = " "
		}

		paragraphWithIndices.WriteString(prefix + word)
	}
	return paragraphWithIndices.String()
}

func isPunctuation(text string) bool {
	return strings.Contains(".!?,", text)
}
