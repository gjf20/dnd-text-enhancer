package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetBodyFromFile(t *testing.T, filename string) []byte {
	currentWorkingDirectory, err := os.Getwd()
	require.NoError(t, err, "Did not expect error getting current directory")

	filepath := filepath.Join(currentWorkingDirectory, "json_response_examples", filename)
	jsonFileContents, err := os.Open(filepath)
	require.NoError(t, err, fmt.Sprintf("Did not expect error opening file %v", filename))
	defer jsonFileContents.Close()

	body, err := ioutil.ReadAll(jsonFileContents)
	require.NoError(t, err, fmt.Sprintf("Did not expect error reading body of file %v", filename))

	return body
}
