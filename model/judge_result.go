package model

import (
	_ "database/sql"
	_ "github.com/satori/go.uuid"
	_ "time"
)

type JudgeResult struct {
	ID            int          `gorm:"primary_key"`
	Alias		  string 	   `gorm:"not null"`
	En			  string 	   `gorm:"not null"`
	Zh		      string	   `gorm:"not null"`
	Color         string  	   `gorm:"not null"`
}

func (judgeResult *JudgeResult) TableName() string {
	return "judge_result"
}

func (judgeResult *JudgeResult) GetData() map[string]interface{} {
	return map[string]interface{}{
		"alias": judgeResult.Alias,
		"en": judgeResult.En,
		"zh": judgeResult.Zh,
		"color": judgeResult.Color,
	}
}