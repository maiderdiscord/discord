package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
	"h12.io/socks"
)

const discordBaseURL = "https://discord.com"

type ProxyType int

const (
	ProxyTypeHTTP ProxyType = iota
	ProxyTypeSOCKS5
)

type Discord struct {
	client   *http.Client
	Token    string
	Platform Platform
}

func New(token string, platform Platform, proxy string, proxyType ProxyType) (*Discord, error) {
	if !SupportedPlatform(platform) {
		return nil, xerrors.New("platform is not supported")
	}

	client := new(http.Client)

	if proxy != "" {
		switch proxyType {
		default:
		case ProxyTypeHTTP:
			proxyURL, err := url.Parse(fmt.Sprintf("http://%s", proxy))
			if err != nil {
				return nil, xerrors.Errorf("failed to parse URL: %w", err)
			}

			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}

		case ProxyTypeSOCKS5:
			p := socks.Dial(fmt.Sprintf("socks5://%s", proxy))
			client.Transport = &http.Transport{
				Dial: p,
			}
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

	if err := d.Do(ctx, http.MethodGet, "/api/v9/users/@me", nil, data, "", ""); err != nil {
		return nil, xerrors.Errorf("failed to get me: %w", err)
	}

	return data, nil
}

func (d *Discord) AcceptInvite(ctx context.Context, code string) error {
	data := new(GetInviteResponse)

	if err := d.Do(ctx, http.MethodGet, "/api/v9/invites/"+code, nil, data, "", ""); err != nil {
		return err
	}

	if err := d.Do(ctx, http.MethodPost, "/api/v9/invites/"+code, struct{}{}, nil, data.Channel.ID, data.Guild.ID); err != nil {
		return err
	}

	return nil
}

func (d *Discord) LeaveGuild(ctx context.Context, guildID string) error {
	req := struct {
		Lurking bool `json:"lurking"`
	}{
		Lurking: false,
	}
	if err := d.Do(ctx, http.MethodDelete, "/api/v9/users/@me/guilds/"+guildID, req, nil, "", ""); err != nil {
		return err
	}
	return nil
}

func (d *Discord) Do(ctx context.Context, method string, path string, requestBody interface{}, result interface{}, channelID, guildID string) error {
	requestBodyText := make([]byte, 0)

	if requestBody != nil {
		b, err := json.Marshal(requestBody)
		if err != nil {
			return xerrors.Errorf("failed to marshal json")
		}

		requestBodyText = b
	}

	req, err := http.NewRequestWithContext(ctx, method, discordBaseURL+path, bytes.NewBuffer(requestBodyText))
	if err != nil {
		return xerrors.Errorf("failed to create request: %w", err)
	}

	req.Header = GetHeaders(d.Token, d.Platform)

	if len(channelID) > 0 && len(guildID) > 0 {
		props, err := GetContentProperties(channelID, guildID)
		if err != nil {
			return err
		}
		req.Header.Add("X-Content-Properties", props)
		req.Header.Add("Referer", fmt.Sprintf("https://discordapp.com/channels/%s/%s", channelID, guildID))
	}

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
