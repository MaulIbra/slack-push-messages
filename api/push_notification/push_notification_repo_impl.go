package push_notification

import "github.com/slack-go/slack"

type pushNotificationRepo struct {
	slackClient  *slack.Client
	slackChannel string
}

func (p pushNotificationRepo) SendMessages(messages string) error {
	_, _, err := p.slackClient.PostMessage(
		p.slackChannel,
		slack.MsgOptionText(messages, false),
	)
	return err
}

func NewPushNotificationRepo(slackClient *slack.Client, slackChannel string) IPushNotificationRepo {
	return &pushNotificationRepo{slackClient: slackClient, slackChannel: slackChannel}
}
