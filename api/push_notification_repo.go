package api

type IPushNotificationRepo interface {
	SendMessages(messages string) error
}
