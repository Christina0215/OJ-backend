package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ContestXProblem struct {
	gorm.Model
	ContestId       uuid.UUID
	ProblemId       uuid.UUID
	Order           int
	BaseScore       int

	Contest         Contest         `gorm:"foreignKey:ContestId;association_foreignKey:ID"`
	Record			[]Record		`gorm:"foreignKey:ProblemId;association_foreignKey:ID"`
}



func (ContestXProblem *ContestXProblem) GetData() map[string]interface{} {
	return map[string]interface{}{
		"record": ContestXProblem.Record,
	}
}