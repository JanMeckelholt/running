package strava

import (
	"context"
	"encoding/json"

	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Client struct {
	stravaURL url.URL
}

type tokenResp struct {
	token string
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
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		// refreshTokenResp, err := c.refreshToken(clientId, clientSecret, refreshToken)
		if err != nil {
			return nil, err
		}
		//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", refreshTokenResp.Refresh_token))
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) GetActivities(ctx context.Context, token string, since int64) (*http.Response, error) {
	log.Infof("Token: %s", token)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/athlete/activities?after=%d", c.stravaURL.Host, since), nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		//refreshTokenResp, err := c.refreshToken(clientId, clientSecret, refreshToken)
		if err != nil {
			return nil, err
		}
		//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", refreshTokenResp.Refresh_token))
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (c *Client) UseRefreshToken(ctx context.Context, clientId, clientSecret, refreshToken string) (string, error) {
	log.Infof("Refreshing: %s", clientId)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oauth/token?client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", c.stravaURL.Host, clientId, clientSecret, refreshToken), nil)
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)

	tokenResp := &tokenResp{}
	err = json.NewDecoder(resp.Body).Decode(tokenResp)
	return tokenResp.token, err
}
