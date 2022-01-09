## Wordle Solver

A simple solver for [Wordle](https://www.powerlanguage.co.uk/wordle/) puzzles
that uses letter- and word-frequencies to narrow down possible guesses.

### Approach

Wordle maintains two word lists: a list of 2,315 solution words and 10,657
possible words. The union of these two lists forms the set of all valid
guesses; other 5-letter words will be rejected from the game. This combined set
forms the initial universe of possible words. The solver narrows this set until
a solution is found.

For each guess, the solver first calculates the letter frequencies of the
possible words. The candidates for the next guess begin as the set of words
that contain the most frequent letter. From there, this set is narrowed by
iteratively performing an intersection of the set with the next-most-frequent
letter. The intersection operations continue until either the set has fewer
than 10 guesses or the list of letters is exhausted.

When comparing the solution list to the possible word list, it appears that the
words in the solution list are more common than the possible list, with the
possible list containing some proper nouns and archaic words. To account for
this, the solver selects a guess from the candidate list by finding the word
that appears most frequently in the [American National Corpus
(ONC)](https://www.anc.org/).

If the guess is incorrect, the solver will ask for the green and yellow
letters. The solver will use this information to filter the universe of
possible words. The words must contain any yellow letters, not contain yellow
letters in positions that have already been considered, contain green letters
in the correct spots, and not include any previous incorrect letters.

Once the universe is narrowed, letter frequencies are re-calculated and
guessing repeats.

### Running the solver

The solver requires a recent Go install (>1.11).

In the root directly, simply run `go run cmd/main.go`.

The solver assumes two files exist in the current directory: 1) `words.txt`,
which contain the possible words to guess and 2) `counts.txt` which contains
space-separated word frequencies.
