package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateNotification(ctx context.Context, input map[string]interface{}) (item *Notification, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateNotification(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateNotificationHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Notification{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Notification",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes NotificationChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["groupID"]; ok && (item.GroupID != changes.GroupID) && (item.GroupID == nil || changes.GroupID == nil || *item.GroupID != *changes.GroupID) {
		item.GroupID = changes.GroupID

		event.AddNewValue("groupID", changes.GroupID)
	}

	if _, ok := input["subject"]; ok && (item.Subject != changes.Subject) && (item.Subject == nil || changes.Subject == nil || *item.Subject != *changes.Subject) {
		item.Subject = changes.Subject

		event.AddNewValue("subject", changes.Subject)
	}

	if _, ok := input["message"]; ok && (item.Message != changes.Message) {
		item.Message = changes.Message

		event.AddNewValue("message", changes.Message)
	}

	if _, ok := input["seen"]; ok && (item.Seen != changes.Seen) {
		item.Seen = changes.Seen

		event.AddNewValue("seen", changes.Seen)
	}

	if _, ok := input["highlighted"]; ok && (item.Highlighted != changes.Highlighted) {
		item.Highlighted = changes.Highlighted

		event.AddNewValue("highlighted", changes.Highlighted)
	}

	if _, ok := input["url"]; ok && (item.URL != changes.URL) && (item.URL == nil || changes.URL == nil || *item.URL != *changes.URL) {
		item.URL = changes.URL

		event.AddNewValue("url", changes.URL)
	}

	if _, ok := input["channel"]; ok && (item.Channel != changes.Channel) && (item.Channel == nil || changes.Channel == nil || *item.Channel != *changes.Channel) {
		item.Channel = changes.Channel

		event.AddNewValue("channel", changes.Channel)
	}

	if _, ok := input["principal"]; ok && (item.Principal != changes.Principal) && (item.Principal == nil || changes.Principal == nil || *item.Principal != *changes.Principal) {
		item.Principal = changes.Principal

		event.AddNewValue("principal", changes.Principal)
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		item.Reference = changes.Reference

		event.AddNewValue("reference", changes.Reference)
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) && (item.ReferenceID == nil || changes.ReferenceID == nil || *item.ReferenceID != *changes.ReferenceID) {
		item.ReferenceID = changes.ReferenceID

		event.AddNewValue("referenceID", changes.ReferenceID)
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		item.Date = changes.Date

		event.AddNewValue("date", changes.Date)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) UpdateNotification(ctx context.Context, id string, input map[string]interface{}) (item *Notification, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateNotification(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateNotificationHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Notification{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Notification",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes NotificationChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["groupID"]; ok && (item.GroupID != changes.GroupID) && (item.GroupID == nil || changes.GroupID == nil || *item.GroupID != *changes.GroupID) {
		event.AddOldValue("groupID", item.GroupID)
		event.AddNewValue("groupID", changes.GroupID)
		item.GroupID = changes.GroupID
	}

	if _, ok := input["subject"]; ok && (item.Subject != changes.Subject) && (item.Subject == nil || changes.Subject == nil || *item.Subject != *changes.Subject) {
		event.AddOldValue("subject", item.Subject)
		event.AddNewValue("subject", changes.Subject)
		item.Subject = changes.Subject
	}

	if _, ok := input["message"]; ok && (item.Message != changes.Message) {
		event.AddOldValue("message", item.Message)
		event.AddNewValue("message", changes.Message)
		item.Message = changes.Message
	}

	if _, ok := input["seen"]; ok && (item.Seen != changes.Seen) {
		event.AddOldValue("seen", item.Seen)
		event.AddNewValue("seen", changes.Seen)
		item.Seen = changes.Seen
	}

	if _, ok := input["highlighted"]; ok && (item.Highlighted != changes.Highlighted) {
		event.AddOldValue("highlighted", item.Highlighted)
		event.AddNewValue("highlighted", changes.Highlighted)
		item.Highlighted = changes.Highlighted
	}

	if _, ok := input["url"]; ok && (item.URL != changes.URL) && (item.URL == nil || changes.URL == nil || *item.URL != *changes.URL) {
		event.AddOldValue("url", item.URL)
		event.AddNewValue("url", changes.URL)
		item.URL = changes.URL
	}

	if _, ok := input["channel"]; ok && (item.Channel != changes.Channel) && (item.Channel == nil || changes.Channel == nil || *item.Channel != *changes.Channel) {
		event.AddOldValue("channel", item.Channel)
		event.AddNewValue("channel", changes.Channel)
		item.Channel = changes.Channel
	}

	if _, ok := input["principal"]; ok && (item.Principal != changes.Principal) && (item.Principal == nil || changes.Principal == nil || *item.Principal != *changes.Principal) {
		event.AddOldValue("principal", item.Principal)
		event.AddNewValue("principal", changes.Principal)
		item.Principal = changes.Principal
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		event.AddOldValue("reference", item.Reference)
		event.AddNewValue("reference", changes.Reference)
		item.Reference = changes.Reference
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) && (item.ReferenceID == nil || changes.ReferenceID == nil || *item.ReferenceID != *changes.ReferenceID) {
		event.AddOldValue("referenceID", item.ReferenceID)
		event.AddNewValue("referenceID", changes.ReferenceID)
		item.ReferenceID = changes.ReferenceID
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		event.AddOldValue("date", item.Date)
		event.AddNewValue("date", changes.Date)
		item.Date = changes.Date
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteNotification(ctx context.Context, id string) (item *Notification, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteNotification(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteNotificationHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Notification{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Notification",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("notifications")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllNotifications(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllNotifications(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllNotificationsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Notification{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
