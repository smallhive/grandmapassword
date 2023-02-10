package main

import (
	"context"
	"log"
	"sort"

	"github.com/smallhive/grandmapassword/internal/dictionary"
	"github.com/smallhive/grandmapassword/internal/passwd"
)

func main() {
	ctx := context.Background()

	loader := dictionary.NewFileLoader("words.txt")
	words, err := passwd.ProcessDictionary(ctx, loader)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(words)
}
