package src

import (
	"github.com/graphql-services/notifications/gen"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func GetMigrations(db *gen.DB) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		&gormigrate.Migration{
			ID: "0000_INIT",
			Migrate: func(tx *gorm.DB) error {
				return db.AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				// there's not much we can do if initialization/automigration failes
				return nil
			},
		},
	}
}
