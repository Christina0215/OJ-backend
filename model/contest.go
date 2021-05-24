package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Contest struct {
	ID           uuid.UUID    `gorm:"primary_key;unique"`
	UserId       uuid.UUID
	Title        string

	Introduction string
	Notification string
	Enabled      bool

	StartAt int64
	EndAt   int64

	ContestXUser      []ContestXUser `gorm:"foreignKey:ID;association_foreignKey:ContestId"`
	ContestXProblem      []ContestXProblem        `gorm:"foreignKey:ID;association_foreignKey:ContestId"`
}

func (contest *Contest) TableName() string {
	return "contest"
}

func (contest *Contest)GetData(kind string) map[string]interface{} {
	var data = map[string]interface{} {
		"id": 			contest.ID,
	}
	switch kind {
	case "detail":
		data["title"] =	contest.Title
		data["introduction"] = contest.Introduction
		data["notification"] = contest.Notification
		data["start_at"] = contest.StartAt
		data["end_at"] = contest.EndAt
		return data
	case "list":
		data["title"] = contest.Title
		data["start_at"] = contest.StartAt
		data["end_at"] =	contest.EndAt
		data["enabled"] = contest.Enabled
		now:=time.Now().Unix()
		if (contest.StartAt > now) {
			data["status"] = "Pending"
		} else if (now < contest.EndAt) {
			data["status"] = "Running"
		} else {
			data["status"] = "Finished"
		}
		return data
	default:
		return map[string]interface{}{}
	}
}

