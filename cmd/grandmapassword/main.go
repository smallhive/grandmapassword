package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallhive/grandmapassword/internal/passwd"
	"github.com/smallhive/grandmapassword/internal/word"
)

const (
	minLength = 3
)

func main() {
	ctx := context.Background()
	word.SetMinWordLength(minLength)

	loader := word.NewFileLoader("words.txt", minLength)
	dictionary, err := word.LoadDictionary(ctx, loader)
	if err != nil {
		log.Fatal(err)
	}

	pwd, err := passwd.Generate(dictionary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Password", pwd)
}
