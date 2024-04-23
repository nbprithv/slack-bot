package handlers

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// HandleAppMentionEventToBot is used to take care of the AppMentionEvent when the bot is mentioned
func HandleAppMentionEventToBot(event *slackevents.AppMentionEvent, client *slack.Client) error {

	// Grab the user name based on the ID of the one who mentioned the bot
	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}
	// Check if the user said Hello to the bot
	text := strings.ToLower(event.Text)

	// Create the attachment and assigned based on the message
	attachment := slack.Attachment{}
	// Add Some default context like user who mentioned the bot
	// attachment.Fields = []slack.AttachmentField{
	// 	{
	// 		Title: "Date",
	// 		Value: time.Now().String(),
	// 	}, {
	// 		Title: "Initializer",
	// 		Value: user.Name,
	// 	},
	// }
	if strings.Contains(text, "hello") || strings.Contains(text, "hi") {
		// Greet the user
		attachment.Text = fmt.Sprintf("Hello %s", user.Name)
		// attachment.Pretext = "Greetings"
		attachment.Color = "#4af030"
	} else if strings.Contains(text, "weather") {
		// Send a message to the user
		attachment.Text = fmt.Sprintf("Weather is sunny today. %s", user.Name)
		// attachment.Pretext = "How can I be of service"
		attachment.Color = "#4af030"
	} else {
		// Send a message to the user
		attachment.Text = fmt.Sprintf("I am good. How are you %s?", user.Name)
		// attachment.Pretext = "How can I be of service"
		attachment.Color = "#4af030"
	}
	// Send the message to the channel
	// The Channel is available in the event message
	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}
