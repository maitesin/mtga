package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maitesin/mtga/internal/domain"
	"github.com/stretchr/testify/require"
)

func validCard() *domain.Card {
	return domain.NewCard(
		uuid.New(),
		"Thoughtseize",
		"es",
		"https://scryfall.com/card/ths/107/thoughtseize",
		"Theros",
		"rare",
		"{B}",
		true,
		"12.00",
		time.Now(),
		1,
		"nm",
		107,
		"ths",
	)
}

func TestCard_UpdatePrice(t *testing.T) {
	tests := []struct {
		name  string
		price string
	}{
		{
			name: `Given a working card,
				   when the price is updated,
                   then the new price is the one in the card`,
			price: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := validCard()
			c.UpdatePrice(tt.price)
			require.Equal(t, tt.price, c.Price)
		})
	}
}

func TestConditionFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCond domain.Condition
		wantErr  error
	}{
		{
			name:     `Mint (long)`,
			input:    "mint",
			wantCond: domain.Mint,
			wantErr:  nil,
		},
		{
			name:     `Mint (short)`,
			input:    "m",
			wantCond: domain.Mint,
			wantErr:  nil,
		},
		{
			name:     `Near Mint (long)`,
			input:    "near_mint",
			wantCond: domain.NearMint,
			wantErr:  nil,
		},
		{
			name:     `Near Mint (short)`,
			input:    "nm",
			wantCond: domain.NearMint,
			wantErr:  nil,
		},
		{
			name:     `Excellent (long)`,
			input:    "excellent",
			wantCond: domain.Excellent,
			wantErr:  nil,
		},
		{
			name:     `Excellent (short)`,
			input:    "e",
			wantCond: domain.Excellent,
			wantErr:  nil,
		},
		{
			name:     `Good (long)`,
			input:    "good",
			wantCond: domain.Good,
			wantErr:  nil,
		},
		{
			name:     `Good (short)`,
			input:    "g",
			wantCond: domain.Good,
			wantErr:  nil,
		},
		{
			name:     `Light Played (long)`,
			input:    "light_played",
			wantCond: domain.LightPlayed,
			wantErr:  nil,
		},
		{
			name:     `Light Played (short)`,
			input:    "lp",
			wantCond: domain.LightPlayed,
			wantErr:  nil,
		},
		{
			name:     `Played (long)`,
			input:    "played",
			wantCond: domain.Played,
			wantErr:  nil,
		},
		{
			name:     `Played (short)`,
			input:    "p",
			wantCond: domain.Played,
			wantErr:  nil,
		},
		{
			name:     `Poor (long)`,
			input:    "poor",
			wantCond: domain.Poor,
			wantErr:  nil,
		},
		{
			name:     `Poor (short)`,
			input:    "po",
			wantCond: domain.Poor,
			wantErr:  nil,
		},
		{
			name:     `Unknown condition`,
			input:    "wololo",
			wantCond: domain.Unknown,
			wantErr:  errors.New(`condition "wololo" unknown`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.ConditionFromString(tt.input)
			if err != nil {
				require.Error(t, tt.wantErr, err)
			} else {
				require.Equal(t, tt.wantCond, got)
			}
		})
	}
}

func TestOptsFromInt(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []domain.Opt
	}{
		{
			name:  `Regular`,
			input: 1,
			want:  []domain.Opt{domain.Regular},
		},
		{
			name:  `Foil`,
			input: 2,
			want:  []domain.Opt{domain.Foil},
		},
		{
			name:  `Signed`,
			input: 4,
			want:  []domain.Opt{domain.Signed},
		},
		{
			name:  `Altered`,
			input: 8,
			want:  []domain.Opt{domain.Altered},
		},
		{
			name:  `Regular|Foil|Signed|Altered`,
			input: 15,
			want:  []domain.Opt{domain.Regular, domain.Foil, domain.Signed, domain.Altered},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := domain.OptsFromInt(tt.input)

			require.Equal(t, tt.want, got)
		})
	}
}
