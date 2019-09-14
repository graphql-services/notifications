package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	CreateNotification     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Notification, err error)
	UpdateNotification     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Notification, err error)
	DeleteNotification     func(ctx context.Context, r *GeneratedResolver, id string) (item *Notification, err error)
	DeleteAllNotifications func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryNotification      func(ctx context.Context, r *GeneratedResolver, opts QueryNotificationHandlerOptions) (*Notification, error)
	QueryNotifications     func(ctx context.Context, r *GeneratedResolver, opts QueryNotificationsHandlerOptions) (*NotificationResultType, error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{

		CreateNotification:     CreateNotificationHandler,
		UpdateNotification:     UpdateNotificationHandler,
		DeleteNotification:     DeleteNotificationHandler,
		DeleteAllNotifications: DeleteAllNotificationsHandler,
		QueryNotification:      QueryNotificationHandler,
		QueryNotifications:     QueryNotificationsHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
