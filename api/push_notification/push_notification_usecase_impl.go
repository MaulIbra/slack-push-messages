package push_notification

import (
	"github.com/slack-go/slack"
)

type pushNotificationUseCase struct {
}

func (p pushNotificationUseCase) SendMessage(notification *PushNotification) error {
	slackToken := slack.New(notification.SlackToken)
	pushNotificationRepo := NewPushNotificationRepo(slackToken, notification.SlackChannelId)
	return pushNotificationRepo.SendMessages(notification.Message)
}

func NewPushNotificationUseCase() IPushNotificationUseCase {
	return &pushNotificationUseCase{}
}
