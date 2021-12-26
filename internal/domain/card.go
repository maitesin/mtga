package domain

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID         uuid.UUID
	Name       string
	Language   string
	URL        string
	SetName    string
	Rarity     string
	Image      string
	ManaCost   string
	Reprint    bool
	Price      string
	ReleasedAt time.Time
	Opts       []Opt
}

type Opt int

const (
	Plain Opt = 1 << iota
	Foil
	Signed
	Altered
)

func NewCard(
	id uuid.UUID,
	name string,
	language string,
	URL string,
	setName string,
	rarity string,
	image string,
	manaCost string,
	reprint bool,
	price string,
	releasedAt time.Time,
	opts ...Opt) *Card {
	return &Card{
		ID:         id,
		Name:       name,
		Language:   language,
		URL:        URL,
		SetName:    setName,
		Rarity:     rarity,
		Image:      image,
		ManaCost:   manaCost,
		Reprint:    reprint,
		Price:      price,
		ReleasedAt: releasedAt,
		Opts:       opts,
	}
}

func (c *Card) UpdatePrice(price string) {
	c.Price = price
}
