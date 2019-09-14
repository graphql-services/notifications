package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *NotificationFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *NotificationFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("notifications"), wheres, values, joins)
}
func (f *NotificationFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *NotificationFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Message != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" = ?")
		values = append(values, f.Message)
	}
	if f.MessageNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" != ?")
		values = append(values, f.MessageNe)
	}
	if f.MessageGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" > ?")
		values = append(values, f.MessageGt)
	}
	if f.MessageLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" < ?")
		values = append(values, f.MessageLt)
	}
	if f.MessageGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" >= ?")
		values = append(values, f.MessageGte)
	}
	if f.MessageLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" <= ?")
		values = append(values, f.MessageLte)
	}
	if f.MessageIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" IN (?)")
		values = append(values, f.MessageIn)
	}
	if f.MessageLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MessageLike, "?", "_", -1), "*", "%", -1))
	}
	if f.MessagePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MessagePrefix))
	}
	if f.MessageSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MessageSuffix))
	}

	if f.Seen != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" = ?")
		values = append(values, f.Seen)
	}
	if f.SeenNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" != ?")
		values = append(values, f.SeenNe)
	}
	if f.SeenGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" > ?")
		values = append(values, f.SeenGt)
	}
	if f.SeenLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" < ?")
		values = append(values, f.SeenLt)
	}
	if f.SeenGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" >= ?")
		values = append(values, f.SeenGte)
	}
	if f.SeenLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" <= ?")
		values = append(values, f.SeenLte)
	}
	if f.SeenIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" IN (?)")
		values = append(values, f.SeenIn)
	}

	if f.Channel != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" = ?")
		values = append(values, f.Channel)
	}
	if f.ChannelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" != ?")
		values = append(values, f.ChannelNe)
	}
	if f.ChannelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" > ?")
		values = append(values, f.ChannelGt)
	}
	if f.ChannelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" < ?")
		values = append(values, f.ChannelLt)
	}
	if f.ChannelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" >= ?")
		values = append(values, f.ChannelGte)
	}
	if f.ChannelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" <= ?")
		values = append(values, f.ChannelLte)
	}
	if f.ChannelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" IN (?)")
		values = append(values, f.ChannelIn)
	}
	if f.ChannelLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ChannelLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ChannelPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ChannelPrefix))
	}
	if f.ChannelSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ChannelSuffix))
	}

	if f.Principal != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" = ?")
		values = append(values, f.Principal)
	}
	if f.PrincipalNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" != ?")
		values = append(values, f.PrincipalNe)
	}
	if f.PrincipalGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" > ?")
		values = append(values, f.PrincipalGt)
	}
	if f.PrincipalLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" < ?")
		values = append(values, f.PrincipalLt)
	}
	if f.PrincipalGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" >= ?")
		values = append(values, f.PrincipalGte)
	}
	if f.PrincipalLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" <= ?")
		values = append(values, f.PrincipalLte)
	}
	if f.PrincipalIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" IN (?)")
		values = append(values, f.PrincipalIn)
	}
	if f.PrincipalLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalLike, "?", "_", -1), "*", "%", -1))
	}
	if f.PrincipalPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalPrefix))
	}
	if f.PrincipalSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalSuffix))
	}

	if f.Reference != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" = ?")
		values = append(values, f.Reference)
	}
	if f.ReferenceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" != ?")
		values = append(values, f.ReferenceNe)
	}
	if f.ReferenceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" > ?")
		values = append(values, f.ReferenceGt)
	}
	if f.ReferenceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" < ?")
		values = append(values, f.ReferenceLt)
	}
	if f.ReferenceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" >= ?")
		values = append(values, f.ReferenceGte)
	}
	if f.ReferenceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" <= ?")
		values = append(values, f.ReferenceLte)
	}
	if f.ReferenceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IN (?)")
		values = append(values, f.ReferenceIn)
	}
	if f.ReferenceLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ReferencePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferencePrefix))
	}
	if f.ReferenceSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceSuffix))
	}

	if f.ReferenceID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" = ?")
		values = append(values, f.ReferenceID)
	}
	if f.ReferenceIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" != ?")
		values = append(values, f.ReferenceIDNe)
	}
	if f.ReferenceIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" > ?")
		values = append(values, f.ReferenceIDGt)
	}
	if f.ReferenceIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" < ?")
		values = append(values, f.ReferenceIDLt)
	}
	if f.ReferenceIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" >= ?")
		values = append(values, f.ReferenceIDGte)
	}
	if f.ReferenceIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" <= ?")
		values = append(values, f.ReferenceIDLte)
	}
	if f.ReferenceIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IN (?)")
		values = append(values, f.ReferenceIDIn)
	}
	if f.ReferenceIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceIDLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ReferenceIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceIDPrefix))
	}
	if f.ReferenceIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceIDSuffix))
	}

	if f.Date != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" = ?")
		values = append(values, f.Date)
	}
	if f.DateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" != ?")
		values = append(values, f.DateNe)
	}
	if f.DateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" > ?")
		values = append(values, f.DateGt)
	}
	if f.DateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" < ?")
		values = append(values, f.DateLt)
	}
	if f.DateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" >= ?")
		values = append(values, f.DateGte)
	}
	if f.DateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" <= ?")
		values = append(values, f.DateLte)
	}
	if f.DateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IN (?)")
		values = append(values, f.DateIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *NotificationFilterType) AndWith(f2 ...*NotificationFilterType) *NotificationFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &NotificationFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *NotificationFilterType) OrWith(f2 ...*NotificationFilterType) *NotificationFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &NotificationFilterType{
		Or: append(_f2, f),
	}
}
