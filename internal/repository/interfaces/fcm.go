package interfaces

import (
	"context"

	"firebase.google.com/go/messaging"
)

type FCM interface {
	SendMessage(context.Context, *messaging.MulticastMessage) error
}
