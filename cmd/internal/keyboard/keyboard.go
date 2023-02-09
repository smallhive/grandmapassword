package keyboard

import (
	"math"
)

var (
	keyboard = [][]string{
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l", ""},
		{"z", "x", "c", "v", "b", "n", "m", "", "", ""},
	}

	distance map[string]int
)

func init() {
	distance = make(map[string]int)

	calculateDistance()
}

func key(a, b string) string {
	return a + b
}

func Distance(a, b string) int {
	k := key(a, b)
	d, ok := distance[k]
	if !ok {
		return 0
	}

	return d
}

func calculateDistance() {
	for keyBoardRowID, rowSymbols := range keyboard {
		for id, pivot := range rowSymbols {
			if pivot == "" {
				continue
			}

			rowAndFullKeyboardDistance(keyBoardRowID, id, pivot)
		}
	}
}

func rowAndFullKeyboardDistance(originRowID int, pivotID int, pivot string) {
	for keyBoardRowID, rowSymbols := range keyboard {
		// detecting different rows
		diff := int(math.Abs(float64(keyBoardRowID - originRowID)))

		for i, letter := range rowSymbols {
			if letter == "" {
				continue
			}

			storageKey := key(pivot, letter)
			_, ok := distance[storageKey]
			if ok {
				continue
			}

			if pivot == letter {
				distance[storageKey] = 0
				continue
			}

			var d int
			// next letter is on the right on current pivot latter
			if i >= pivotID {
				d = i + diff - pivotID
			} else {
				// we are comparing letters from different rows
				if diff > 0 {
					// next letter closer and closer to us
					d = pivotID + diff - i
				} else {
					d = i + diff + pivotID
				}
			}

			distance[storageKey] = d
			distance[key(letter, pivot)] = d
		}
	}
}
