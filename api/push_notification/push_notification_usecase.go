package push_notification

type IPushNotificationUseCase interface {
	SendMessage(notification *PushNotification) error
}
