package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID           uuid.UUID
	Name         string
	Language     string
	URL          string
	SetName      string
	Rarity       string
	ManaCost     string
	Reprint      bool
	Price        string
	ReleasedAt   time.Time
	Opts         []Opt
	Quantity     int
	Condition    Condition
	SetNumber    int
	SetShortName string
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

type Condition string

const (
	Mint        Condition = "mint"
	NearMint    Condition = "near_mint"
	Excellent   Condition = "excellent"
	Good        Condition = "good"
	LightPlayed Condition = "light_played"
	Played      Condition = "played"
	Poor        Condition = "poor"
	Unknown     Condition = "unknown"
)

func ConditionFromString(cond string) (Condition, error) {
	switch cond {
	case "m", "mint":
		return Mint, nil
	case "nm", "near_mint":
		return NearMint, nil
	case "e", "excellent":
		return Excellent, nil
	case "g", "good":
		return Good, nil
	case "lp", "light_played":
		return LightPlayed, nil
	case "p", "played":
		return Played, nil
	case "po", "poor":
		return Poor, nil
	default:
		return Unknown, fmt.Errorf("condition %q unknown", cond)
	}
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
	quantity int,
	condition Condition,
	setNumber int,
	setShortName string,
	opts ...Opt) *Card {
	return &Card{
		ID:           id,
		Name:         name,
		Language:     language,
		URL:          URL,
		SetName:      setName,
		Rarity:       rarity,
		ManaCost:     manaCost,
		Reprint:      reprint,
		Price:        price,
		ReleasedAt:   releasedAt,
		Opts:         opts,
		Quantity:     quantity,
		Condition:    condition,
		SetNumber:    setNumber,
		SetShortName: setShortName,
	}
}

func (c *Card) UpdatePrice(price string) {
	c.Price = price
}
