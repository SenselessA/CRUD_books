package handler

import "github.com/sirupsen/logrus"

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(statusCode int, message string) {
		logrus.Error(message)
}