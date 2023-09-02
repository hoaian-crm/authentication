package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

type Message struct {
	Code        int               `json:"code"`
	Description string            `json:"description"`
	Field       string            `json:"field"`
	MetaData    map[string]string `json:"meta_data"`
}

func MessagesBuilder(err error) []Message {
	var messages = []Message{}
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				message := Message{
					Field:       e.Field(),
					Code:        Messages[e.Tag()].Code,
					Description: Messages[e.Tag()].Description,
				}
				if e.Param() != "" {
					message.MetaData = map[string]string{
						e.Tag(): e.Param(),
					}
				}
				messages = append(messages, message)
			}
		}
	}
	return messages
}

var Messages = map[string]Message{
	"is_not_empty": {
		Code:        1,
		Description: "not allow empty",
	},
	"is_string": {
		Code:        2,
		Description: "must be text",
	},
	"email": {
		Code:        3,
		Description: "must be an email address",
	},
	"min_length": {
		Code:        4,
		Description: "must be have length not less than $min_length",
	},
	"must_unique": {
		Code:        5,
		Description: "must be unique",
	},
	"invalid_access_token": {
		Code:        6,
		Description: "invalid access token or access token expired",
	},
	"unauthorization": {
		Code:        7,
		Description: "unauthorization",
	},
	"login_success": {
		Code:        8,
		Description: "login success",
	},
	"invalid_email_password": {
		Code:        9,
		Description: "invalid email or password",
	},
	"missing_token": {
		Code:        10,
		Description: "missing authorization token (Bearer <token>)",
	},
	"get_success": {
		Code:        11,
		Description: "get success",
	},
	"invalid_otp_code": {
		Code:        12,
		Description: "otp code invalid",
	},
	"email_not_registered": {
		Code:        13,
		Description: "email is not in database",
	},
	"update_success": {
		Code:        14,
		Description: "update successfuly",
	},
	"is_number": {
		Code:        15,
		Description: "must be a number",
	},
	"unknown_model": {
		Code:        16,
		Description: "unknown model to implement",
	},
	"no_filter": {
		Code:        17,
		Description: "not have filter",
	},
	"no_data": {
		Code:        18,
		Description: "not have data body",
	},
	"update_failed": {
		Code:        19,
		Description: "update failed",
	},
	"delete_success": {
		Code:        20,
		Description: "delete success",
	},
	"delete_failed": {
		Code:        21,
		Description: "delete failed",
	},
	"get_failed": {
		Code:        22,
		Description: "get failed",
	},
	"create_failed": {
		Code:        23,
		Description: "create failed",
	},
	"create_success": {
		Code:        24,
		Description: "create success",
	},
	"invalid_body_data": {
		Code:        25,
		Description: "request send invalid body data",
	},
	"not_found_channel": {
		Code:        26,
		Description: "not found channel",
	},
}
