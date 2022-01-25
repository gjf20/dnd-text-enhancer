package thesaurus_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/gjf20/dnd-text-enhancer/mocks"
	"github.com/gjf20/dnd-text-enhancer/pkg/common"
	"github.com/gjf20/dnd-text-enhancer/pkg/thesaurus"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetSynonymsFromCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPortal := mocks.NewMockOnlineThesaurus(ctrl)

	testWord := "shoe"
	expectedSynonym := "boot"
	synonymMap := map[string][]string{testWord: {expectedSynonym}}
	tc := thesaurus.ThesaurusCache{
		WordSynonyms: synonymMap,
		OnlinePortal: mockPortal,
	}

	mockPortal.EXPECT().Query(gomock.Any()).Times(0)

	synList, err := tc.GetSynonyms(testWord)
	require.NoError(t, err)
	require.Equal(t, 1, len(synList))
	require.Equal(t, synList[0], expectedSynonym)
}

func TestGetSynonymsQueriesOnlineThesaurus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPortal := mocks.NewMockOnlineThesaurus(ctrl)

	testWord := "shoe"
	expectedSynonym := "boot"
	synonymMap := map[string][]string{}
	tc := thesaurus.ThesaurusCache{
		WordSynonyms: synonymMap,
		OnlinePortal: mockPortal,
	}

	mockPortal.EXPECT().Query(testWord).Return([]string{expectedSynonym}, nil)

	synList, err := tc.GetSynonyms(testWord)
	require.NoError(t, err)
	require.Equal(t, 1, len(synList))
	require.Equal(t, synList[0], expectedSynonym)
}

func TestGetSynonymsReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPortal := mocks.NewMockOnlineThesaurus(ctrl)

	testWord := "shoe"
	synonymMap := map[string][]string{}
	tc := thesaurus.ThesaurusCache{
		WordSynonyms: synonymMap,
		OnlinePortal: mockPortal,
	}

	testError := "fake error"
	mockPortal.EXPECT().Query(testWord).Return(nil, errors.New(testError))

	synList, err := tc.GetSynonyms(testWord)
	require.Error(t, err)
	require.Contains(t, err.Error(), testError)
	require.Nil(t, synList)
	_, ok := tc.WordSynonyms[testWord]
	require.False(t, ok, "New should not have been added to the cache when there is an error getting the data from the online portal")
}

func TestQuerySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mocks.NewMockHttpClient(ctrl)

	testKey := "my test key"
	testWord := "shoe"

	returnedBody := common.GetBodyFromFile(t, "adjacent_json_result.json")

	thesaurus := thesaurus.MerriamWebsterThesaurus{ApiKey: testKey, Client: mockClient}
	mockClient.EXPECT().Get(HasSubstring(testKey, testWord)).Return(returnedBody, nil)

	synonyms, err := thesaurus.Query(testWord)

	require.NoError(t, err)
	require.True(t, len(synonyms) > 0)
}

func TestQueryClientFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mocks.NewMockHttpClient(ctrl)

	testKey := "my test key"
	testWord := "shoe"
	testError := "fake error"

	thesaurus := thesaurus.MerriamWebsterThesaurus{ApiKey: testKey, Client: mockClient}
	mockClient.EXPECT().Get(HasSubstring(testKey, testWord)).Return(nil, errors.New(testError))

	synonyms, err := thesaurus.Query(testWord)

	require.Error(t, err)
	require.Contains(t, err.Error(), testError)
	require.Nil(t, synonyms)
}

func TestQueryExtractionFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mocks.NewMockHttpClient(ctrl)

	testKey := "my test key"
	testWord := "shoe"
	returnedBody := common.GetBodyFromFile(t, "malformed_slice.json")

	thesaurus := thesaurus.MerriamWebsterThesaurus{ApiKey: testKey, Client: mockClient}
	mockClient.EXPECT().Get(HasSubstring(testKey, testWord)).Return(returnedBody, nil)

	synonyms, err := thesaurus.Query(testWord)

	require.Error(t, err)
	require.Nil(t, synonyms)
}

type hasSubstring struct {
	values []string
}

func (m hasSubstring) Matches(arg interface{}) bool {
	sarg := arg.(string)
	for _, s := range m.values {
		if !strings.Contains(sarg, s) {
			return false
		}
	}
	return true
}

// Unused, but needed for the hasSubstring type to satisfy the matcher interface
func (m hasSubstring) String() string {
	return strings.Join(m.values, ", ")
}

func HasSubstring(values ...string) gomock.Matcher {
	return hasSubstring{values: values}
}
