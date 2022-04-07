package models

import (
	"gorm.io/gorm"
)


type Books struct {
	ID			uint		`gorm:"primary key;autoIncrement" json:"id"`
	Author		*string		`json:"author"`
	title		*string		`json:"title"`
	Publisher	*string		`json:"Publisher"`
	Genre		*string		`json:"Genre"`
}

func Migratebooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}