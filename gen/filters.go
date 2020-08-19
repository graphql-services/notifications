package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *NotificationFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *NotificationFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("notifications"), wheres, whereValues, havings, havingValues, joins)
}
func (f *NotificationFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
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

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.GroupID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" = ?")
		values = append(values, f.GroupID)
	}

	if f.GroupIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" != ?")
		values = append(values, f.GroupIDNe)
	}

	if f.GroupIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" > ?")
		values = append(values, f.GroupIDGt)
	}

	if f.GroupIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" < ?")
		values = append(values, f.GroupIDLt)
	}

	if f.GroupIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" >= ?")
		values = append(values, f.GroupIDGte)
	}

	if f.GroupIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" <= ?")
		values = append(values, f.GroupIDLte)
	}

	if f.GroupIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" IN (?)")
		values = append(values, f.GroupIDIn)
	}

	if f.GroupIDNull != nil {
		if *f.GroupIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("groupID")+" IS NOT NULL")
		}
	}

	if f.Subject != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" = ?")
		values = append(values, f.Subject)
	}

	if f.SubjectNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" != ?")
		values = append(values, f.SubjectNe)
	}

	if f.SubjectGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" > ?")
		values = append(values, f.SubjectGt)
	}

	if f.SubjectLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" < ?")
		values = append(values, f.SubjectLt)
	}

	if f.SubjectGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" >= ?")
		values = append(values, f.SubjectGte)
	}

	if f.SubjectLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" <= ?")
		values = append(values, f.SubjectLte)
	}

	if f.SubjectIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" IN (?)")
		values = append(values, f.SubjectIn)
	}

	if f.SubjectLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SubjectLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SubjectPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SubjectPrefix))
	}

	if f.SubjectSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SubjectSuffix))
	}

	if f.SubjectNull != nil {
		if *f.SubjectNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("subject")+" IS NOT NULL")
		}
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

	if f.MessageNull != nil {
		if *f.MessageNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("message")+" IS NOT NULL")
		}
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

	if f.SeenNull != nil {
		if *f.SeenNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("seen")+" IS NOT NULL")
		}
	}

	if f.URL != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" = ?")
		values = append(values, f.URL)
	}

	if f.URLNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" != ?")
		values = append(values, f.URLNe)
	}

	if f.URLGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" > ?")
		values = append(values, f.URLGt)
	}

	if f.URLLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" < ?")
		values = append(values, f.URLLt)
	}

	if f.URLGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" >= ?")
		values = append(values, f.URLGte)
	}

	if f.URLLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" <= ?")
		values = append(values, f.URLLte)
	}

	if f.URLIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IN (?)")
		values = append(values, f.URLIn)
	}

	if f.URLLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLPrefix))
	}

	if f.URLSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLSuffix))
	}

	if f.URLNull != nil {
		if *f.URLNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IS NOT NULL")
		}
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

	if f.ChannelNull != nil {
		if *f.ChannelNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("channel")+" IS NOT NULL")
		}
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

	if f.PrincipalNull != nil {
		if *f.PrincipalNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("principal")+" IS NOT NULL")
		}
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

	if f.ReferenceNull != nil {
		if *f.ReferenceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NOT NULL")
		}
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

	if f.ReferenceIDNull != nil {
		if *f.ReferenceIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NOT NULL")
		}
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

	if f.DateNull != nil {
		if *f.DateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NOT NULL")
		}
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

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
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

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
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

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
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

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}
func (f *NotificationFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.GroupIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") = ?")
		values = append(values, f.GroupIDMin)
	}

	if f.GroupIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") = ?")
		values = append(values, f.GroupIDMax)
	}

	if f.GroupIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") != ?")
		values = append(values, f.GroupIDMinNe)
	}

	if f.GroupIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") != ?")
		values = append(values, f.GroupIDMaxNe)
	}

	if f.GroupIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") > ?")
		values = append(values, f.GroupIDMinGt)
	}

	if f.GroupIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") > ?")
		values = append(values, f.GroupIDMaxGt)
	}

	if f.GroupIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") < ?")
		values = append(values, f.GroupIDMinLt)
	}

	if f.GroupIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") < ?")
		values = append(values, f.GroupIDMaxLt)
	}

	if f.GroupIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") >= ?")
		values = append(values, f.GroupIDMinGte)
	}

	if f.GroupIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") >= ?")
		values = append(values, f.GroupIDMaxGte)
	}

	if f.GroupIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") <= ?")
		values = append(values, f.GroupIDMinLte)
	}

	if f.GroupIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") <= ?")
		values = append(values, f.GroupIDMaxLte)
	}

	if f.GroupIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("groupID")+") IN (?)")
		values = append(values, f.GroupIDMinIn)
	}

	if f.GroupIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("groupID")+") IN (?)")
		values = append(values, f.GroupIDMaxIn)
	}

	if f.SubjectMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") = ?")
		values = append(values, f.SubjectMin)
	}

	if f.SubjectMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") = ?")
		values = append(values, f.SubjectMax)
	}

	if f.SubjectMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") != ?")
		values = append(values, f.SubjectMinNe)
	}

	if f.SubjectMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") != ?")
		values = append(values, f.SubjectMaxNe)
	}

	if f.SubjectMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") > ?")
		values = append(values, f.SubjectMinGt)
	}

	if f.SubjectMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") > ?")
		values = append(values, f.SubjectMaxGt)
	}

	if f.SubjectMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") < ?")
		values = append(values, f.SubjectMinLt)
	}

	if f.SubjectMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") < ?")
		values = append(values, f.SubjectMaxLt)
	}

	if f.SubjectMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") >= ?")
		values = append(values, f.SubjectMinGte)
	}

	if f.SubjectMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") >= ?")
		values = append(values, f.SubjectMaxGte)
	}

	if f.SubjectMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") <= ?")
		values = append(values, f.SubjectMinLte)
	}

	if f.SubjectMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") <= ?")
		values = append(values, f.SubjectMaxLte)
	}

	if f.SubjectMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") IN (?)")
		values = append(values, f.SubjectMinIn)
	}

	if f.SubjectMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") IN (?)")
		values = append(values, f.SubjectMaxIn)
	}

	if f.SubjectMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SubjectMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SubjectMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.SubjectMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.SubjectMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SubjectMinPrefix))
	}

	if f.SubjectMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.SubjectMaxPrefix))
	}

	if f.SubjectMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SubjectMinSuffix))
	}

	if f.SubjectMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("subject")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.SubjectMaxSuffix))
	}

	if f.MessageMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") = ?")
		values = append(values, f.MessageMin)
	}

	if f.MessageMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") = ?")
		values = append(values, f.MessageMax)
	}

	if f.MessageMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") != ?")
		values = append(values, f.MessageMinNe)
	}

	if f.MessageMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") != ?")
		values = append(values, f.MessageMaxNe)
	}

	if f.MessageMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") > ?")
		values = append(values, f.MessageMinGt)
	}

	if f.MessageMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") > ?")
		values = append(values, f.MessageMaxGt)
	}

	if f.MessageMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") < ?")
		values = append(values, f.MessageMinLt)
	}

	if f.MessageMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") < ?")
		values = append(values, f.MessageMaxLt)
	}

	if f.MessageMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") >= ?")
		values = append(values, f.MessageMinGte)
	}

	if f.MessageMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") >= ?")
		values = append(values, f.MessageMaxGte)
	}

	if f.MessageMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") <= ?")
		values = append(values, f.MessageMinLte)
	}

	if f.MessageMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") <= ?")
		values = append(values, f.MessageMaxLte)
	}

	if f.MessageMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") IN (?)")
		values = append(values, f.MessageMinIn)
	}

	if f.MessageMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") IN (?)")
		values = append(values, f.MessageMaxIn)
	}

	if f.MessageMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MessageMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.MessageMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.MessageMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.MessageMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MessageMinPrefix))
	}

	if f.MessageMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.MessageMaxPrefix))
	}

	if f.MessageMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MessageMinSuffix))
	}

	if f.MessageMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("message")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.MessageMaxSuffix))
	}

	if f.SeenMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") = ?")
		values = append(values, f.SeenMin)
	}

	if f.SeenMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") = ?")
		values = append(values, f.SeenMax)
	}

	if f.SeenMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") != ?")
		values = append(values, f.SeenMinNe)
	}

	if f.SeenMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") != ?")
		values = append(values, f.SeenMaxNe)
	}

	if f.SeenMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") > ?")
		values = append(values, f.SeenMinGt)
	}

	if f.SeenMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") > ?")
		values = append(values, f.SeenMaxGt)
	}

	if f.SeenMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") < ?")
		values = append(values, f.SeenMinLt)
	}

	if f.SeenMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") < ?")
		values = append(values, f.SeenMaxLt)
	}

	if f.SeenMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") >= ?")
		values = append(values, f.SeenMinGte)
	}

	if f.SeenMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") >= ?")
		values = append(values, f.SeenMaxGte)
	}

	if f.SeenMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") <= ?")
		values = append(values, f.SeenMinLte)
	}

	if f.SeenMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") <= ?")
		values = append(values, f.SeenMaxLte)
	}

	if f.SeenMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("seen")+") IN (?)")
		values = append(values, f.SeenMinIn)
	}

	if f.SeenMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("seen")+") IN (?)")
		values = append(values, f.SeenMaxIn)
	}

	if f.URLMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") = ?")
		values = append(values, f.URLMin)
	}

	if f.URLMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") = ?")
		values = append(values, f.URLMax)
	}

	if f.URLMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") != ?")
		values = append(values, f.URLMinNe)
	}

	if f.URLMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") != ?")
		values = append(values, f.URLMaxNe)
	}

	if f.URLMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") > ?")
		values = append(values, f.URLMinGt)
	}

	if f.URLMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") > ?")
		values = append(values, f.URLMaxGt)
	}

	if f.URLMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") < ?")
		values = append(values, f.URLMinLt)
	}

	if f.URLMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") < ?")
		values = append(values, f.URLMaxLt)
	}

	if f.URLMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") >= ?")
		values = append(values, f.URLMinGte)
	}

	if f.URLMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") >= ?")
		values = append(values, f.URLMaxGte)
	}

	if f.URLMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") <= ?")
		values = append(values, f.URLMinLte)
	}

	if f.URLMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") <= ?")
		values = append(values, f.URLMaxLte)
	}

	if f.URLMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") IN (?)")
		values = append(values, f.URLMinIn)
	}

	if f.URLMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") IN (?)")
		values = append(values, f.URLMaxIn)
	}

	if f.URLMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLMinPrefix))
	}

	if f.URLMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLMaxPrefix))
	}

	if f.URLMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLMinSuffix))
	}

	if f.URLMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLMaxSuffix))
	}

	if f.ChannelMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") = ?")
		values = append(values, f.ChannelMin)
	}

	if f.ChannelMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") = ?")
		values = append(values, f.ChannelMax)
	}

	if f.ChannelMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") != ?")
		values = append(values, f.ChannelMinNe)
	}

	if f.ChannelMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") != ?")
		values = append(values, f.ChannelMaxNe)
	}

	if f.ChannelMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") > ?")
		values = append(values, f.ChannelMinGt)
	}

	if f.ChannelMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") > ?")
		values = append(values, f.ChannelMaxGt)
	}

	if f.ChannelMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") < ?")
		values = append(values, f.ChannelMinLt)
	}

	if f.ChannelMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") < ?")
		values = append(values, f.ChannelMaxLt)
	}

	if f.ChannelMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") >= ?")
		values = append(values, f.ChannelMinGte)
	}

	if f.ChannelMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") >= ?")
		values = append(values, f.ChannelMaxGte)
	}

	if f.ChannelMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") <= ?")
		values = append(values, f.ChannelMinLte)
	}

	if f.ChannelMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") <= ?")
		values = append(values, f.ChannelMaxLte)
	}

	if f.ChannelMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") IN (?)")
		values = append(values, f.ChannelMinIn)
	}

	if f.ChannelMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") IN (?)")
		values = append(values, f.ChannelMaxIn)
	}

	if f.ChannelMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ChannelMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ChannelMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ChannelMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ChannelMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ChannelMinPrefix))
	}

	if f.ChannelMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ChannelMaxPrefix))
	}

	if f.ChannelMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ChannelMinSuffix))
	}

	if f.ChannelMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("channel")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ChannelMaxSuffix))
	}

	if f.PrincipalMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") = ?")
		values = append(values, f.PrincipalMin)
	}

	if f.PrincipalMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") = ?")
		values = append(values, f.PrincipalMax)
	}

	if f.PrincipalMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") != ?")
		values = append(values, f.PrincipalMinNe)
	}

	if f.PrincipalMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") != ?")
		values = append(values, f.PrincipalMaxNe)
	}

	if f.PrincipalMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") > ?")
		values = append(values, f.PrincipalMinGt)
	}

	if f.PrincipalMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") > ?")
		values = append(values, f.PrincipalMaxGt)
	}

	if f.PrincipalMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") < ?")
		values = append(values, f.PrincipalMinLt)
	}

	if f.PrincipalMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") < ?")
		values = append(values, f.PrincipalMaxLt)
	}

	if f.PrincipalMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") >= ?")
		values = append(values, f.PrincipalMinGte)
	}

	if f.PrincipalMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") >= ?")
		values = append(values, f.PrincipalMaxGte)
	}

	if f.PrincipalMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") <= ?")
		values = append(values, f.PrincipalMinLte)
	}

	if f.PrincipalMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") <= ?")
		values = append(values, f.PrincipalMaxLte)
	}

	if f.PrincipalMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") IN (?)")
		values = append(values, f.PrincipalMinIn)
	}

	if f.PrincipalMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") IN (?)")
		values = append(values, f.PrincipalMaxIn)
	}

	if f.PrincipalMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PrincipalMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PrincipalMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalMinPrefix))
	}

	if f.PrincipalMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalMaxPrefix))
	}

	if f.PrincipalMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalMinSuffix))
	}

	if f.PrincipalMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principal")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalMaxSuffix))
	}

	if f.ReferenceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMin)
	}

	if f.ReferenceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMax)
	}

	if f.ReferenceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMinNe)
	}

	if f.ReferenceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMaxNe)
	}

	if f.ReferenceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMinGt)
	}

	if f.ReferenceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMaxGt)
	}

	if f.ReferenceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMinLt)
	}

	if f.ReferenceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMaxLt)
	}

	if f.ReferenceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMinGte)
	}

	if f.ReferenceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMaxGte)
	}

	if f.ReferenceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMinLte)
	}

	if f.ReferenceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMaxLte)
	}

	if f.ReferenceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMinIn)
	}

	if f.ReferenceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMaxIn)
	}

	if f.ReferenceMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMinPrefix))
	}

	if f.ReferenceMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMaxPrefix))
	}

	if f.ReferenceMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMinSuffix))
	}

	if f.ReferenceMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMaxSuffix))
	}

	if f.ReferenceIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMin)
	}

	if f.ReferenceIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMax)
	}

	if f.ReferenceIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMinNe)
	}

	if f.ReferenceIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMaxNe)
	}

	if f.ReferenceIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMinGt)
	}

	if f.ReferenceIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMaxGt)
	}

	if f.ReferenceIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMinLt)
	}

	if f.ReferenceIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMaxLt)
	}

	if f.ReferenceIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMinGte)
	}

	if f.ReferenceIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMaxGte)
	}

	if f.ReferenceIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMinLte)
	}

	if f.ReferenceIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMaxLte)
	}

	if f.ReferenceIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMinIn)
	}

	if f.ReferenceIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMaxIn)
	}

	if f.ReferenceIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceIDMinPrefix))
	}

	if f.ReferenceIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceIDMaxPrefix))
	}

	if f.ReferenceIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceIDMinSuffix))
	}

	if f.ReferenceIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceIDMaxSuffix))
	}

	if f.DateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMin)
	}

	if f.DateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMax)
	}

	if f.DateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMinNe)
	}

	if f.DateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMaxNe)
	}

	if f.DateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMinGt)
	}

	if f.DateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMaxGt)
	}

	if f.DateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMinLt)
	}

	if f.DateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMaxLt)
	}

	if f.DateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMinGte)
	}

	if f.DateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMaxGte)
	}

	if f.DateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMinLte)
	}

	if f.DateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMaxLte)
	}

	if f.DateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMinIn)
	}

	if f.DateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
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
