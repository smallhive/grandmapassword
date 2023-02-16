package keyboard

import (
	"math"
)

const (
	emptyRune = '0'
)

var (
	keyboard = [][]rune{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', emptyRune},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm', emptyRune, emptyRune, emptyRune},
	}

	keysDistance map[rune]map[rune]int
)

func init() {
	keysDistance = make(map[rune]map[rune]int)
	calculate()
}

func key(a, b rune) int {
	// Cantor Pairing Function
	return int((a+b)*(a+b+1)/2 + a)
}

func Distance(a, b rune) int {
	level1, ok := keysDistance[a]
	if !ok {
		return 0
	}

	v, ok := level1[b]
	if !ok {
		return 0
	}

	return v
}

func calculate() {
	for originRowID, rowSymbols := range keyboard {
		for pivotID, pivot := range rowSymbols {
			if pivot == emptyRune {
				continue
			}

			calculateKeyboardWithEachRow(originRowID, pivotID, pivot)
		}
	}
}

func calculateKeyboardWithEachRow(originRowID int, pivotID int, pivot rune) {
	for keyBoardRowID, rowSymbols := range keyboard {
		for i, letter := range rowSymbols {
			symbolDistance(letter, originRowID, keyBoardRowID, pivotID, i, pivot)
		}
	}
}

func symbolDistance(letter rune, originRowID, keyBoardRowID, pivotID, i int, pivot rune) {
	if letter == emptyRune {
		return
	}

	_, ok := keysDistance[pivot]
	if !ok {
		keysDistance[pivot] = make(map[rune]int)
	}

	_, ok = keysDistance[letter]
	if !ok {
		keysDistance[letter] = make(map[rune]int)
	}

	if pivot == letter {
		keysDistance[pivot][pivot] = 0
		return
	}

	// manhattan keysDistance
	distance := int(math.Abs(float64(originRowID-keyBoardRowID))) + int(math.Abs(float64(pivotID-i)))

	keysDistance[pivot][letter] = distance
	keysDistance[letter][pivot] = distance
}
