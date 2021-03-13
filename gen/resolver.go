package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateNotification     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Notification, err error)
	UpdateNotification     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Notification, err error)
	DeleteNotification     func(ctx context.Context, r *GeneratedResolver, id string) (item *Notification, err error)
	DeleteAllNotifications func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryNotification      func(ctx context.Context, r *GeneratedResolver, opts QueryNotificationHandlerOptions) (*Notification, error)
	QueryNotifications     func(ctx context.Context, r *GeneratedResolver, opts QueryNotificationsHandlerOptions) (*NotificationResultType, error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

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
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
