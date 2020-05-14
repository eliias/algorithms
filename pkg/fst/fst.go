package fst

import (
	"unicode/utf8"
)

type Transition struct {
	Letter      string
	Value       int
	Transitions Transitions
}

type Transitions map[rune]*Transition

type FST struct {
	root        Transitions
	suffixCache map[string]*Transition
}

func NewFST() *FST {
	return &FST{
		root:        Transitions{},
		suffixCache: make(map[string]*Transition),
	}
}

func (f *FST) Add(word string, value int) {
	f.add(word, value, f.root)
}

func (f *FST) add(word string, value int, base Transitions) {
	// no more runes, stop
	if len(word) == 0 {
		return
	}

	// get next rune
	r, _ := utf8.DecodeRuneInString(word)

	// get suffix
	suffix := word[1:]

	// walk path if transition already exists, subtract current transition value
	// from new word value
	if t, ok := base[r]; ok {
		// next rune
		f.add(suffix, value-t.Value, t.Transitions)
		return
	}

	// hit suffix cache?
	if t, ok := f.suffixCache[suffix]; ok {
		base[r] = t
		return
	}

	// create new state
	t := &Transition{
		Letter:      string(r),
		Value:       value,
		Transitions: Transitions{},
	}
	base[r] = t

	// save to cache
	f.suffixCache[word] = t

	// TODO: evict "frozen" paths from cache

	// next rune
	f.add(suffix, 0, t.Transitions)
}

func (f *FST) GC() {
	f.suffixCache = make(map[string]*Transition)
}
