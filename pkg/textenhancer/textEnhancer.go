package textenhancer

import (
	"fmt"

	"github.com/gjf20/dnd-text-enhancer/pkg/thesaurus"
	prose "github.com/jdkato/prose/v2"
)

func EnhanceText(inputText string, thesaurus thesaurus.Thesaurus, wordsToRetain map[string]bool) (Paragraph, error) {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(inputText)
	if err != nil {
		return Paragraph{}, fmt.Errorf("error processing the input text: %v", err)
	}

	paragraph := NewParagraph(doc.Tokens(), thesaurus)

	fmt.Println("Processing text... please wait a moment")

	for _, tok := range doc.Tokens() {
		if wordsToRetain[tok.Text] {
			continue
		}
		if tagIsInteresting(tok.Tag) {
			paragraph.replaceWithRandomSynonym(tok.Text)
		}
	}

	return paragraph, nil
}

func tagIsInteresting(tag string) bool {
	return tagIsAdjective(tag) || tagIsNoun(tag) || tagIsAdverb(tag)
}

func tagIsAdjective(tag string) bool {
	return tag == "JJ" || tag == "JJR" || tag == "JJS"
}

func tagIsNoun(tag string) bool {
	return tag == "NN" || tag == "NNS"
}

func tagIsAdverb(tag string) bool {
	return tag == "RB" || tag == "RBR" || tag == "RBS" || tag == "RP"
}
