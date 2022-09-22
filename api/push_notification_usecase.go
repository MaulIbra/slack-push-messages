package api

type IPushNotificationUseCase interface {
	SendMessage(notification *PushNotification) error
}
