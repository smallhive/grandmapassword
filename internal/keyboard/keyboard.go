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

	// letter z has 122 index in asci table
	keysDistance [123][123]int
)

func init() {
	keysDistance = [123][123]int{}
	calculate()
}

func key(a, b rune) int {
	// Cantor Pairing Function
	return int((a+b)*(a+b+1)/2 + a)
}

func Distance(a, b rune) int {
	return keysDistance[a][b]
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

	if pivot == letter {
		keysDistance[pivot][pivot] = 0
		return
	}

	// manhattan keysDistance
	distance := int(math.Abs(float64(originRowID-keyBoardRowID))) + int(math.Abs(float64(pivotID-i)))

	keysDistance[pivot][letter] = distance
	keysDistance[letter][pivot] = distance
}
