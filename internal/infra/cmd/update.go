package cmd

import (
	"context"
	"golang.org/x/time/rate"
	"net/http"
	"time"

	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
)

type Update struct {
}

func UpdateHandler(ctx context.Context, _ Update, repository app.CardsRepository) error {
	cardsFetcher := scryfall.NewFetcher(http.DefaultClient, rate.NewLimiter(rate.Every(time.Second), 10))

	cards, err := repository.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, card := range cards {
		updatedCard, err := cardsFetcher.Fetch(card.SetNumber, card.SetShortName, card.Language)
		if err != nil {
			return err
		}
		card.UpdatePrice(updatedCard.Price)
		repository.Update(ctx, card)
	}

	return nil
}
