package strava

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	stravaURL url.URL
	token     string
}

func NewClient(starvaURL url.URL, token string) *Client {
	return &Client{
		stravaURL: starvaURL,
		token:     token,
	}
}

func (c *Client) GetAthlet(ctx context.Context) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint("%s/athlete", c.stravaURL), nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}
