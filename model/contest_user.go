package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ContestXUser struct {
	gorm.Model
	ContestId    uuid.UUID
	UserId       uuid.UUID

	Contest      []Contest         `gorm:"foreignKey:ContestId;association_foreignKey:ID"`
}


