package dictionaries

type Dictionary struct {
	words    []string
	strategy checkingWordStrategy
}

type checkingWordStrategy interface {
	Check(word string) bool
}

func NewDictionary(source []string, strategy checkingWordStrategy) *Dictionary {
	return &Dictionary{source, strategy}
}

func (d *Dictionary) wordExists(word string) bool {
	return d.strategy.Check(word)
}
