package main

import (
	"context"
	"fmt"
	"github.com/maitesin/mtga/config"
	"github.com/maitesin/mtga/internal/infra/cmd"
	sqlx "github.com/maitesin/mtga/internal/infra/sql"
	"github.com/upper/db/v4/adapter/sqlite"
	"os"
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

	if err := cmd.Handle(context.Background(), repository); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
