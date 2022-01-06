package main

import (
	"context"
	"github.com/maitesin/mtga/internal/infra/gen"

	"github.com/maitesin/mtga/config"
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

	//store, err := storage.NewFileSystemStorage(cfg.Storage.Path)
	//if err != nil {
	//	panic(err)
	//}

	cards, err := repository.GetAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, card := range cards {
		generator.Generate(ctx, card)
	}
}
