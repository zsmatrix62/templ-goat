package gotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/zsmatrix62/templ-goat/pkg/config"
	"github.com/zsmatrix62/templ-goat/pkg/logger"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

var Client *client

// Config stores the configuration for the Gotify client
type Config struct {
	Url string
	Key string
}

// Client represents the Gotify client
type client struct {
	Config Config
}

// NewClient initializes a new Gotify client
func NewClient(url, token string) *client {
	return &client{
		Config: Config{
			Url: url,
			Key: token,
		},
	}
}

func DefaultClient() *client {
	return NewClient(config.GetString("gotify.url"), config.GetString("gotify.token"))
}

// SendMessage sends a message to the Gotify server
func (c *client) SendMessage(title string, priority int, message ...any) {
	var err error

	defer func() {
		logger.LogWarnIf(err)
	}()

	url := fmt.Sprintf("%s/message?token=%s", c.Config.Url, c.Config.Key)

	messageStrings := lo.Map(message, func(m any, idx int) string {
		return cast.ToString(m)
	})

	// Create the message payload
	data := map[string]interface{}{
		"message":  strings.Join(messageStrings, ", "),
		"title":    title,
		"priority": priority,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("failed to marshal message: %v", err)
		return
	}

	// Send the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		err = fmt.Errorf("failed to create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response is OK
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("gotify returned non-OK status: %v", resp.Status)
		return
	}
}
