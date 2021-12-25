package main

import (
	"fmt"
	"github.com/maitesin/mtga/config"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/internal/infra/cmd"
	sqlx "github.com/maitesin/mtga/internal/infra/sql"
	"github.com/upper/db/v4/adapter/sqlite"
)

func main() {
	cfg := config.NewConfig()

	var opts cmd.Options

	var parser = flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(0)
		}
	}

	var settings = sqlite.ConnectionURL{
		Database: cfg.SQL.DatabaseURL(),
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}

	repository := sqlx.NewCardsRepository(sess)

	addCommandHandler := app.NewCreateCardHandler(repository)

	_ = addCommandHandler

	fmt.Printf("Opts %+v\n", opts)
}
