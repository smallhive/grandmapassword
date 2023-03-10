package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/smallhive/grandmapassword/internal/passwd"
	"github.com/smallhive/grandmapassword/internal/word"
)

const (
	minLength = 3
)

var (
	fileToLoad = "words_5k.txt"
)

func main() {
	if len(os.Args) > 1 {
		fileToLoad = os.Args[1]
	}

	t := time.Now()

	ctx := context.Background()
	word.SetMinWordLength(minLength)

	loader := word.NewFileLoader(fileToLoad, minLength)
	dictionary, err := word.LoadDictionary(ctx, loader)
	if err != nil {
		log.Fatal(err)
	}

	pwd, err := passwd.Generate(dictionary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nPassword", pwd, "generated in", time.Since(t))
}
