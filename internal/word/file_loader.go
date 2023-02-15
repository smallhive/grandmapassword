package word

import (
	"bufio"
	"context"
	"os"
)

type FileLoader struct {
	filePath  string
	minLength int
}

func NewFileLoader(filePath string, minLength int) *FileLoader {
	return &FileLoader{filePath: filePath, minLength: minLength}
}

func (l *FileLoader) Load(_ context.Context) ([]string, error) {
	readFile, err := os.Open(l.filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result []string
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if len(txt) < l.minLength {
			continue
		}

		result = append(result, txt)
	}

	return result, nil
}
