package word

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/smallhive/grandmapassword/internal/keyboard"
)

var (
	ErrWordTooShort = errors.New("word too short")
)

type Word struct {
	Word       string
	Length     int
	Difficulty int
}

func (w Word) String() string {
	return fmt.Sprintf("length:%d  difficulty:%d  %s", w.Length, w.Difficulty, w.Word)
}

type Dictionary []Word

func (a Dictionary) Len() int { return len(a) }

func (a Dictionary) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Dictionary) Less(i, j int) bool {
	if a[i].Difficulty == a[j].Difficulty {
		return a[i].Length < a[j].Length
	}

	return a[i].Difficulty < a[j].Difficulty
}

func (a Dictionary) Collapse() Word {
	var word Word
	for i, w := range a {
		word.Word += w.Word
		word.Difficulty += w.Difficulty
		word.Length += w.Length

		// difficulty between words
		if i > 0 {
			prevW := a[i-1].Word
			word.Difficulty += keyboard.Distance(rune(w.Word[0]), rune(prevW[len(prevW)-1]))
		}
	}

	return word
}
