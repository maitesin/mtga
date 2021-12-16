package main

import (
	"fmt"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
	"net/http"
)

func main() {
	scryfallFetcher := scryfall.NewFetcher(scryfall.WithClient(http.DefaultClient))

	card, err := scryfallFetcher.Fetch(96, "xln")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", card)
}
