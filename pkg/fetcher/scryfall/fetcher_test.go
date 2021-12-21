package scryfall_test

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/maitesin/mtga/pkg/fetcher"
	"github.com/maitesin/mtga/pkg/fetcher/scryfall"
	"github.com/stretchr/testify/require"
)

type testingRoundTripper struct {
	body string
	statusCode int
	err error
}

func (trt *testingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	if trt.err != nil {
		return nil, trt.err
	} else {
		return &http.Response{
			StatusCode: trt.statusCode,
			Body: io.NopCloser(strings.NewReader(trt.body)),
		}, nil
	}
}

var validRoundTripper http.RoundTripper = &testingRoundTripper{
	body: `{
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

var validCard = fetcher.Card{
	Name: "Opt",
	Language: "en",
	URL: "https://scryfall.com/card/xln/65/opt?utm_source=api",
	SetName: "Ixalan",
	Rarity: "common",
	Image: "ewogICJuYW1lIjogIk9wdCIsCiAgImxhbmciOiAiZW4iLAogICJyZWxlYXNlZF9hdCI6ICIyMDE3LTA5LTI5IiwKICAic2NyeWZhbGxfdXJpIjogImh0dHBzOi8vc2NyeWZhbGwuY29tL2NhcmQveGxuLzY1L29wdD91dG1fc291cmNlPWFwaSIsCiAgImltYWdlX3VyaXMiOiB7CiAgICAicG5nIjogImh0dHBzOi8vYzEuc2NyeWZhbGwuY29tL2ZpbGUvc2NyeWZhbGwtY2FyZHMvcG5nL2Zyb250LzYvNi82NmMzYWE5Zi01Y2IwLTRjOGQtYTA1MC00MjkzODM5ODA3MWIucG5nPzE1NjI1NTY4MDUiCiAgfSwKICAibWFuYV9jb3N0IjogIntVfSIsCiAgInNldF9uYW1lIjogIkl4YWxhbiIsCiAgInJhcml0eSI6ICJjb21tb24iLAogICJwcmljZXMiOiB7CiAgICAiZXVyIjogIjAuMTIiLAogICAgImV1cl9mb2lsIjogIjEuNTQiCiAgfQp9",
	ManaCost: "{U}",
	Reprint: false,
	Price: "0.12",
	ReleasedAt: time.Date(2017, 9, 29, 0, 0, 0, 0, time.UTC),
}

type roundTripperMutator func(http.RoundTripper) http.RoundTripper

var noopRoundTripperMutator = func(tripper http.RoundTripper) http.RoundTripper {
	return tripper
}

func TestFetcher_Fetch(t *testing.T) {
	tests := []struct {
		name    string
		roundTripperMutator roundTripperMutator
		number int
		set    string
		opts   []fetcher.Opt
		wantCard fetcher.Card
		wantErr error
	}{
	{
		name: `Given a working http client, card number and set name,
when the Fetch method on the scryfall Fetcher is called,
then we receive a valid card and no error`,
		roundTripperMutator: noopRoundTripperMutator,
		number: 65,
		set: "xln",
		opts: nil,
		wantCard: validCard,
		wantErr: nil,
	},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			httpClient := http.DefaultClient
			httpClient.Transport = tt.roundTripperMutator(validRoundTripper)

			f := scryfall.NewFetcher(scryfall.WithClient(httpClient))
			got, err := f.Fetch(tt.number, tt.set, tt.opts...)
			if tt.wantErr != nil {
				require.ErrorAs(t, err, &tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wantCard, got)
		})
	}
}

