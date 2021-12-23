package main

import (
	"database/sql"
	"github.com/maitesin/mtga/internal/app"
	sql2 "github.com/maitesin/mtga/internal/infra/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/upper/db/v4/adapter/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./cards.db")
	if err != nil {
		panic(err)
	}

	dbConn, err := sqlite.New(db)
	if err != nil {
		panic(err)
	}

	repository := sql2.NewCardsRepository(dbConn)

	addCommandHandler := app.NewCreateCardHandler(repository)

	_ = addCommandHandler
}
