package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
)

// Command defines the interface of the commands to be performed
type Command interface {
	Name() string
}

// CommandHandler defines the interface of the handler to run commands
type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) error
}

// CreateCardCmd is a VTO
type CreateCardCmd struct {
	ID           uuid.UUID
	CardName     string
	Language     string
	URL          string
	SetName      string
	Rarity       string
	Image        string
	ManaCost     string
	Reprint      bool
	Price        string
	ReleaseAt    time.Time
	Opts         []domain.Opt
	Quantity     int
	Condition    domain.Condition
	SetNumber    int
	SetShortName string
}

// Name returns the name of the command to create a card
func (c CreateCardCmd) Name() string {
	return "createCard"
}

// CreateCardHandler is the handler to create a card
type CreateCardHandler struct {
	repository CardsRepository
}

// NewCreateCardHandler is a constructor
func NewCreateCardHandler(repository CardsRepository) CreateCardHandler {
	return CreateCardHandler{
		repository: repository,
	}
}

// Handle creates a canvas
func (c CreateCardHandler) Handle(ctx context.Context, cmd Command) error {
	createCmd, ok := cmd.(CreateCardCmd)
	if !ok {
		return InvalidCommandError{Expected: CreateCardCmd{}, Received: cmd}
	}

	card := domain.NewCard(
		createCmd.ID,
		createCmd.CardName,
		createCmd.Language,
		createCmd.URL,
		createCmd.SetName,
		createCmd.Rarity,
		createCmd.ManaCost,
		createCmd.Reprint,
		createCmd.Price,
		createCmd.ReleaseAt,
		createCmd.Quantity,
		createCmd.Condition,
		createCmd.SetNumber,
		createCmd.SetShortName,
		createCmd.Opts...,
	)

	return c.repository.Insert(ctx, *card)
}
