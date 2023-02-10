package passwd

import (
	"context"
	"runtime"

	"github.com/pkg/errors"
	worker_pool "github.com/smallhive/worker-pool"

	"github.com/smallhive/grandmapassword/internal/word"
)

type Loader interface {
	Load(ctx context.Context) ([]string, error)
}

func ProcessDictionary(ctx context.Context, loader Loader) (ProcessedWordSlice, error) {
	words, err := loader.Load(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "dictionary load failed")
	}

	p := processDictionary(ctx, words)
	return p, nil
}

func processDictionary(ctx context.Context, words []string) []ProcessedWord {
	pool := worker_pool.NewExportWorker[string](runtime.NumCPU())

	result := make([]ProcessedWord, 0, len(words))
	processedWordsChan := make(chan ProcessedWord, runtime.NumCPU())
	go func() {
		for {
			pw, ok := <-processedWordsChan
			if !ok {
				break
			}

			result = append(result, pw)
		}
	}()

	consumer := func(_ context.Context, task string) {
		d, err := word.Difficulty(task)
		if err != nil {
			return
		}

		processedWordsChan <- ProcessedWord{
			Word:       task,
			Difficulty: d,
		}
	}
	pool.Consume(ctx, consumer)

	producer := func(_ context.Context, taskInput worker_pool.TaskChan[string]) {
		for _, task := range words {
			taskInput <- task
		}
	}

	pool.Produce(ctx, producer)
	close(processedWordsChan)

	return result
}
