package fcm

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var APP *firebase.App

func New() (err error) {
	opt := option.WithCredentialsFile("../private/fcm.json")
	APP, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}

	return
}
