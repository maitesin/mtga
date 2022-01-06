package gen

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/maitesin/mtga/internal/domain"
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

func (f *InfoCardGenerator) Generate(_ context.Context, card domain.Card) error {
	file, err := os.Create(filepath.Join(f.path, card.ID.String()+".md"))
	if err != nil {
		return err
	}

	value := fmt.Sprintf(`
+++
title = %q
date = %q
+++
`,
		card.Name,
		card.ReleasedAt.Format(time.RFC3339),
	)

	_, err = io.Copy(file, ioutil.NopCloser(strings.NewReader(value)))

	return err
}
