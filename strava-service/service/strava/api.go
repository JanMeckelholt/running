package strava

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Client struct {
	stravaURL url.URL
}

func NewClient(starvaURL url.URL) *Client {
	return &Client{
		stravaURL: starvaURL,
	}
}

func (c *Client) GetAthlet(ctx context.Context, token string) (*http.Response, error) {
	log.Infof("Token: %s", token)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/athlete", c.stravaURL.Host), nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}
