package model

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"qkcode/boot/orm"
	"time"
	_ "time"
)

type Solution struct {
	ID              uuid.UUID  `gorm:"primary_key;unique"`
	UserId          uuid.UUID  `gorm:"not null"`
	Title  			string     `gorm:"not null"`
	Content         string 	   `gorm:"not null"`
	ProblemID		string	   `gorm:"not null"`
	Language        string     `gorm:"not null"`

	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time

}

func (solution *Solution) TableName() string {
	return "solution"
}


func (solution *Solution)BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", time.Now())
	err = scope.SetColumn("CreatedAt", time.Now())
	return
}

func (solution *Solution)AfterUpdate(scope *gorm.Scope) (err error)  {
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (solution *Solution)GetData(kind string) map[string]interface{} {
	db := orm.GetDB()
	var author User
	db.Table("user").Where("id = ?", solution.UserId).Find(&author)
	var data = map[string]interface{} {
		"id": 			solution.ID,
		"title":		solution.Title,
		"user":		    author.Username,
	}

	if kind=="detail" {
		data["content"]=solution.Content
	}
	return data
}
