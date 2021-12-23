package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/upper/db/v4"
)

const CardsTable = "cards"

type Card struct {
	ID         uuid.UUID    `db:"id"`
	Name       string       `dn:"name"`
	Language   string       `dn:"language"`
	URL        string       `dn:"url"`
	SetName    string       `dn:"set_name"`
	Rarity     string       `dn:"rarity"`
	Image      string       `dn:"image"`
	ManaCost   string       `dn:"mana_cost"`
	Reprint    bool         `dn:"reprint"`
	Price      string       `dn:"price"`
	ReleasedAt time.Time    `dn:"released_at"`
	Opts       []domain.Opt `dn:"opts"`
}

type CardsRepository struct {
	sess db.Session
}

func (c *CardsRepository) Insert(ctx context.Context, card domain.Card) error {
	_, err := c.sess.WithContext(ctx).SQL().InsertInto(CardsTable).Values(fromDomain(card)).Exec()
	return err
}

func (c *CardsRepository) Update(ctx context.Context, card domain.Card) error {
	_, err := c.sess.WithContext(ctx).SQL().Update(CardsTable).Set(fromDomain(card)).Exec()
	return err
}

func (c *CardsRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Card, error) {
	//rows, err := c.sess.WithContext(ctx).SQL().Query()
	//TODO implement me
	panic("implement me")
}

func (c *CardsRepository) GetAll(ctx context.Context) ([]domain.Card, error) {
	//TODO implement me
	panic("implement me")
}

func NewCardsRepository(session db.Session) *CardsRepository {
	return &CardsRepository{
		sess: session,
	}
}

func fromDomain(card domain.Card) Card {
	return Card{
		ID:         card.ID,
		Name:       card.Name,
		Language:   card.Language,
		URL:        card.URL,
		SetName:    card.SetName,
		Rarity:     card.Rarity,
		Image:      card.Image,
		ManaCost:   card.ManaCost,
		Reprint:    card.Reprint,
		Price:      card.Price,
		ReleasedAt: card.ReleasedAt,
		Opts:       card.Opts,
	}
}
