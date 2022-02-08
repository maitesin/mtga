package main

import (
	"context"
	"github.com/maitesin/mtga/config"
	"github.com/maitesin/mtga/internal/infra/gen"
	sqlx "github.com/maitesin/mtga/internal/infra/sql"
	"github.com/upper/db/v4/adapter/sqlite"
)

func main() {
	cfg := config.NewConfig()
	ctx := context.Background()

	var settings = sqlite.ConnectionURL{
		Database: cfg.SQL.DatabaseURL(),
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}

	repository := sqlx.NewCardsRepository(sess)

	generator, err := gen.NewInfoCardGenerator("docs/content/cards")
	if err != nil {
		panic(err)
	}

	merger := gen.CardMerger{}

	domainCards, err := repository.GetAll(ctx)
	if err != nil {
		panic(err)
	}

	cards := merger.Merge(domainCards)

	for _, card := range cards {
		err = generator.Generate(ctx, card)
		if err != nil {
			panic(err)
		}
	}
}
