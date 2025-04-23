package uspsgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

var tokenEndpoint = endpointBase + "/oauth2/v3/token"

type Client struct {
	client *http.Client
}

func New(key, secret string) Client {
	return NewWithContext(context.Background(), key, secret)
}

func NewWithContext(ctx context.Context, key, secret string) Client {
	conf := &clientcredentials.Config{
		ClientID:     key,
		ClientSecret: secret,
		TokenURL:     tokenEndpoint,
	}

	return Client{client: conf.Client(context.Background())}
}

func (c Client) request(ctx context.Context, endpoint string, params any, resp any) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?%s", endpoint, toParams(params).Encode()), nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json")

	r, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		var errResp ErrorDetails
		if err := json.NewDecoder(r.Body).Decode(&errResp); err != nil {
			return err
		}
		return Error{err: errResp}
	}
	return json.NewDecoder(r.Body).Decode(resp)
}
