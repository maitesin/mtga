package gen

import (
	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
)

type Card struct {
	ID        uuid.UUID
	Name      string
	Languages []string
	Price     string
	Quantity  int
}

type Merger interface {
	Merge(cards []domain.Card) []Card
}

type CardMerger struct{}

func (cm *CardMerger) Merge(domainCards []domain.Card) []Card {
	var cards []Card

	for _, dcard := range domainCards {
		cards = append(cards, Card{
			dcard.ID,
			dcard.Name,
			[]string{dcard.Language},
			dcard.Price,
			dcard.Quantity,
		})
	}

	return cards
}
