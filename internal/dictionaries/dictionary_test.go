package dictionaries

import (
	_ "embed"
	"github.com/maslokeeper01/golambdas/src/golambdas"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

func TestValidatingCorrectPolishWord(t *testing.T) {
	assert := assert.New(t)

	words := []string{"poprawne", "prawdziwe", "zgodne", "niepoprawne", "wziąć", "świerszcz", "gżegżółka", "psss", "źgać"}

	bytes, err := ioutil.ReadFile("../../resources/languages/polish/words.txt")
	require.NoError(t, err)
	dictContent := strings.Split(strings.Replace(string(bytes), "\r\n", "\n", -1), "\n")
	sort.Strings(dictContent)
	dict := NewDictionary(dictContent, NewTextFileCheckingStrategy(dictContent))

	wordsStream := golambdas.NewStream(words)
	wordsStream.ForEach(func(word string) {
		assert.Truef(dict.wordExists(word), "Provided word should be marked as valid")
	})
}

func TestValidatingIncorrectPolishWord(t *testing.T) {
	assert := assert.New(t)

	words := []string{"niepoprawnexxx", "", "psssii", "qxvvxq", "incorrect", "schlecht"}

	bytes, err := ioutil.ReadFile("../../resources/languages/polish/words.txt")
	require.NoError(t, err)
	dictContent := strings.Split(strings.Replace(string(bytes), "\r\n", "\n", -1), "\n")
	sort.Strings(dictContent)
	dict := NewDictionary(dictContent, NewTextFileCheckingStrategy(dictContent))
	wordsStream := golambdas.NewStream(words)
	wordsStream.ForEach(func(word string) {
		assert.Falsef(dict.wordExists(word), "Provided word should be marked as invalid")
	})

}
