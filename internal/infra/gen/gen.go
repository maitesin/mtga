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
	file, err := os.Create(filepath.Join(f.path, card.ID.String()+".md"))
	if err != nil {
		return err
	}

	value := fmt.Sprintf(
		`+++
title = %q
name = %q
lang = ["%s"]
price = %q
quantity = %d
+++
`,
		card.ID,
		card.Name,
		strings.Join(card.Languages, `","`),
		card.Price,
		card.Quantity,
	)

	_, err = io.Copy(file, ioutil.NopCloser(strings.NewReader(value)))

	return err
}
