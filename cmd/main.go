package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jscottmiller/wordle-solver/cmd/stringset"
)

func main() {
	words := readWords()
	frequencies := readFrequencies()

	var positionedLetters [5]rune
	var knownLetters []rune
	var incorrectLetters [5][]rune
	var absentLetters []rune

	for {
		byLetter := wordsByLetter(words)
		letters := frequentLetters(byLetter)

		i := 0
		var candidates stringset.Set

		for i < len(letters) && (candidates == nil || candidates.Size() > 10) {
			r := letters[i]
			if candidates == nil {
				candidates = byLetter[r]
			} else {
				next := candidates.Intersection(byLetter[r])
				if next.Size() == 0 {
					break
				}
				candidates = next
			}
			i += 1
		}

		guess := candidates.Choose(frequencies)
		fmt.Printf("Guess: %q\n", guess)

		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Correct? (y/n)")
		if input.Scan() {
			if input.Text() == "y" {
				return
			}
		}

	GreenLoop:
		for {
		GetGreen:
			fmt.Print("Green letters? ")
			for input.Scan() {
				for _, letter := range input.Text() {
					idx := strings.Index(guess, string(letter))
					if idx < 0 {
						goto GetGreen
					}
					positionedLetters[idx] = letter

					n := 0
					for _, known := range knownLetters {
						if letter != known {
							knownLetters[n] = known
							n++
						}
					}
					knownLetters = knownLetters[:n]
				}
				break GreenLoop
			}
		}

	YellowLoop:
		for {
		GetYellow:
			fmt.Print("Yellow letters? ")
			for input.Scan() {
				for _, letter := range input.Text() {
					idx := strings.Index(guess, string(letter))
					if idx < 0 {
						goto GetYellow
					}
					knownLetters = append(knownLetters, letter)
					incorrectLetters[idx] = append(incorrectLetters[idx], letter)
				}
				break YellowLoop
			}
		}

	Missing:
		for _, l := range guess {
			for _, p := range positionedLetters {
				if l == p {
					continue Missing
				}
			}
			for _, k := range knownLetters {
				if l == k {
					continue Missing
				}
			}
			absentLetters = append(absentLetters, l)
		}

		var newWords []string
	Word:
		for _, word := range words {
			runes := []rune(word)
			for _, l := range absentLetters {
				if strings.Index(word, string(l)) >= 0 {
					continue Word
				}
			}
			for i, letters := range incorrectLetters {
				for _, l := range letters {
					if strings.Index(word, string(l)) == i {
						continue Word
					}
				}
			}
			for i, l := range positionedLetters {
				if l == rune(0) {
					continue
				} else if runes[i] != l {
					continue Word
				}
			}
			if len(knownLetters) > 0 {
				allKnown := true
				for _, l := range knownLetters {
					allKnown = allKnown && strings.Index(word, string(l)) >= 0
				}
				if !allKnown {
					continue Word
				}
			}
			newWords = append(newWords, word)
		}
		words = newWords

		if len(words) == 0 {
			fmt.Println("No more words found. WTF.")
			return
		}
	}
}

func readWords() []string {
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var words []string
	s := bufio.NewScanner(f)

	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

type frequency struct {
	word  string
	count int
}

func readFrequencies() map[string]int {
	f, err := os.Open("counts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counts := make(map[string]int)
	s := bufio.NewScanner(f)

	for s.Scan() {
		t := s.Text()
		p := strings.Split(t, " ")
		count, _ := strconv.Atoi(p[1])
		counts[p[0]] = count
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return counts
}

func wordsByLetter(words []string) map[rune]stringset.Set {
	byLetter := make(map[rune]stringset.Set)

	for _, word := range words {
		for _, r := range word {
			if _, ok := byLetter[r]; !ok {
				byLetter[r] = stringset.NewSet()
			}
			byLetter[r].Add(word)
		}
	}

	return byLetter
}

type pair struct {
	Key   rune
	Value int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Less(i, j int) bool { return p[i].Value >= p[j].Value }

func frequentLetters(byLetter map[rune]stringset.Set) []rune {
	counts := make(map[rune]int)
	for r, words := range byLetter {
		counts[r] = words.Size()
	}
	var pairs pairList
	for k, v := range counts {
		pairs = append(pairs, pair{k, v})
	}
	sort.Sort(pairs)

	i := 0
	ordered := make([]rune, len(pairs))
	for _, p := range pairs {
		ordered[i] = p.Key
		i++
	}
	return ordered
}
