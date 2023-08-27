package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Host struct {
	ID   uint   `gorm:"primaryKey;autoIncrementIncrement:true"`
	Name string `gorm:"uniqueIndex"`
	IP   string `gorm:"unique"`
}

func Connect(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
