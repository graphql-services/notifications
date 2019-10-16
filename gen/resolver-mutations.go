package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateNotification(ctx context.Context, input map[string]interface{}) (item *Notification, err error) {
	return r.Handlers.CreateNotification(ctx, r.GeneratedResolver, input)
}
func CreateNotificationHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Notification{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

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
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["message"]; ok && (item.Message != changes.Message) {
		item.Message = changes.Message
		event.AddNewValue("message", changes.Message)
	}

	if _, ok := input["seen"]; ok && (item.Seen != changes.Seen) {
		item.Seen = changes.Seen
		event.AddNewValue("seen", changes.Seen)
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

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateNotification(ctx context.Context, id string, input map[string]interface{}) (item *Notification, err error) {
	return r.Handlers.UpdateNotification(ctx, r.GeneratedResolver, id, input)
}
func UpdateNotificationHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Notification{}
	now := time.Now()
	tx := r.DB.db.Begin()

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
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

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

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteNotification(ctx context.Context, id string) (item *Notification, err error) {
	return r.Handlers.DeleteNotification(ctx, r.GeneratedResolver, id)
}
func DeleteNotificationHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Notification, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Notification{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
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
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllNotifications(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllNotifications(ctx, r.GeneratedResolver)
}
func DeleteAllNotificationsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Notification{}).Error
	return err == nil, err
}
