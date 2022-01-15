package cmd

import (
	"context"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/internal/domain"
)

type Export struct {
	File flags.Filename `short:"f" long:"file" description:"File to write the export information"`
}

func ExportHandler(ctx context.Context, opts Export, repository app.CardsRepository) error {
	file, err := os.Create(string(opts.File))
	if err != nil {
		return err
	}

	domainCards, err := repository.GetAll(ctx)
	if err != nil {
		return err
	}

	return gocsv.Marshal(fromDomain(domainCards), file)
}

type card struct {
	Name       string    `csv:"name"`
	Language   string    `csv:"language"`
	URL        string    `csv:"url"`
	SetName    string    `csv:"set_name"`
	Rarity     string    `csv:"rarity"`
	Price      string    `csv:"price"`
	ReleasedAt time.Time `csv:"released_at"`
	Quantity   int       `csv:"quantity"`
	Condition  string    `csv:"condition"`
}

func fromDomain(domainCards []domain.Card) []card {
	cards := make([]card, len(domainCards))

	for i, dCard := range domainCards {
		cards[i] = card{
			dCard.Name,
			dCard.Language,
			dCard.URL,
			dCard.SetName,
			dCard.Rarity,
			dCard.Price,
			dCard.ReleasedAt,
			dCard.Quantity,
			string(dCard.Condition),
		}
	}

	return cards
}