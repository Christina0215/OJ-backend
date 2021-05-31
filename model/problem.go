package model

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
	_ "time"
)

type Problem struct {
	ID             uuid.UUID   `gorm:"primary_key;unique"`
	UserId         uuid.UUID   `gorm:"not null"`
	Title          string 	   `gorm:"not null"`
	Type           string 	   `gorm:"not null"`
	Difficulty     string 	   `gorm:"not null"`
	Content        string 	   `gorm:"not null"`
	Samples        string 	   `gorm:"not null"`
	TimeLimit      string 	   `gorm:"not null"`
	MemoryLimit    string 	   `gorm:"not null"`
	StandardInput  string 	   `gorm:"not null"`
	StandardOutput string 	   `gorm:"not null"`
	Tip            string
	TestdataNumber int 		   `gorm:"not null"`

	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time

	User            User       `gorm:"foreignKey:UserId;association_foreignKey:ID"`
	Record			[]Record   `gorm:"foreignKey:ProblemId;association_foreignKey:ID"`
}

func (problem *Problem) TableName() string {
	return "problem"
}


func (problem *Problem)BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", time.Now())
	err = scope.SetColumn("CreatedAt", time.Now())
	return
}

func (problem *Problem)AfterUpdate(scope *gorm.Scope) (err error)  {
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (problem *Problem)GetData(kind string) map[string]interface{} {
	var data = map[string]interface{} {
		"id": 			problem.ID,
		"title":		problem.Title,
	}
	switch kind {
	case "detail":
		data["type"] =	problem.Type
		data["difficulty"] = problem.Difficulty
		data["content"] = problem.Content
		data["samples"] = problem.Samples
		data["timeLimit"] = problem.TimeLimit
		data["memoryLimit"] = problem.MemoryLimit
		data["standardInput"] = problem.StandardInput
		data["standardOutput"] = problem.StandardOutput
		data["tip"] = problem.Tip
		return data
	case "normal_list":
		data["id"] = problem.ID
		data["title"] = problem.Title
		data["type"] =	problem.Type
		data["difficulty"] = problem.Difficulty
		return data
	case "simplify_list":
		data["id"] = problem.ID
		data["title"] = problem.Title
		data["difficulty"] = problem.Difficulty
		return data
	default:
		return map[string]interface{}{}
	}
}
