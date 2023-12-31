package models

import "gorm.io/gorm"

type Books struct {
	ID        uint
	Author    *string
	Title     *string
	Publisher *string
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
