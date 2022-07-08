package gen

import (
	"context"
	"fmt"
	"github.com/maitesin/mtga/internal/domain"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type InfoGenerator interface {
	Generate(ctx context.Context, card domain.Card) error
}

type InfoCardGenerator struct {
	path string
}

func NewInfoCardGenerator(path string) (*InfoCardGenerator, error) {
	return &InfoCardGenerator{
		path: path,
	}, nil
}

func (f *InfoCardGenerator) Generate(_ context.Context, card Card) error {
	file, err := os.Create(filepath.Join(f.path, fmt.Sprintf("%s_%t.md", card.ID.String(), card.Foil)))
	if err != nil {
		return err
	}

	title := strings.Replace(
		fmt.Sprintf("%s_%s_%s_%s_%t", card.ReleasedAt.Format(time.RFC3339), card.Set, card.Name, card.Languages, card.Foil),
		" ",
		"_",
		-1,
	)

	value := fmt.Sprintf(
		`+++
title = %q
id = %q
date = %q
name = %q
lang = ["%s"]
price = %q
quantity = %d
foil = %t
set = %q
+++
`,
		title,
		card.ID,
		card.ReleasedAt.Format(time.RFC3339),
		card.Name,
		strings.Join(card.Languages, `","`),
		card.Price,
		card.Quantity,
		card.Foil,
		card.Set,
	)

	_, err = io.Copy(file, ioutil.NopCloser(strings.NewReader(value)))

	return err
}
