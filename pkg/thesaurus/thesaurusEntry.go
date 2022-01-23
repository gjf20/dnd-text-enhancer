package thesaurus

import (
	"encoding/json"
	"fmt"
)

type thesaurusEntry struct {
	ShortDef            []string          `json:"shortdef"`
	MetaData            metaData          `json:"meta"`
	HeadWordInformation map[string]string `json:"hwi"`
	FunctionalLabel     string            `json:"fl"`
	Definition          []interface{}     `json:"def"`
}

type metaData struct {
	Antonyms  []interface{}     `json:"ants"` //this is an interface because not all entries have antonyms
	ID        string            `json:"id"`
	Offensive bool              `json:"offensive"`
	Section   string            `json:"section"`
	Source    string            `json:"source"`
	Stems     []string          `json:"stems"`
	Synonyms  [][]string        `json:"syns"`
	Target    map[string]string `json:"target"`
	UUID      string            `json:"uuid"`
}

func (te *thesaurusEntry) GetSynonymList() []string {
	return te.MetaData.Synonyms[0]
}

func extractThesaurusEntry(body []byte) (thesaurusEntry, error) {
	var result []thesaurusEntry
	err := json.Unmarshal(body, &result)
	if err != nil {
		return thesaurusEntry{}, fmt.Errorf("recieved error unmarshalling json: %v", err)
	}

	if len(result) < 1 {
		return thesaurusEntry{}, fmt.Errorf("JSON contained no entries: %v", err)
	}

	return result[0], nil
}
