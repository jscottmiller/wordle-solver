package stringset

import "math/rand"

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

func (s Set) Choose() string {
	var words []string
	for w := range s {
		words = append(words, w)
	}
	idx := rand.Intn(len(words))
	return words[idx]
}
