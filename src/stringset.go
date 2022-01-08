package main

import "math/rand"

type set map[string]struct{}

func NewSet() set {
	return make(set)
}

func (s set) Size() int {
	return len(s)
}

func (s set) Has(item string) bool {
	_, present := s[item]
	return present
}

func (s set) Add(item string) {
	s[item] = struct{}{}
}

func (s set) Union(other set) set {
	u := NewSet()
	for item := range s {
		u.Add(item)
	}
	for item := range other {
		u.Add(item)
	}
	return u
}

func (s set) Intersection(other set) set {
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

func (s set) Choose() string {
	var words []string
	for w := range s {
		words = append(words, w)
	}
	idx := rand.Intn(len(words))
	return words[idx]
}
