package src

import (
	"context"

	"github.com/graphql-services/notifications/gen"
)

func New(db *gen.DB, ec *gen.EventController) *Resolver {
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

func (r *MutationResolver) CreateNotificationBatchUpdate(ctx context.Context, input gen.NotificationBatchUpdateCreateInput) (res bool, err error) {

	filter := &gen.NotificationFilterType{
		Principal: &input.Principal,
	}

	if input.Channel != nil {
		filter.Channel = input.Channel
	}
	if input.Reference != nil {
		filter.Reference = input.Reference
	}
	if input.ReferenceID != nil {
		filter.ReferenceID = input.ReferenceID
	}

	limit := 500
	items, err := r.GeneratedResolver.NotificationsItems(ctx, gen.QueryNotificationsHandlerOptions{
		Limit:  &limit,
		Filter: filter,
	})

	db := r.GetDB(ctx)
	for _, item := range items {
		item.Seen = input.Seen
		err = db.Save(item).Error
		if err != nil {
			return
		}
	}
	res = true
	return
}
