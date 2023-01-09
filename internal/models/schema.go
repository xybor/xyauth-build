package models

import "github.com/xybor/xyauth/internal/database"

func DeleteTable() error {
	return database.Get().Migrator().DropTable(
		UserCredential{},
		User{},
		Client{},
	)
}

func Migrate() error {
	return database.Get().AutoMigrate(
		UserCredential{},
		User{},
		Client{},
	)
}
