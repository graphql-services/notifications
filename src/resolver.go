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

func (r *MutationResolver) CreateNotificationBatchUpdate(ctx context.Context, input gen.NotificationBatchUpdateCreateInput) (res *gen.NotificationBatchUpdate, err error) {
	q := r.GetDB(ctx).Model(&gen.Notification{}).Where("principal = ?", input.Principal)
	if input.Channel != nil {
		q = q.Where("channel = ?", *input.Channel)
	}
	if input.Reference != nil {
		q = q.Where("reference = ?", *input.Reference)
	}
	if input.ReferenceID != nil {
		q = q.Where("referenceID = ?", *input.ReferenceID)
	}

	err = q.Update("seen", input.Seen).Error
	if err != nil {
		return
	}
	res = &gen.NotificationBatchUpdate{ID: "done"}
	return
}
