package model

import (
	_ "database/sql"
	_ "github.com/satori/go.uuid"
	_ "time"
)

type Language struct {
	ID              int        `gorm:"primary_key"`
	DisplayName		string	   `gorm:"not null"`
	Extension		string	   `gorm:"not null"`
}

func (language *Language) TableName() string {
	return "language"
}

func (language *Language) GetData() map[string]interface{} {
	return map[string]interface{}{
		"displayName": language.DisplayName,
		"extension": language.Extension,
	}
}