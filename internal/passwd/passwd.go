package passwd

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/pkg/errors"
	worker_pool "github.com/smallhive/worker-pool"

	"github.com/smallhive/grandmapassword/internal/word"
)

var (
	ErrZeroGenerated = errors.New("can't generate even a single password")
)

type pairIndex struct {
	I int
	J int
}

const (
	minLength = 20
	maxLength = 24

	// 24 is a total length. Min sum length other pair should be minimum 6
	maxPairLength = 18
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
	t := time.Now()
	pairs := generatePairs(words)
	fmt.Println("generatePairs", time.Since(t))
	sort.Sort(pairs)

	t = time.Now()
	variants := generateVariants(words, pairs)
	fmt.Println("generateVariants", time.Since(t))
	sort.Sort(variants)

	if len(variants) == 0 {
		return nil, ErrZeroGenerated
	}

	// return best password according its difficulty
	return &variants[0], nil
}

func generatePairs(words word.Dictionary) PairSlice {
	length := len(words)
	bestDifficultByLength := make(map[int]int)
	bestPairsByLength := make(map[int][]Pair)

	d := 0

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			a := i
			b := j

			// try to understand the best combination of words
			d1 := words[i].Distance(words[j])
			d2 := words[j].Distance(words[i])
			d = d1
			if d2 < d1 {
				a = j
				b = i
				d = d2
			}

			sum := words[a].Length + words[b].Length
			if sum > maxPairLength {
				continue
			}

			p := Pair{
				First:      a,
				Second:     b,
				Sum:        sum,
				Difficulty: words[a].Difficulty + words[b].Difficulty + d,
				W1:         words[a],
				W2:         words[b],
			}

			perfectDifficulty, ok := bestDifficultByLength[p.Sum]
			if !ok {
				bestDifficultByLength[p.Sum] = p.Difficulty
				bestPairsByLength[p.Sum] = []Pair{p}
				continue
			}

			if p.Difficulty <= perfectDifficulty {
				bestDifficultByLength[p.Sum] = p.Difficulty
				bestPairsByLength[p.Sum] = append(bestPairsByLength[p.Sum], p)
			}
		}
	}

	aux := make(PairSlice, 0, len(bestPairsByLength))
	for _, p := range bestPairsByLength {
		aux = append(aux, p...)
	}

	return aux
}

func generateVariants(words word.Dictionary, pairs PairSlice) word.Dictionary {
	mu := sync.Mutex{}
	var variants word.Dictionary

	size := len(pairs)
	muDiff := sync.RWMutex{}
	bestDifficulty := 1_000_000

	pool := worker_pool.NewExportWorker[pairIndex](runtime.NumCPU())
	consumer := func(_ context.Context, task pairIndex) {
		i := task.I
		j := task.J
		s := pairs[i].Sum + pairs[j].Sum

		if isLengthOk(s) && !hasIndexIntersection(pairs[i], pairs[j]) {
			d1 := words[pairs[i].Second].Distance(words[pairs[j].First])
			currentDifficulty := pairs[i].Difficulty + pairs[j].Difficulty + d1

			muDiff.RLock()
			isLessOrEqual := currentDifficulty <= bestDifficulty
			muDiff.RUnlock()

			if isLessOrEqual {
				muDiff.Lock()
				bestDifficulty = currentDifficulty
				muDiff.Unlock()

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

				mu.Lock()
				variants = append(variants, dictionary.Collapse())
				mu.Unlock()
			}
		}
	}

	ctx := context.Background()
	pool.Consume(ctx, consumer)

	producer := func(_ context.Context, taskInput worker_pool.TaskChan[pairIndex]) {
		for i := 0; i < size-1; i++ {
			for j := i + 1; j < size; j++ {
				taskInput <- pairIndex{I: i, J: j}
			}
		}
	}

	pool.Produce(ctx, producer)

	return variants
}
