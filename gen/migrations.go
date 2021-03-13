package gen

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB, options *gormigrate.Options, migrations []*gormigrate.Migration) error {
	m := gormigrate.New(db, options, migrations)

	// // it's possible to use this, but in case of any specific keys or columns are created in migrations, they will not be generated by automigrate
	// m.InitSchema(func(tx *gorm.DB) error {
	// 	return AutoMigrate(db)
	// })

	return m.Migrate()
}

func AutoMigrate(db *gorm.DB) (err error) {
	_db := db.AutoMigrate(
		Notification{},
	)

	if _db.Dialect().GetName() != "sqlite3" {

	}
	return _db.Error
}
