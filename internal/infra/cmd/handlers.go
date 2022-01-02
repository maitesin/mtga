package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/maitesin/mtga/internal/infra/storage"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
	"io/ioutil"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
)

func Handle(ctx context.Context, repository app.CardsRepository, storage storage.Storage) error {
	var opts Options
	var parser = flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				return nil
			}
			return err
		default:
			return nil
		}
	}

	cardsFetcher := scryfall.NewFetcher(scryfall.WithClient(http.DefaultClient))
	cardF, err := cardsFetcher.Fetch(opts.Number, opts.Set)
	if err != nil {
		return err
	}

	card := domain.NewCard(
		cardF.ID,
		cardF.Name,
		cardF.Language,
		cardF.URL,
		cardF.SetName,
		cardF.Rarity,
		cardF.Image,
		cardF.ManaCost,
		cardF.Reprint,
		cardF.Price,
		cardF.ReleasedAt,
		domain.Regular,
	)

	b, err := base64.RawStdEncoding.DecodeString(card.Image)
	if err != nil {
		return err
	}

	storage.Store(ctx, ioutil.NopCloser(bytes.NewReader(b)))

	return repository.Insert(ctx, *card)
}
