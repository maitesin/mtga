package cmd

import (
	"context"

	"github.com/jessevdk/go-flags"
	"github.com/maitesin/mtga/internal/app"
)

type Export struct {
	File flags.Filename `short:"f" long:"file" description:"File to write the export information"`
}

func ExportHandler(ctx context.Context, opts Export, repository app.CardsRepository) error {
	return nil
}
