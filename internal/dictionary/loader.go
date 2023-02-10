package dictionary

import (
	"bytes"
	"context"
	"os"

	"github.com/pkg/errors"
)

type Loader interface {
	Load(ctx context.Context) ([]string, error)
}

type FileLoader struct {
	filePath string
}

func NewFileLoader(filePath string) *FileLoader {
	return &FileLoader{filePath: filePath}
}

func (l *FileLoader) Load(_ context.Context) ([]string, error) {
	bts, err := os.ReadFile(l.filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "file %s read err", l.filePath)
	}

	strings := bytes.Split(bts, []byte("\n"))

	result := make([]string, 0, len(strings))
	for _, str := range strings {
		result = append(result, string(str))
	}

	return result, nil
}
