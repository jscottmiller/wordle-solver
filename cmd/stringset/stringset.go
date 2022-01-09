package stringset

import (
	"fmt"
	"sort"
)

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Has(item string) bool {
	_, present := s[item]
	return present
}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) Union(other Set) Set {
	u := NewSet()
	for item := range s {
		u.Add(item)
	}
	for item := range other {
		u.Add(item)
	}
	return u
}

func (s Set) Intersection(other Set) Set {
	i := NewSet()
	for item := range s {
		if other.Has(item) {
			i.Add(item)

		}
	}
	for item := range other {
		if s.Has(item) {
			i.Add(item)

		}
	}
	return i
}

type wordCountPair struct {
	word  string
	count int
}

type byFrequency []wordCountPair

func (a byFrequency) Len() int           { return len(a) }
func (a byFrequency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFrequency) Less(i, j int) bool { return a[i].count > a[j].count }

func (s Set) Choose(bias map[string]int) string {
	var words byFrequency
	for w := range s {
		words = append(words, wordCountPair{
			w,
			bias[w],
		})
	}
	sort.Sort(words)
	fmt.Println(words)
	return words[0].word
}
