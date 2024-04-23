package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func HandleSlashCommandEvent(event socketmode.Event, client *slack.Client) error {
	log.Print("-------------------------------------------------------")
	log.Printf("Event type is %s", strings.ToLower(string(event.Type)))
	log.Print("-------------------------------------------------------")

	data := event.Data.(slack.SlashCommand)
	_, _, err := client.PostMessage(data.ChannelID, slack.MsgOptionText(fmt.Sprintf("Hi, I'm Test Bot I got your command."), false))
	if err != nil {
		return err
	}
	return nil

}
