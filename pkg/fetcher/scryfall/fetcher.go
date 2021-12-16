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

type Fetcher struct {
	client *http.Client
}

type imageUris struct {
	PngURI string `json:"png"`
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
	ReleasedAt string    `json:"released_at"`
}

type Opt func(fetcher *Fetcher) *Fetcher

func WithClient(client *http.Client) func(fetcher *Fetcher) *Fetcher {
	return func(fetcher *Fetcher) *Fetcher {
		fetcher.client = client
		return fetcher
	}
}

func NewFetcher(opts ...Opt) *Fetcher {
	f := &Fetcher{}

	for _, opt := range opts {
		f = opt(f)
	}

	return f
}

func (f *Fetcher) Fetch(number int, set string, _ ...fetcher.Opt) (fetcher.Card, error) {
	response, err := f.client.Get(fmt.Sprintf("%s/cards/%s/%d/", scryfallUrl, set, number))
	if err != nil {
		return fetcher.Card{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fetcher.Card{}, err
	}

	var c card
	err = json.Unmarshal(body, &c)
	if err != nil {
		return fetcher.Card{}, err
	}

	pngResponse, err := f.client.Get(c.URIs.PngURI)
	if err != nil {
		return fetcher.Card{}, nil
	}
	defer pngResponse.Body.Close()

	pngBody, err := ioutil.ReadAll(pngResponse.Body)
	if err != nil {
		return fetcher.Card{}, nil
	}

	pngBase64 := base64.StdEncoding.EncodeToString(pngBody)

	t, err := time.Parse("2006-01-02", c.ReleasedAt)

	return fetcher.Card{
		c.Name,
		c.Language,
		c.Scryfall,
		c.SetName,
		c.Rarity,
		pngBase64,
		c.ManaCost,
		c.Reprint,
		t,
	}, nil
}
