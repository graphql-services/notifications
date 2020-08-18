package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s NotificationSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("notifications"), sorts, joins)
}
func (s NotificationSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Message != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("message"), Direction: s.Message.String()}
		*sorts = append(*sorts, sort)
	}

	if s.MessageMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("message") + ")", Direction: s.MessageMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MessageMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("message") + ")", Direction: s.MessageMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Seen != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("seen"), Direction: s.Seen.String()}
		*sorts = append(*sorts, sort)
	}

	if s.SeenMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("seen") + ")", Direction: s.SeenMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SeenMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("seen") + ")", Direction: s.SeenMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Channel != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("channel"), Direction: s.Channel.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ChannelMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("channel") + ")", Direction: s.ChannelMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ChannelMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("channel") + ")", Direction: s.ChannelMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Principal != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("principal"), Direction: s.Principal.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PrincipalMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("principal") + ")", Direction: s.PrincipalMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PrincipalMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("principal") + ")", Direction: s.PrincipalMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Reference != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("reference"), Direction: s.Reference.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("referenceID"), Direction: s.ReferenceID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Date != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("date"), Direction: s.Date.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DateMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DateMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}
