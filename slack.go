package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	sl "github.com/slack-go/slack"
)

// SlClient obj
type SlClient struct {
	WebhookURL string
	DryRun     bool
}

// NewClient creates a new Slack client
func NewClient(webhookURL string) SlClient {
	return SlClient{
		WebhookURL: webhookURL,
	}
}

// Report sends a message to the Slack channel
func (slack SlClient) Report(msg string) error {

	if slack.DryRun {
		fmt.Println(msg)
		return nil
	}
	payload := sl.WebhookMessage{
		Text: msg,
		// LinkNames:   true,
		// Mrkdwn:      true,
		// Username:    *userName,
		// IconEmoji:   *iconEmoji,
		// Channel:     *channel,
		// Attachments: []sl.Attachment{attachment},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(data)

	request, err := http.NewRequest("POST", slack.WebhookURL, body)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

