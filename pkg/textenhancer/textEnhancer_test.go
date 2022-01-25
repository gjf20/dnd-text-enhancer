package textenhancer

import (
	"fmt"
	"testing"

	"github.com/gjf20/dnd-text-enhancer/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestEnhanceText(t *testing.T) {
	word1 := "hard"
	synonym1 := "durable"
	word2 := "tree"
	synonym2 := "plant"
	inputText := word1 + " " + word2
	expectedReadableText := synonym1 + " " + synonym2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(word1).Return([]string{synonym1}, nil)
	mockThesaurus.EXPECT().GetSynonyms(word2).Return([]string{synonym2}, nil)

	paragraph, err := EnhanceText(inputText, mockThesaurus, nil)
	require.NoError(t, err)
	require.Equal(t, expectedReadableText, paragraph.GetReadableText())
}

func TestEnhanceTextRetainWord(t *testing.T) {
	word1 := "hard"
	synonym1 := "durable"
	word2 := "tree"
	inputText := word1 + " " + word2
	expectedReadableText := synonym1 + " " + word2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(word1).Return([]string{synonym1}, nil)
	mockThesaurus.EXPECT().GetSynonyms(word2).Times(0)

	paragraph, err := EnhanceText(inputText, mockThesaurus, map[string]bool{word2: true})
	require.NoError(t, err)
	require.Equal(t, expectedReadableText, paragraph.GetReadableText())
}

var adjTags = []string{"JJ", "JJR", "JJS"}
var nounTags = []string{"NN", "NNS"}
var advTags = []string{"RB", "RBR", "RBS", "RP"}

func TestTagIsAdjective(t *testing.T) {
	for _, test := range adjTags {
		t.Run(fmt.Sprintf("is adjective: %v", test), func(t *testing.T) {
			require.True(t, tagIsAdjective(test))
		})
	}

	otherTags := append(nounTags, advTags...)
	for _, test := range otherTags {
		t.Run(fmt.Sprintf("is adjective: %v", test), func(t *testing.T) {
			require.False(t, tagIsAdjective(test))
		})
	}
}

func TestTagIsNoun(t *testing.T) {
	for _, test := range nounTags {
		t.Run(fmt.Sprintf("is noun: %v", test), func(t *testing.T) {
			require.True(t, tagIsNoun(test))
		})
	}

	otherTags := append(adjTags, advTags...)
	for _, test := range otherTags {
		t.Run(fmt.Sprintf("is noun: %v", test), func(t *testing.T) {
			require.False(t, tagIsNoun(test))
		})
	}
}

func TestTagIsAdverb(t *testing.T) {
	for _, test := range advTags {
		t.Run(fmt.Sprintf("is adverb: %v", test), func(t *testing.T) {
			require.True(t, tagIsAdverb(test))
		})
	}

	otherTags := append(adjTags, nounTags...)
	for _, test := range otherTags {
		t.Run(fmt.Sprintf("is adverb: %v", test), func(t *testing.T) {
			require.False(t, tagIsAdverb(test))
		})
	}
}
