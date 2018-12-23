package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	errRequestFailed = errors.New("request failed")
)

type Bot struct {
	ID    string
	Token string

	httpClient *http.Client
}

func NewBot(id, token string) *Bot {
	return &Bot{
		ID:    strings.TrimPrefix(id, "bot"),
		Token: token,

		httpClient: &http.Client{},
	}
}

func (b Bot) apiURL(action string) *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprintf("bot%s:%s/%s", b.ID, b.Token, action),
	}
}

type ChatMessage struct {
	ChatID string
	Text   string
}

func (b Bot) SendMessage(m ChatMessage) error {
	url := b.apiURL("sendMessage")
	qs := url.Query()
	qs.Set("chat_id", m.ChatID)
	qs.Set("text", m.Text)
	url.RawQuery = qs.Encode()

	resp, err := b.httpClient.Post(url.String(), "", nil)
	if err != nil {
		return err
	}

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	rv := map[string]interface{}{}
	if err := json.Unmarshal(bodyRaw, &rv); err != nil {
		return err
	}

	if ok, yes := rv["ok"].(bool); yes {
		if ok {
			return nil
		}
	}

	return errors.New(string(bodyRaw))
}
