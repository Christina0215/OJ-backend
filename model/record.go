package model

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Record struct {
	gorm.Model
	ContestId     int
	UserId        uuid.UUID `gorm:"not null"`
	ProblemId     string    `gorm:"not null"`
	LanguageId    int       `gorm:"not null"`
	JudgeResultId int 	    `gorm:"not null;default:1"`
	CompileInfo   string    `gorm:"type:text"`
	TimeCost      int
	MemoryCost    int
	Score		  int

	Code        Code        `gorm:"ForeignKey:RecordId;association_foreignKey:ID"`
	JudgeResult JudgeResult `gorm:"foreignKey:JudgeResultId;association_foreignKey:ID"`
	Language	Language 	`gorm:"foreignKey:LanguageId;association_foreignKey:ID"`
	Testcase    Testcase    `gorm:"foreignKey:RecordId;association_foreignKey:ID"`
	Problem		Problem		`gorm:"foreignKey:ID;association_foreignKey:ProblemId"`
	ContestXProblem		ContestXProblem	`gorm:"foreignKey:ProblemId;association_foreignKey:ID"`
}


func (record *Record) TableName() string {
	return "record"
}

func (record *Record)GetData(kind string) map[string]interface{} {
	var data = map[string]interface{} {
		"id": 			record.ID,
		"problemId":		record.ProblemId,
	}
	switch kind {
	case "list":
		data["createdAt"] =	record.CreatedAt.Format("2006-01-02 15:04:05")
		data["timeCost"] = record.TimeCost
		data["memoryCost"] = record.MemoryCost
		data["language"] = record.Language.GetData()
		data["judgeResult"] = record.JudgeResult.GetData()
		data["compileInfo"] = record.CompileInfo
		data["score"] = record.Score
		return data
	case "detail":
		data["timeCost"] =	record.TimeCost
		data["memoryCost"] = record.MemoryCost
		data["language"] = record.Language.GetData()
		data["judgeResult"] = record.JudgeResult.GetData()
		data["compileInfo"] = record.CompileInfo
		data["testdataNum"] = record.Problem.TestdataNumber
		data["code"] = record.Code.GetData()
		data["score"] = record.Score
		return data
	default:
		return map[string]interface{}{}
	}
}