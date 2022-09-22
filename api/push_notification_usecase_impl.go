package api

import (
	"github.com/slack-go/slack"
	"os"
)

type pushNotificationUseCase struct {
}

func (p pushNotificationUseCase) SendMessage(notification *PushNotification) error {

	if len(os.Args) > 1 && os.Args[1] == "local" {
		notification.SlackToken = os.Getenv("SLACK_TOKEN")
		notification.SlackChannelId = os.Getenv("SLACK_CHANNEL_ID")
	}
	slackToken := slack.New(notification.SlackToken)
	pushNotificationRepo := NewPushNotificationRepo(slackToken, notification.SlackChannelId)
	return pushNotificationRepo.SendMessages(notification.Message)
}

func NewPushNotificationUseCase() IPushNotificationUseCase {
	return &pushNotificationUseCase{}
}
