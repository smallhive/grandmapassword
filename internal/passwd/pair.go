package passwd

import (
	"fmt"

	"github.com/smallhive/grandmapassword/internal/word"
)

type Pair struct {
	First      int
	Second     int
	Sum        int
	Difficulty int
	W1         word.Word
	W2         word.Word
}

func (p Pair) String() string {
	return fmt.Sprintf("%d %d %d", p.First, p.Second, p.Difficulty)
}

type PairSlice []Pair

func (a PairSlice) Len() int      { return len(a) }
func (a PairSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PairSlice) Less(i, j int) bool {
	if a[i].Sum == a[j].Sum {
		return a[i].Difficulty < a[j].Difficulty
	}

	return a[i].Sum < a[j].Sum
}
