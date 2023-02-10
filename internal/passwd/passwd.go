package passwd

import (
	"sort"

	"github.com/pkg/errors"

	"github.com/smallhive/grandmapassword/internal/word"
)

var (
	ErrZeroGenerated = errors.New("can't generate even a single password")
)

const (
	minLength = 20
	maxLength = 24
)

func isLengthOk(totalLength int) bool {
	return minLength <= totalLength && totalLength <= maxLength
}

func hasIndexIntersection(a, b Pair) bool {
	return a.First == b.First ||
		a.First == b.Second ||
		a.Second == b.First ||
		a.Second == b.Second
}

func Generate(words word.Dictionary) (*word.Word, error) {
	pairs := generatePairs(words)
	sort.Sort(pairs)

	variants := generateVariants(words, pairs)
	sort.Sort(variants)

	if len(variants) == 0 {
		return nil, ErrZeroGenerated
	}

	// return best password according its difficulty
	return &variants[0], nil
}

func generatePairs(words word.Dictionary) PairSlice {
	length := len(words)
	aux := make(PairSlice, 0, length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			aux = append(aux, Pair{
				First:  i,
				Second: j,
				Sum:    words[i].Length + words[j].Length,
			})
		}
	}

	return aux
}

func generateVariants(words word.Dictionary, pairs PairSlice) word.Dictionary {
	var variants word.Dictionary

	i := 0
	j := len(pairs) - 1
	size := len(pairs)

	for i < size && j >= 0 {
		s := pairs[i].Sum + pairs[j].Sum

		if isLengthOk(s) && !hasIndexIntersection(pairs[i], pairs[j]) {
			var dictionary word.Dictionary

			dictionary = append(dictionary,
				word.Word{
					Word:       words[pairs[i].First].Word,
					Length:     words[pairs[i].First].Length,
					Difficulty: words[pairs[i].First].Difficulty,
				},
				word.Word{
					Word:       words[pairs[i].Second].Word,
					Length:     words[pairs[i].Second].Length,
					Difficulty: words[pairs[i].Second].Difficulty,
				},
				word.Word{
					Word:       words[pairs[j].First].Word,
					Length:     words[pairs[j].First].Length,
					Difficulty: words[pairs[j].First].Difficulty,
				},
				word.Word{
					Word:       words[pairs[j].Second].Word,
					Length:     words[pairs[j].Second].Length,
					Difficulty: words[pairs[j].Second].Difficulty,
				},
			)

			variants = append(variants, dictionary.Collapse())
			i++
			j--
		} else if s < minLength {
			i++
		} else {
			j--
		}
	}

	return variants
}
