package gen

import (
	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
	"strings"
	"time"
)

type Card struct {
	ID         uuid.UUID
	Name       string
	Languages  []string
	Price      string
	Quantity   int
	Foil       bool
	Set        string
	ReleasedAt time.Time
}

type key struct {
	url  string
	foil bool
}

type Merger interface {
	Merge(cards []domain.Card) []Card
}

type CardMerger struct{}

func (cm *CardMerger) Merge(domainCards []domain.Card) []Card {
	m := make(map[key]*Card)

	// This is not accounting for foils

	for _, dcard := range domainCards {
		url := dcard.URL
		if strings.Count(url, "/") != 6 {
			url = url[:len(url)-3]
		}
		k := key{
			url:  url,
			foil: dcard.IsFoil(),
		}
		if _, ok := m[k]; !ok {
			m[k] = &Card{
				dcard.ID,
				dcard.Name,
				[]string{dcard.Language},
				dcard.Price,
				dcard.Quantity,
				dcard.IsFoil(),
				dcard.SetName,
				dcard.ReleasedAt,
			}
		} else {
			m[k].Quantity += dcard.Quantity
			m[k].Languages = append(m[k].Languages, dcard.Language)
		}
	}

	cards := make([]Card, 0, len(m))

	for _, c := range m {
		var langs []string
		for _, lang := range c.Languages {
			found := false
			for _, l := range langs {
				if l == lang {
					found = true
					break
				}
			}
			if !found {
				langs = append(langs, lang)
			}
		}
		c.Languages = langs
		cards = append(cards, *c)
	}

	return cards
}
