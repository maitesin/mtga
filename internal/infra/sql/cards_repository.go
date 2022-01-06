package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/upper/db/v4"
)

const CardsTable = "cards"
const fields = "id, name, language, url, set_name, rarity, mana_cost, reprint, price, released_at, opts, quantity, condition"

type Card struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Language   string    `db:"language"`
	URL        string    `db:"url"`
	SetName    string    `db:"set_name"`
	Rarity     string    `db:"rarity"`
	ManaCost   string    `db:"mana_cost"`
	Reprint    bool      `db:"reprint"`
	Price      string    `db:"price"`
	ReleasedAt time.Time `db:"released_at"`
	Opts       int       `db:"opts"`
	Quantity   int       `db:"quantity"`
	Condition  string    `db:"condition"`
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
	rows, err := c.sess.WithContext(ctx).SQL().Query(fmt.Sprintf("SELECT %s FROM %s WHERE ID = %s", fields, CardsTable, id))
	if err != nil {
		return domain.Card{}, err
	}
	defer rows.Close()

	var result Card
	if rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Language,
			&result.URL,
			&result.SetName,
			&result.Rarity,
			&result.ManaCost,
			&result.Reprint,
			&result.Price,
			&result.ReleasedAt,
			&result.Opts,
			&result.Quantity,
			&result.Condition,
		)
		if err != nil {
			return domain.Card{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return domain.Card{}, err
	}

	return toDomain(result)
}

func (c *CardsRepository) GetAll(ctx context.Context) ([]domain.Card, error) {
	rows, err := c.sess.WithContext(ctx).SQL().Query(fmt.Sprintf("SELECT %s FROM %s", fields, CardsTable))
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	var results []Card
	var result Card
	var i int
	for rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Language,
			&result.URL,
			&result.SetName,
			&result.Rarity,
			&result.ManaCost,
			&result.Reprint,
			&result.Price,
			&result.ReleasedAt,
			&result.Opts,
			&result.Quantity,
			&result.Condition,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
		i++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	cards := make([]domain.Card, len(results))
	for i = range results {
		domainCard, err := toDomain(results[i])
		if err != nil {
			return nil, err
		}
		cards[i] = domainCard
	}

	return cards, nil
}

func NewCardsRepository(session db.Session) *CardsRepository {
	return &CardsRepository{
		sess: session,
	}
}

func fromDomain(card domain.Card) Card {
	var opts int
	for _, opt := range card.Opts {
		opts += int(opt)
	}
	return Card{
		ID:         card.ID,
		Name:       card.Name,
		Language:   card.Language,
		URL:        card.URL,
		SetName:    card.SetName,
		Rarity:     card.Rarity,
		ManaCost:   card.ManaCost,
		Reprint:    card.Reprint,
		Price:      card.Price,
		ReleasedAt: card.ReleasedAt,
		Opts:       opts,
		Quantity:   card.Quantity,
		Condition:  string(card.Condition),
	}
}

func toDomain(card Card) (domain.Card, error) {
	condition, err := domain.ConditionFromString(card.Condition)
	if err != nil {
		return domain.Card{}, err
	}
	return domain.Card{
		ID:         card.ID,
		Name:       card.Name,
		Language:   card.Language,
		URL:        card.URL,
		SetName:    card.SetName,
		Rarity:     card.Rarity,
		ManaCost:   card.ManaCost,
		Reprint:    card.Reprint,
		Price:      card.Price,
		ReleasedAt: card.ReleasedAt,
		Opts:       domain.OptsFromInt(card.Opts),
		Quantity:   card.Quantity,
		Condition:  condition,
	}, nil
}
