package scryfall

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/maitesin/mtga/pkg/fetcher"
	"io/ioutil"
	"net/http"
	"time"
)

const scryfallUrl = "https://api.scryfall.com"

type imageUris struct {
	PngURI string `json:"png"`
}

type prices struct {
	NonFoil string `json:"eur"`
	Foil    string `json:"eur_foil"`
}

type card struct {
	Name       string    `json:"name"`
	Language   string    `json:"language"`
	Scryfall   string    `json:"scryfall_uri"`
	SetName    string    `json:"set_name"`
	Rarity     string    `json:"rarity"`
	URIs       imageUris `json:"image_uris"`
	ManaCost   string    `json:"mana_cost"`
	Reprint    bool      `json:"reprint"`
	Prices     prices    `json:"prices"`
	ReleasedAt string    `json:"released_at"`
}

type Fetcher struct {
	client *http.Client
}

type Opt func(fetcher *Fetcher) *Fetcher

func WithClient(client *http.Client) func(fetcher *Fetcher) *Fetcher {
	return func(fetcher *Fetcher) *Fetcher {
		fetcher.client = client
		return fetcher
	}
}

func NewFetcher(opts ...Opt) *Fetcher {
	f := &Fetcher{
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		f = opt(f)
	}

	return f
}

func (f *Fetcher) Fetch(number int, set string, _ ...fetcher.Opt) (fetcher.Card, error) {
	response, err := f.client.Get(fmt.Sprintf("%s/cards/%s/%d/", scryfallUrl, set, number))
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}

	var c card
	err = json.Unmarshal(body, &c)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}

	pngResponse, err := f.client.Get(c.URIs.PngURI)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}
	defer pngResponse.Body.Close()

	pngBody, err := ioutil.ReadAll(pngResponse.Body)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}

	pngBase64 := base64.StdEncoding.EncodeToString(pngBody)

	t, err := time.Parse("2006-01-02", c.ReleasedAt)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}

	return fetcher.Card{
		Name:       c.Name,
		Language:   c.Language,
		URL:        c.Scryfall,
		SetName:    c.SetName,
		Rarity:     c.Rarity,
		Image:      pngBase64,
		ManaCost:   c.ManaCost,
		Reprint:    c.Reprint,
		Price:      c.Prices.NonFoil,
		ReleasedAt: t,
	}, nil
}
