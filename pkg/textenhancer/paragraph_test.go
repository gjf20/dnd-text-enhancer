package textenhancer

import (
	"errors"
	"testing"

	"github.com/gjf20/dnd-text-enhancer/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jdkato/prose/v2"
	"github.com/stretchr/testify/require"
)

func TestNewParagraph(t *testing.T) {
	expectedToken := prose.Token{Text: "small", Tag: "JJ"}
	tokens := []prose.Token{expectedToken}

	paragraph := NewParagraph(tokens, nil)

	require.Equal(t, 0, paragraph.currentWordInd)
	require.Equal(t, 1, len(paragraph.words))
	require.Equal(t, expectedToken.Text, paragraph.words[0])
	require.NotNil(t, paragraph.replacedWords)
}

func TestReplaceWithRandomSynonym(t *testing.T) {
	expectedReplacement := "tiny"
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	bigToken := prose.Token{Text: "big", Tag: "JJ"}
	tokens := []prose.Token{smallToken, bigToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(smallToken.Text).Return([]string{expectedReplacement}, nil)

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(smallToken.Text)

	require.Equal(t, expectedReplacement, paragraph.words[0])
	require.Equal(t, smallToken.Text, paragraph.replacedWords[0])
}

func TestReplaceWithRandomSynonymNoOp(t *testing.T) {
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	bigToken := prose.Token{Text: "big", Tag: "JJ"}
	tokens := []prose.Token{smallToken, bigToken}
	missingWord := "tiny"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(missingWord).Times(0)

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(missingWord)

	require.Equal(t, 0, len(paragraph.replacedWords))
	require.ElementsMatch(t, []string{smallToken.Text, bigToken.Text}, paragraph.words)
}

func TestGetEditInfo(t *testing.T) {
	expectedReplacement := "tiny"
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	bigToken := prose.Token{Text: "big", Tag: "JJ"}
	tokens := []prose.Token{smallToken, bigToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(smallToken.Text).Return([]string{expectedReplacement}, nil)

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(smallToken.Text)

	editText, editableWords := paragraph.GetEditInfo()

	require.Contains(t, editText, expectedReplacement+"[1]", "Replaced words should have a suffix containing the edit index")
	require.NotContains(t, editText, "[2]", "Words that were not replaced should not have a suffix containing the edit index")
	require.Equal(t, 1, editableWords)
}

func TestGetEditInfoProcessesPunctuation(t *testing.T) {
	period := prose.Token{Text: ".", Tag: "."}
	exclamation := prose.Token{Text: "!", Tag: "!"}
	question := prose.Token{Text: "?", Tag: "?"}
	comma := prose.Token{Text: ",", Tag: ","}
	tokens := []prose.Token{period, exclamation, question, comma}
	expectedText := period.Text + exclamation.Text + question.Text + comma.Text

	paragraph := NewParagraph(tokens, nil)

	editText, editableWords := paragraph.GetEditInfo()

	require.Equal(t, expectedText, editText, "Punctuatuion should not have spaces prepended")
	require.Equal(t, 0, editableWords)
}

func TestGetReadableTextProcessesPunctuation(t *testing.T) {
	period := prose.Token{Text: ".", Tag: "."}
	exclamation := prose.Token{Text: "!", Tag: "!"}
	question := prose.Token{Text: "?", Tag: "?"}
	comma := prose.Token{Text: ",", Tag: ","}
	semicolon := prose.Token{Text: ":", Tag: ":"}
	lessThan := prose.Token{Text: "<", Tag: "<"}
	percent := prose.Token{Text: "%", Tag: "%"}
	dollar := prose.Token{Text: "$", Tag: "$"}

	tokens := []prose.Token{period, exclamation, question, comma, semicolon, lessThan, percent, dollar}
	expectedText := period.Text + exclamation.Text + question.Text + comma.Text + semicolon.Text + lessThan.Text + percent.Text + dollar.Text

	paragraph := NewParagraph(tokens, nil)

	editText := paragraph.GetReadableText()

	require.Equal(t, expectedText, editText, "Punctuatuion should not have spaces prepended")
}

func TestGetEditOptions(t *testing.T) {
	expectedReplacement := "tiny"
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	bigToken := prose.Token{Text: "big", Tag: "JJ"}
	tokens := []prose.Token{smallToken, bigToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(smallToken.Text).Return([]string{expectedReplacement}, nil).Times(2)

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(smallToken.Text)

	editOptions := paragraph.GetEditOptions(0)

	require.Equal(t, 2, len(editOptions))
	require.Equal(t, smallToken.Text, editOptions[0])
	require.Equal(t, expectedReplacement, editOptions[1])
}

func TestSwapWord(t *testing.T) {
	expectedReplacement := "tiny"
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	bigToken := prose.Token{Text: "big", Tag: "JJ"}
	tokens := []prose.Token{bigToken, smallToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(smallToken.Text).Return([]string{expectedReplacement}, nil)

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(smallToken.Text)

	testWord := "test"
	paragraph.SwapWord(1, testWord) //edit index is 1-indexed

	require.Equal(t, bigToken.Text+" "+testWord, paragraph.GetReadableText())
}

func TestReplaceWithRandomSynonymThesaurusError(t *testing.T) {
	expectedReplacement := "foo" //made up word
	smallToken := prose.Token{Text: "small", Tag: "JJ"}
	fooToken := prose.Token{Text: "foo", Tag: "NN"}
	tokens := []prose.Token{smallToken, fooToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockThesaurus := mocks.NewMockThesaurus(ctrl)
	mockThesaurus.EXPECT().GetSynonyms(expectedReplacement).Return(nil, errors.New("fake error"))

	paragraph := NewParagraph(tokens, mockThesaurus)
	paragraph.replaceWithRandomSynonym(expectedReplacement)

	require.Equal(t, 0, len(paragraph.replacedWords))
	require.ElementsMatch(t, []string{smallToken.Text, fooToken.Text}, paragraph.words)
}

func TestGetIndexUpdatesCurrentIndex(t *testing.T) {
	type test struct {
		name                     string
		wordToSearch             string
		expectedReturnedIndex    int
		expectedCurrentWordIndex int
	}

	tests := []test{
		{
			"updates current index when word is found - first word",
			"small",
			0,
			1,
		},
		{
			"updates current index when word is found - second word",
			"big",
			1,
			2,
		},
		{
			"does not update current index when word is not found",
			"tiny",
			-1,
			0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			smallToken := prose.Token{Text: "small", Tag: "JJ"}
			bigToken := prose.Token{Text: "big", Tag: "JJ"}
			tokens := []prose.Token{smallToken, bigToken}

			paragraph := NewParagraph(tokens, nil)
			ind := paragraph.getIndex(test.wordToSearch)

			require.Equal(t, test.expectedReturnedIndex, ind)
			require.Equal(t, test.expectedCurrentWordIndex, paragraph.currentWordInd)
		})
	}

}
