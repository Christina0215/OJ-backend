package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"primary_key;unique"`
	RoleID       int8      `gorm:"default:2;not null"`
	Email        string    `gorm:"unique;not null"`
	Username     string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null"`
	Gender       bool      `gorm:"default:0;not null"`
	Introduction string    `gorm:"type:text"`
	Avatar       string
	School       string
	Company      string
	Github       string
	solved       int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	ApiToken []ApiToken `gorm:"foreignKey:ID;association_foreignKey:UserId"`
	Problems []Problem  `gorm:"foreignKey:ID;association_foreignKey:UserId"`
	Records  []Record   `gorm:"foreignKey:ID;association_foreignKey:UserId"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("ID", uuid.NewV4())
	err = scope.SetColumn("CreatedAt", time.Now())
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (user *User) AfterUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (user *User) GetData(kind string) map[string]interface{} {
	db := orm.GetDB()
	var role Role
	var problems []Problem
	if err := db.Model(user).Related(&role).Find(&role).Error; err != nil {
		panic(err)
	}
	db.Model(user).Preload("Problems").Find(&problems)
	var passedRecord []string
	db.Raw("select distinct problem_id from record where user_id = ? and judge_result_id = 3",user.ID).Find(&passedRecord)
	switch kind {
	case "detail":
		return map[string]interface{}{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"avatar":       user.Avatar,
			"gender":       user.Gender,
			"role":         role.GetData(),
			"introduction": user.Introduction,
			"solved":       len(passedRecord),
		}
	default:
		return make(map[string]interface{})
	}
}
