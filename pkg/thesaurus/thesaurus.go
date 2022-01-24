package thesaurus

import (
	"fmt"
	"io"
	"net/http"
)

type Thesaurus interface {
	GetSynonyms(string) []string
}

type OnlineThesaurus interface {
	Query(string) ([]string, error)
}

type HttpClient interface {
	Get(string) ([]byte, error)
}

type ThesaurusCache struct {
	WordSynonyms map[string][]string
	OnlinePortal OnlineThesaurus
}

type MerriamWebsterThesaurus struct {
	ApiKey string
	Client HttpClient
}

type Client struct{}

func (t *ThesaurusCache) GetSynonyms(word string) []string {
	syns, ok := t.WordSynonyms[word]
	if !ok {
		var err error
		syns, err = t.OnlinePortal.Query(word)
		if err != nil {
			return nil
		}
		t.WordSynonyms[word] = syns
	}

	return syns
}

func (mwt *MerriamWebsterThesaurus) Query(word string) ([]string, error) {
	const thesaurusURL = "https://www.dictionaryapi.com/api/v3/references/thesaurus/json/%v?key=%v"
	wordQuery := fmt.Sprintf(thesaurusURL, word, mwt.ApiKey)
	body, err := mwt.Client.Get(wordQuery)
	if err != nil {
		return nil, fmt.Errorf("unexpected error from http client: %v", err)
	}

	thesaurusEntry, err := extractThesaurusEntry(body)
	if err != nil {
		return nil, fmt.Errorf("unexpected error from extracting json: %v", err)
	}

	syns := thesaurusEntry.GetSynonymList()
	return syns, nil
}

func (c *Client) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("recieved error executing request: %v", err)
		return nil, nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("recieved error reading body: %v", err)
		return nil, nil
	}

	return body, nil
}
