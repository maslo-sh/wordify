package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingWordFromCorrectRowPlacement(t *testing.T) {
	assert := assert.New(t)

	testPlacements := []placement{{Coords{0, 0}, 'a'},
		placement{Coords{0, 6}, 'a'},
		placement{Coords{0, 4}, 'b'},
		placement{Coords{0, 2}, 'f'},
		placement{Coords{0, 1}, 'm'},
		placement{Coords{0, 5}, 'i'},
		placement{Coords{0, 3}, 'i'},
	}

	testMove := Move{0, testPlacements}

	word, _ := createWordFromMove(testMove)

	assert.Equalf("amfibia", word, "Expected word 'amfibia'")
}

func TestCreatingWordFromCorrectColumnPlacement(t *testing.T) {
	assert := assert.New(t)

	testPlacements := []placement{{Coords{5, 3}, 'g'},
		placement{Coords{5, 6}, 'ź'},
		placement{Coords{5, 4}, 'r'},
		placement{Coords{5, 5}, 'o'},
		placement{Coords{5, 7}, 'b'},
		placement{Coords{5, 8}, 'a'},
	}

	testMove := Move{0, testPlacements}

	word, _ := createWordFromMove(testMove)

	assert.Equalf("groźba", word, "Expected word 'groźba'")
}

func TestRegisteringMoveWithIncorrectPlacement(t *testing.T) {
	assert := assert.New(t)

	testPlacements := []placement{{Coords{5, 3}, 'g'},
		placement{Coords{5, 6}, 'ź'},
		placement{Coords{5, 4}, 'r'},
		placement{Coords{6, 5}, 'o'},
		placement{Coords{5, 7}, 'b'},
		placement{Coords{4, 8}, 'a'},
	}

	testMove := Move{0, testPlacements}

	_, err := createWordFromMove(testMove)
	assert.Errorf(err, "Can't create word from an incorrect placement")
}
