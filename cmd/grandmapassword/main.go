package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallhive/grandmapassword/internal/passwd"
	"github.com/smallhive/grandmapassword/internal/word"
)

func main() {
	ctx := context.Background()

	loader := word.NewFileLoader("words.txt")
	dictionary, err := word.LoadDictionary(ctx, loader)
	if err != nil {
		log.Fatal(err)
	}

	pwd, err := passwd.Generate(dictionary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pwd)
}
