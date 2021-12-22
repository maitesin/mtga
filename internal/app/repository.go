package app

import (
	"context"

	"github.com/maitesin/mtga/internal/domain"
)

type CardsRepository interface {
	Insert(ctx context.Context, card *domain.Card) error
	Update(ctx context.Context, card *domain.Card) error
}
