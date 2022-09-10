package model

import "fmt"

type count int
type letter rune

type Bag map[letter]count

func (b Bag) Size() int {
	sum := 0
	for _, v := range b {
		sum += int(v)
	}
	return sum
}

func (b Bag) Add(x letter) {
	b[x]++
}

func (b Bag) Delete(x letter) {
	_, ok := b[x]
	if ok {
		b[x]--
	}
}

func (b Bag) Find(x letter) (int, bool) {
	count, ok := b[x]
	return int(count), ok
}

func (b Bag) FindAll(x letter) (int, bool) {
	return b.Find(x) // not useful for this implementation
}

func NewBag() Bag {
	bag := make(Bag)
	for i := 0; i < 44; i++ {
		fmt.Printf("%s", string(letter(i%26+97)))
		bag.Add(letter(i%26 + 97))
	}

	return bag
}
