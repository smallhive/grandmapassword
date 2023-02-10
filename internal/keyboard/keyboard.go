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

	keysDistance map[string]int
)

func init() {
	keysDistance = make(map[string]int)

	calculate()
}

func key(a, b string) string {
	return a + b
}

func Distance(a, b string) int {
	k := key(a, b)
	d, ok := keysDistance[k]
	if !ok {
		return 0
	}

	return d
}

func calculate() {
	for originRowID, rowSymbols := range keyboard {
		for pivotID, pivot := range rowSymbols {
			if pivot == "" {
				continue
			}

			calculateKeyboardWithEachRow(originRowID, pivotID, pivot)
		}
	}
}

func calculateKeyboardWithEachRow(originRowID int, pivotID int, pivot string) {
	for keyBoardRowID, rowSymbols := range keyboard {
		for i, letter := range rowSymbols {
			symbolDistance(letter, originRowID, keyBoardRowID, pivotID, i, pivot)
		}
	}
}

func symbolDistance(letter string, originRowID, keyBoardRowID, pivotID, i int, pivot string) {
	if letter == "" {
		return
	}

	storageKey := key(pivot, letter)
	_, ok := keysDistance[storageKey]
	if ok {
		return
	}

	if pivot == letter {
		keysDistance[storageKey] = 0
		return
	}

	// manhattan keysDistance
	distance := int(math.Abs(float64(originRowID-keyBoardRowID))) + int(math.Abs(float64(pivotID-i)))

	keysDistance[storageKey] = distance
	keysDistance[key(letter, pivot)] = distance
}
