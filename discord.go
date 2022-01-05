package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
)

const discordBaseURL = "https://discord.com"

type Discord struct {
	client   *http.Client
	Token    string
	Platform Platform
}

func New(token string, platform Platform, proxyURL *url.URL) (*Discord, error) {
	if !SupportedPlatform(platform) {
		return nil, xerrors.New("platform is not supported")
	}

	client := new(http.Client)

	if proxyURL != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	return &Discord{
		client:   client,
		Token:    token,
		Platform: platform,
	}, nil
}

func (d *Discord) Me(ctx context.Context) (*MeResponse, error) {
	data := new(MeResponse)

	if err := d.Do(ctx, http.MethodGet, "/api/v9/users/@me", nil, data); err != nil {
		return nil, xerrors.Errorf("failed to get me: %w", err)
	}

	return data, nil
}

func (d *Discord) Do(ctx context.Context, method string, path string, requestBody interface{}, result interface{}) error {
	requestBodyText := make([]byte, 0)

	if requestBody != nil {
		b, err := json.Marshal(requestBody)
		if err != nil {
			return xerrors.Errorf("failed to marshal json")
		}

		requestBodyText = b
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, discordBaseURL+path, bytes.NewBuffer(requestBodyText))
	if err != nil {
		return xerrors.Errorf("failed to create request: %w", err)
	}

	req.Header = GetHeaders(d.Token, d.Platform)

	res, err := d.client.Do(req)
	if err != nil {
		return xerrors.Errorf("failed to send request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return xerrors.Errorf("status code is invalid: %d", res.StatusCode)
	}

	if result != nil {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return xerrors.Errorf("failed to read body: %w", err)
		}

		if err := json.Unmarshal(body, result); err != nil {
			return xerrors.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	return nil
}
