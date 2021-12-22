package fetcher

import "time"

type Card struct {
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
}

type Opt int

const (
	Plain Opt = iota
	Foil
)

type Fetcher interface {
	Fetch(number int, set string, opts ...Opt) (Card, error)
}
