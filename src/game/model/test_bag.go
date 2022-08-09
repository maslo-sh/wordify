package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingNewBag(t *testing.T) {
	assert := assert.New(t)
	bag := NewBag()

	assert.Equalf(26, bag.Tiles.Cardinality(), "Expected 26 tiles in the bag")
}
