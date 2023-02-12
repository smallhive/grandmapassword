package word

import (
	"github.com/pkg/errors"

	"github.com/smallhive/grandmapassword/internal/keyboard"
)

const (
	minWordLength = 3
)

// Difficulty calculates finger moving difficulty for word
func Difficulty(w string) (int, error) {
	if len(w) < minWordLength {
		return 0, errors.Wrap(ErrWordTooShort, w)
	}

	return difficulty(w), nil
}

func difficulty(w string) int {
	var d int

	for i := 0; i < len(w)-1; i++ {
		d += keyboard.Distance(
			string(w[i]),
			string(w[i+1]),
		)
	}

	return d
}
