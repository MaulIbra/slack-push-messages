package push_notification

type IPushNotificationRepo interface {
	SendMessages(messages string) error
}
