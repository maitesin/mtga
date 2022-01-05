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
	ID            uuid.UUID
	CardName      string
	CardLanguage  string
	CardURL       string
	CardSetName   string
	CardRarity    string
	CardImage     string
	CardManaCost  string
	CardReprint   bool
	CardPrice     string
	CardReleaseAt time.Time
	CardOpts      []domain.Opt
	CardQuantity  int
	CardCondition domain.Condition
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
		createCmd.CardLanguage,
		createCmd.CardURL,
		createCmd.CardSetName,
		createCmd.CardRarity,
		createCmd.CardManaCost,
		createCmd.CardReprint,
		createCmd.CardPrice,
		createCmd.CardReleaseAt,
		createCmd.CardQuantity,
		createCmd.CardCondition,
		createCmd.CardOpts...,
	)

	return c.repository.Insert(ctx, *card)
}
