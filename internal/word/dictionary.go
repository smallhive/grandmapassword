package word

import (
	"context"
	"runtime"
	"sync"

	"github.com/pkg/errors"
	worker_pool "github.com/smallhive/worker-pool"
)

type Loader interface {
	Load(ctx context.Context) ([]string, error)
}

func LoadDictionary(ctx context.Context, loader Loader) (Dictionary, error) {
	words, err := loader.Load(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "dictionary load failed")
	}

	return loadDictionary(ctx, words), nil
}

func loadDictionary(ctx context.Context, words []string) Dictionary {
	pool := worker_pool.NewExportWorker[string](runtime.NumCPU())
	dictionary := make(Dictionary, 0, len(words))
	dictionaryChan := make(chan Word, runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(1)

	// store processed word to the dictionary
	go func() {
		defer wg.Done()

		for {
			pw, ok := <-dictionaryChan
			if !ok {
				break
			}

			dictionary = append(dictionary, pw)
		}
	}()

	consumer := func(_ context.Context, task string) {
		d, err := Difficulty(task)
		if err != nil {
			return
		}

		dictionaryChan <- Word{
			Word:       task,
			Difficulty: d,
			Length:     len(task),
		}
	}
	pool.Consume(ctx, consumer)

	producer := func(_ context.Context, taskInput worker_pool.TaskChan[string]) {
		for _, task := range words {
			taskInput <- task
		}
	}

	pool.Produce(ctx, producer)
	close(dictionaryChan)

	wg.Wait()

	return dictionary
}
