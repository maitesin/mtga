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
	ManaCost   string
	Reprint    bool
	Price      string
	ReleasedAt time.Time
	Opts       []Opt
}

type Opt int

const (
	Regular Opt = 1 << iota
	Foil
	Signed
	Altered
)

func OptsFromInt(opts int) []Opt {
	var optsOut []Opt

	for _, value := range []Opt{Regular, Foil, Signed, Altered} {
		if opts&int(value) == 1 {
			optsOut = append(optsOut, value)
		}
	}

	return optsOut
}

func NewCard(
	id uuid.UUID,
	name string,
	language string,
	URL string,
	setName string,
	rarity string,
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
