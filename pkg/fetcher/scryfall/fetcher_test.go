package scryfall_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/maitesin/mtga/pkg/fetcher"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
	"github.com/stretchr/testify/require"
)

type testingRoundTripper struct {
	body       string
	statusCode int
	err        error
	bodyErr    error
}

func (trt *testingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	if trt.err != nil {
		return nil, trt.err
	}
	if trt.bodyErr != nil {
		return &http.Response{
			StatusCode: trt.statusCode,
			Body:       io.NopCloser(newFailingReader(trt.bodyErr)),
		}, nil
	}
	return &http.Response{
		StatusCode: trt.statusCode,
		Body:       io.NopCloser(strings.NewReader(trt.body)),
	}, nil
}

func validRoundTripper() *testingRoundTripper {
	return &testingRoundTripper{
		body: `{
  "id": "66c3aa9f-5cb0-4c8d-a050-42938398071b",
  "name": "Opt",
  "lang": "en",
  "released_at": "2017-09-29",
  "scryfall_uri": "https://scryfall.com/card/xln/65/opt?utm_source=api",
  "image_uris": {
    "png": "https://c1.scryfall.com/file/scryfall-cards/png/front/6/6/66c3aa9f-5cb0-4c8d-a050-42938398071b.png?1562556805"
  },
  "mana_cost": "{U}",
  "set_name": "Ixalan",
  "rarity": "common",
  "prices": {
    "eur": "0.12",
    "eur_foil": "1.54"
  }
}`,
		statusCode: 200,
	}
}

type roundTripperMutator func(*testingRoundTripper) http.RoundTripper

var noopRoundTripperMutator = func(tripper *testingRoundTripper) http.RoundTripper {
	return tripper
}

var validCard = fetcher.Card{
	ID:         uuid.MustParse("66c3aa9f-5cb0-4c8d-a050-42938398071b"),
	Name:       "Opt",
	Language:   "en",
	URL:        "https://scryfall.com/card/xln/65/opt?utm_source=api",
	SetName:    "Ixalan",
	Rarity:     "common",
	Image:      "ewogICJpZCI6ICI2NmMzYWE5Zi01Y2IwLTRjOGQtYTA1MC00MjkzODM5ODA3MWIiLAogICJuYW1lIjogIk9wdCIsCiAgImxhbmciOiAiZW4iLAogICJyZWxlYXNlZF9hdCI6ICIyMDE3LTA5LTI5IiwKICAic2NyeWZhbGxfdXJpIjogImh0dHBzOi8vc2NyeWZhbGwuY29tL2NhcmQveGxuLzY1L29wdD91dG1fc291cmNlPWFwaSIsCiAgImltYWdlX3VyaXMiOiB7CiAgICAicG5nIjogImh0dHBzOi8vYzEuc2NyeWZhbGwuY29tL2ZpbGUvc2NyeWZhbGwtY2FyZHMvcG5nL2Zyb250LzYvNi82NmMzYWE5Zi01Y2IwLTRjOGQtYTA1MC00MjkzODM5ODA3MWIucG5nPzE1NjI1NTY4MDUiCiAgfSwKICAibWFuYV9jb3N0IjogIntVfSIsCiAgInNldF9uYW1lIjogIkl4YWxhbiIsCiAgInJhcml0eSI6ICJjb21tb24iLAogICJwcmljZXMiOiB7CiAgICAiZXVyIjogIjAuMTIiLAogICAgImV1cl9mb2lsIjogIjEuNTQiCiAgfQp9",
	ManaCost:   "{U}",
	Reprint:    false,
	Price:      "0.12",
	ReleasedAt: time.Date(2017, 9, 29, 0, 0, 0, 0, time.UTC),
}

type cardMutator func(fetcher.Card) fetcher.Card

var noopCardMutator = func(card fetcher.Card) fetcher.Card {
	return card
}

type failingReader struct {
	err error
}

func newFailingReader(err error) *failingReader {
	return &failingReader{
		err: err,
	}
}

func (f *failingReader) Read([]byte) (n int, err error) {
	return 0, f.err
}

func TestFetcher_Fetch(t *testing.T) {
	tests := []struct {
		name                string
		roundTripperMutator roundTripperMutator
		number              int
		set                 string
		opts                []fetcher.Opt
		wantCardMutator     cardMutator
		wantErr             error
	}{
		{
			name: `Given a working http client, card number and set name,
when the Fetch method on the scryfall Fetcher is called,
then we receive a valid card and no error`,
			roundTripperMutator: noopRoundTripperMutator,
			number:              65,
			set:                 "xln",
			opts:                nil,
			wantCardMutator:     noopCardMutator,
			wantErr:             nil,
		},
		{
			name: `Given a working http client, card number, set name, and a foil option,
when the Fetch method on the scryfall Fetcher is called,
then we receive a valid card and no error`,
			roundTripperMutator: noopRoundTripperMutator,
			number:              65,
			set:                 "xln",
			opts:                []fetcher.Opt{fetcher.Foil},
			wantCardMutator: func(card fetcher.Card) fetcher.Card {
				card.Price = "1.54"
				return card
			},
			wantErr: nil,
		},
		{
			name: `Given a failing http client,
when the Fetch method on the scryfall Fetcher is called,
then we receive an error`,
			roundTripperMutator: func(tripper *testingRoundTripper) http.RoundTripper {
				tripper.err = errors.New("something went wrong")
				return tripper
			},
			number: 65,
			set:    "xln",
			opts:   nil,
			wantCardMutator: func(fetcher.Card) fetcher.Card {
				return fetcher.Card{}
			},
			wantErr: errors.New("something went wrong"),
		},
		{
			name: `Given a working http client, but with a failing body,
when the Fetch method on the scryfall Fetcher is called,
then we receive an error`,
			roundTripperMutator: func(tripper *testingRoundTripper) http.RoundTripper {
				tripper.bodyErr = errors.New("something went wrong")
				return tripper
			},
			number: 65,
			set:    "xln",
			opts:   nil,
			wantCardMutator: func(fetcher.Card) fetcher.Card {
				return fetcher.Card{}
			},
			wantErr: errors.New("something went wrong"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			httpClient := &http.Client{}
			httpClient.Transport = tt.roundTripperMutator(validRoundTripper())

			f := scryfall.NewFetcher(scryfall.WithClient(httpClient))
			got, err := f.Fetch(tt.number, tt.set, tt.opts...)
			if tt.wantErr != nil {
				require.ErrorAs(t, err, &tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			wantCard := tt.wantCardMutator(validCard)
			require.Equal(t, wantCard, got)
		})
	}
}
