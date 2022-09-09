package scryfall

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/google/uuid"
	"github.com/maitesin/mtga/pkg/fetcher"
)

const scryfallUrl = "https://api.scryfall.com"

type imageUris struct {
	PngURI string `json:"png"`
}

type prices struct {
	NonFoil string `json:"eur"`
	Foil    string `json:"eur_foil"`
}

type face struct {
	URIs *imageUris `json:"image_uris"`
}

type card struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Language   string     `json:"lang"`
	Scryfall   string     `json:"scryfall_uri"`
	SetName    string     `json:"set_name"`
	Rarity     string     `json:"rarity"`
	URIs       *imageUris `json:"image_uris,omitempty"`
	ManaCost   string     `json:"mana_cost"`
	Reprint    bool       `json:"reprint"`
	Prices     prices     `json:"prices"`
	ReleasedAt string     `json:"released_at"`
	Faces      []face     `json:"card_faces,omitempty"`
}

type Fetcher struct {
	client *RLHTTPClient
}

type Opt func(fetcher *Fetcher) *Fetcher

func NewFetcher(c *http.Client, limiter *rate.Limiter) *Fetcher {
	return &Fetcher{
		client: newRLHTTPClient(c, limiter),
	}
}

func (f *Fetcher) Fetch(number int, set string, lang string, opts ...fetcher.Opt) (fetcher.Card, error) {
	foil := containsFoil(opts...)

	cardInfoEn, err := f.downloadCard(number, set, "en")
	if err != nil {
		return fetcher.Card{}, err
	}

	var image []byte
	if cardInfoEn.URIs != nil {
		image, err = f.doRequest(cardInfoEn.URIs.PngURI)
		if err != nil {
			return fetcher.Card{}, err
		}
	} else if len(cardInfoEn.Faces) > 0 {
		image, err = f.doRequest(cardInfoEn.Faces[0].URIs.PngURI)
		if err != nil {
			return fetcher.Card{}, err
		}
	} else {
		cardInfoEn, err = f.downloadCard(number, set, lang)
		if err != nil {
			return fetcher.Card{}, err
		}

		if cardInfoEn.URIs != nil {
			image, err = f.doRequest(cardInfoEn.URIs.PngURI)
			if err != nil {
				return fetcher.Card{}, err
			}
		} else {
			image, err = f.doRequest(cardInfoEn.Faces[0].URIs.PngURI)
			if err != nil {
				return fetcher.Card{}, err
			}
		}
	}

	t, err := time.Parse("2006-01-02", cardInfoEn.ReleasedAt)
	if err != nil {
		return fetcher.Card{}, fetcher.NewCardError(number, set, err)
	}

	price := cardInfoEn.Prices.NonFoil
	if foil {
		price = cardInfoEn.Prices.Foil
	}

	cardInfoLang, err := f.downloadCard(number, set, lang)
	if err != nil {
		return fetcher.Card{}, err
	}

	return fetcher.Card{
		ID:         cardInfoLang.ID,
		Name:       cardInfoLang.Name,
		Language:   cardInfoLang.Language,
		URL:        cardInfoLang.Scryfall,
		SetName:    cardInfoLang.SetName,
		Rarity:     cardInfoLang.Rarity,
		Image:      image,
		ManaCost:   cardInfoLang.ManaCost,
		Reprint:    cardInfoLang.Reprint,
		Price:      price,
		ReleasedAt: t,
	}, nil
}

func (f *Fetcher) downloadCard(number int, set string, lang string) (card, error) {
	body, err := f.doRequest(fmt.Sprintf("%s/cards/%s/%d/%s", scryfallUrl, set, number, lang))
	if err != nil {
		return card{}, err
	}

	var c card
	err = json.Unmarshal(body, &c)
	if err != nil {
		return card{}, fetcher.NewCardError(number, set, err)
	}

	return c, nil
}

func (f *Fetcher) doRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pngBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return pngBody, nil
}

func containsFoil(opts ...fetcher.Opt) bool {
	for _, opt := range opts {
		if opt == fetcher.Foil {
			return true
		}
	}
	return false
}
