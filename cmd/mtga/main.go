package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
	"github.com/maitesin/mtga/internal/infra/cmd"
	sqlx "github.com/maitesin/mtga/internal/infra/sql"
	"github.com/upper/db/v4/adapter/sqlite"
)

func main() {
	var settings = sqlite.ConnectionURL{
		Database: `cards.db`, // Path to database file
	}

	// Attempt to open the 'example.db' database file
	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}

	repository := sqlx.NewCardsRepository(sess)

	addCommandHandler := app.NewCreateCardHandler(repository)

	_ = addCommandHandler

	var opts cmd.Options

	var parser = flags.NewParser(&opts, flags.Default)

	if values, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			fmt.Printf("Values %+v\n", values)
		}
	}
}
