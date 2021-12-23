package app

import (
	"context"

	"github.com/google/uuid"

	"github.com/maitesin/mtga/internal/domain"
)

type CardsRepository interface {
	Insert(ctx context.Context, card domain.Card) error
	Update(ctx context.Context, card domain.Card) error
	GetByID(ctx context.Context, id uuid.UUID) (domain.Card, error)
	GetAll(ctx context.Context) ([]domain.Card, error)
}
