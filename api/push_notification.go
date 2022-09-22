package api

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type PushNotification struct {
	SlackToken     string `validate:"required" json:"slack_token"`
	SlackChannelId string `validate:"required" json:"slack_channel_id"`
	Message        string `validate:"required" json:"message"`
}

type PushNotificationResponse struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (pn PushNotification) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(pn)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
