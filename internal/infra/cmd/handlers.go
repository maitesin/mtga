package cmd

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/maitesin/mtga/internal/infra/storage"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
)

func Handle(ctx context.Context, opts Options, repository app.CardsRepository, storage storage.Storage) error {
	condition, err := domain.ConditionFromString(opts.Condition)
	if err != nil {
		return err
	}

	cardsFetcher := scryfall.NewFetcher(http.DefaultClient, rate.NewLimiter(rate.Every(time.Second), 10))
	cardF, err := cardsFetcher.Fetch(opts.Number, opts.Set)
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
