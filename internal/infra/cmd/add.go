package cmd

import (
	"bytes"
	"context"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/maitesin/mtga/internal/infra/storage"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
)

type Add struct {
	Number    int    `short:"n" long:"number" description:"Number of the card in the set"`
	Set       string `short:"s" long:"set" description:"Magic the Gathering set of the card"`
	Quantity  int    `short:"q" long:"quantity" default:"1" description:"Number of copies of that card"`
	Foil      bool   `short:"f" long:"foil" description:"Foil card"`
	Altered   bool   `short:"a" long:"altered" description:"Altered card"`
	Signed    bool   `short:"i" long:"signed" description:"Signed card"`
	Condition string `short:"c" long:"condition" default:"nm" description:"Card's condition'"`
	Language  string `short:"l" long:"language" default:"en" description:"Card's language'"`
}

func AddHandler(ctx context.Context, opts Add, repository app.CardsRepository, storage storage.Storage) error {
	condition, err := domain.ConditionFromString(opts.Condition)
	if err != nil {
		return err
	}

	cardsFetcher := scryfall.NewFetcher(http.DefaultClient, rate.NewLimiter(rate.Every(time.Second), 10))
	cardF, err := cardsFetcher.Fetch(opts.Number, opts.Set, opts.Language)
	if err != nil {
		return err
	}

	options := domain.Regular
	if opts.Foil {
		options = domain.Foil
	}
	if opts.Altered {
		options = options & domain.Altered
	}

	card := domain.NewCard(
		cardF.ID,
		cardF.Name,
		cardF.Language,
		cardF.URL,
		cardF.SetName,
		cardF.Rarity,
		cardF.ManaCost,
		cardF.Reprint,
		cardF.Price,
		cardF.ReleasedAt,
		opts.Quantity,
		condition,
		options,
	)

	err = storage.Store(ctx, card.ID, ioutil.NopCloser(bytes.NewReader(cardF.Image)))
	if err != nil {
		return err
	}

	return repository.Insert(ctx, *card)
}
