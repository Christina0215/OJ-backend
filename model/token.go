package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ApiToken struct {
	gorm.Model

	UserId    uuid.UUID    `gorm:"not null"`
	Token     string       `gorm:"unique;not null"`
	ExpiredAt time.Time

	User      User         `gorm:"foreignKey:UserId;association_foreignKey:ID"`
}

func (apiToken ApiToken) TableName() string {
	return "token"
}

func (apiToken *ApiToken) BeforeCreate(scope *gorm.Scope) error {
	//var a = uuid.NewV4()
	//fmt.Println(a)
	//_ = scope.SetColumn("Token", interface{}(a))
	apiToken.Token = uuid.NewV4().String()

	return nil
}

func (apiToken *ApiToken) AddTime(remember bool) {
	now := time.Now()
	duration, _ := time.ParseDuration("30m")
	if remember {
		apiToken.ExpiredAt = now.AddDate(0, 1, 0)
	} else {
		apiToken.ExpiredAt = now.Add(duration)
	}
}
