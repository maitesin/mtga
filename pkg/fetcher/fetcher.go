package fetcher

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
	Image      []byte
	ManaCost   string
	Reprint    bool
	Price      string
	ReleasedAt time.Time
}

type Opt int

const (
	Plain Opt = iota
	Foil
)

type Fetcher interface {
	Fetch(number int, set string, lang string, opts ...Opt) (Card, error)
}
