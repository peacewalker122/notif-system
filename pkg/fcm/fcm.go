package fcm

import (
	"context"
	"log/slog"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var APP *firebase.App

func New(path string) (err error) {
	opt := option.WithCredentialsFile(path)
	APP, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}

	slog.Info("Firebase initialized", APP)

	return
}
