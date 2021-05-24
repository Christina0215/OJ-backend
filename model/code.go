package model

import (
	_ "database/sql"
	_ "github.com/jinzhu/gorm"
	"io/ioutil"
)

type Code struct {
	ID         int    `gorm:"primary_key"`
	RecordId   int    `gorm:"not null"`
	LanguageId int    `gorm:"not null"`
	Filename   string `gorm:"not null"`
}

func (code *Code) TableName() string {
	return "code"
}

func (code *Code) GetData() string {
	file, err := ioutil.ReadFile("public/code/" + code.Filename)
	if err != nil {
		panic(err)
	}
	var codeString = string(file)
	return codeString
}
