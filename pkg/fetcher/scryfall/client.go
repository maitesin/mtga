package scryfall

import (
	"context"
	"net/http"

	"golang.org/x/time/rate"
)

//RLHTTPClient Rate Limited HTTP Client
type RLHTTPClient struct {
	client  *http.Client
	limiter *rate.Limiter
}

//NewClient return http client with a ratelimiter
func newRLHTTPClient(c *http.Client, rl *rate.Limiter) *RLHTTPClient {
	return &RLHTTPClient{
		client:  c,
		limiter: rl,
	}
}

//Do dispatches the HTTP request to the network
func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Comment out the below 5 lines to turn off ratelimiting
	ctx := context.Background()
	err := c.limiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
