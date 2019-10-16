package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s NotificationSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("notifications"), sorts, joins)
}
func (s NotificationSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Message != nil {
		*sorts = append(*sorts, aliasPrefix+"message "+s.Message.String())
	}

	if s.Seen != nil {
		*sorts = append(*sorts, aliasPrefix+"seen "+s.Seen.String())
	}

	if s.Channel != nil {
		*sorts = append(*sorts, aliasPrefix+"channel "+s.Channel.String())
	}

	if s.Principal != nil {
		*sorts = append(*sorts, aliasPrefix+"principal "+s.Principal.String())
	}

	if s.Reference != nil {
		*sorts = append(*sorts, aliasPrefix+"reference "+s.Reference.String())
	}

	if s.ReferenceID != nil {
		*sorts = append(*sorts, aliasPrefix+"referenceID "+s.ReferenceID.String())
	}

	if s.Date != nil {
		*sorts = append(*sorts, aliasPrefix+"date "+s.Date.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	return nil
}
