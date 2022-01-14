package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/config"
	"github.com/maitesin/mtga/internal/infra/cmd"
	sqlx "github.com/maitesin/mtga/internal/infra/sql"
	"github.com/maitesin/mtga/internal/infra/storage"
	"github.com/upper/db/v4/adapter/sqlite"
)

func main() {
	cfg := config.NewConfig()

	var settings = sqlite.ConnectionURL{
		Database: cfg.SQL.DatabaseURL(),
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}

	repository := sqlx.NewCardsRepository(sess)

	store, err := storage.NewFileSystemStorage(cfg.Storage.Path)
	if err != nil {
		panic(err)
	}

	var opts cmd.Options
	var parser = flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			panic(err)
		default:
			os.Exit(0)
		}
	}

	if err := cmd.AddHandler(context.Background(), opts.Add, repository, store); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
