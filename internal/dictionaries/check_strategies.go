package dictionaries

import (
	"strings"
)

type TextFileCheckingStrategy struct {
	textFileContent []string
}

func NewTextFileCheckingStrategy(textFileContent []string) TextFileCheckingStrategy {
	return TextFileCheckingStrategy{textFileContent}
}

func (s TextFileCheckingStrategy) Check(word string) bool {
	if len(word) == 0 {
		return false
	}
	return wordExistsInSlice(word, s.textFileContent)
}

func wordExistsInSlice(word string, possibleWords []string) bool {
	if len(possibleWords) == 0 {
		return false
	}

	var centerIndex int

	if len(possibleWords)%2 == 0 {
		centerIndex = len(possibleWords) / 2
	} else {
		centerIndex = (len(possibleWords) - 1) / 2
	}

	comparingRes := strings.Compare(word, possibleWords[centerIndex])

	switch comparingRes {
	case -1:
		return wordExistsInSlice(word, possibleWords[:centerIndex])
	case 0:
		return true
	case 1:
		return wordExistsInSlice(word, possibleWords[centerIndex+1:])
	}

	return false
}
