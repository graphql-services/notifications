package src

import (
	"context"

	"github.com/graphql-services/notifications/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }

	return resolver
}

func (r *MutationResolver) SeenNotification(ctx context.Context, id string) (*gen.Notification, error) {
	input := map[string]interface{}{
		"seen": true,
	}
	return r.Handlers.UpdateNotification(ctx, r.GeneratedResolver, id, input)
}

func (r *MutationResolver) SeenNotifications(ctx context.Context, principal string, channel *string, reference *string, referenceID *string) (ok bool, err error) {
	input := map[string]interface{}{
		"seen": true,
	}
	where := gen.Notification{
		Channel:     channel,
		Principal:   &principal,
		Reference:   reference,
		ReferenceID: referenceID,
	}
	err = r.DB.Query().Model(where).Where(where).Updates(input).Error

	ok = true
	return
}
