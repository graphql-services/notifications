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
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := NotificationQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &NotificationResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
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
		return nil, nil
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
func (r *GeneratedResolver) NotificationsItems(ctx context.Context, opts QueryNotificationsHandlerOptions) (res []*Notification, err error) {
	resultType, err := r.Handlers.QueryNotifications(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("notifications"),
	}, &res)
	if err != nil {
		return
	}
	return
}
func (r *GeneratedResolver) NotificationsCount(ctx context.Context, opts QueryNotificationsHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryNotifications(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("notifications"),
	}, &Notification{})
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
	otps := GetItemsOptions{
		Alias:      TableName("notifications"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Notification{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedNotificationResultTypeResolver) Count(ctx context.Context, obj *NotificationResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("notifications"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Notification{})
}
