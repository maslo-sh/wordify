package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingNewBag(t *testing.T) {
	assert := assert.New(t)
	bag := NewBag()

	assert.Equalf(26, len(bag), "Expected 26 tiles in the bag")

	letters := make([]letter, 0)
	amounts := make([]count, 0)
	for key, value := range bag {
		letters = append(letters, key)
		amounts = append(amounts, value)
	}
}
