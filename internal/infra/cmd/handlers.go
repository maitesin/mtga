package cmd

import (
	"context"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
)

func Handle(ctx context.Context, repository app.CardsRepository) error {
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

	return repository.Insert(ctx, *card)
}
