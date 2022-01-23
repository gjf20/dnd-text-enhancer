package thesaurus

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleFolder = "json_response_examples"

func TestExtractThesaurusEntrySuccesses(t *testing.T) {
	tests := [][]string{
		{"with antonyms", "adjacent_json_result.json", "border in common"},
		{"without antonyms", "umpire_json_result.json", "resolves a dispute"},
	}
	for _, args := range tests {
		t.Run(args[0], func(t *testing.T) {
			currentWorkingDirectory, err := os.Getwd()
			require.NoError(t, err, "Did not expect error getting current directory")

			filepath := filepath.Join(currentWorkingDirectory, exampleFolder, args[1])
			jsonFileContents, err := os.Open(filepath)
			require.NoError(t, err, fmt.Sprintf("Did not expect error opening file %v", args[1]))
			defer jsonFileContents.Close()

			body, err := ioutil.ReadAll(jsonFileContents)
			require.NoError(t, err, fmt.Sprintf("Did not expect error reading body of file %v", args[1]))

			thesaurusEntry, err := extractThesaurusEntry(body)
			require.NoError(t, err, "extraction should have been successful")
			require.Contains(t, thesaurusEntry.ShortDef[0], args[2])
		})
	}
}

func TestExtractThesaurusEntryErrors(t *testing.T) {
	tests := [][]string{
		{"unmarshal error", "malformed_meta.json", "unmarshalling json"},
		{"no entries error", "malformed_slice.json", "contained no entries"},
	}
	for _, args := range tests {
		t.Run(args[0], func(t *testing.T) {
			currentWorkingDirectory, err := os.Getwd()
			require.NoError(t, err, "Did not expect error getting current directory")

			filepath := filepath.Join(currentWorkingDirectory, exampleFolder, args[1])
			jsonFileContents, err := os.Open(filepath)
			require.NoError(t, err, fmt.Sprintf("Did not expect error opening file %v", args[1]))
			defer jsonFileContents.Close()

			body, err := ioutil.ReadAll(jsonFileContents)
			require.NoError(t, err, fmt.Sprintf("Did not expect error reading body of file %v", args[1]))

			_, err = extractThesaurusEntry(body)
			require.Error(t, err, "extraction should have failed")
			require.Contains(t, err.Error(), args[2])
		})
	}
}

func TestGetSynonymsList(t *testing.T) {
	syn1 := "greeting"
	syn2 := "salutation"
	metaData := metaData{Synonyms: [][]string{{syn1, syn2}}}
	thesaurusEntry := thesaurusEntry{MetaData: metaData}

	synonyms := thesaurusEntry.GetSynonymList()
	require.Equal(t, 2, len(synonyms))
	require.Equal(t, syn1, synonyms[0])
	require.Equal(t, syn2, synonyms[1])
}
