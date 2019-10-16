package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryNotificationHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *NotificationFilterType
}

func (r *GeneratedQueryResolver) Notification(ctx context.Context, id *string, q *string, filter *NotificationFilterType) (*Notification, error) {
	opts := QueryNotificationHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryNotification(ctx, r.GeneratedResolver, opts)
}
func QueryNotificationHandler(ctx context.Context, r *GeneratedResolver, opts QueryNotificationHandlerOptions) (*Notification, error) {
	query := NotificationQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &NotificationResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("notifications")+".id = ?", *opts.ID)
	}

	var items []*Notification
	giOpts := GetItemsOptions{
		Alias:      TableName("notifications"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Notification"}
	}
	return items[0], err
}

type QueryNotificationsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*NotificationSortType
	Filter *NotificationFilterType
}

func (r *GeneratedQueryResolver) Notifications(ctx context.Context, offset *int, limit *int, q *string, sort []*NotificationSortType, filter *NotificationFilterType) (*NotificationResultType, error) {
	opts := QueryNotificationsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryNotifications(ctx, r.GeneratedResolver, opts)
}
func QueryNotificationsHandler(ctx context.Context, r *GeneratedResolver, opts QueryNotificationsHandlerOptions) (*NotificationResultType, error) {
	query := NotificationQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &NotificationResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedNotificationResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedNotificationResultTypeResolver) Items(ctx context.Context, obj *NotificationResultType) (items []*Notification, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("notifications"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedNotificationResultTypeResolver) Count(ctx context.Context, obj *NotificationResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Notification{})
}
