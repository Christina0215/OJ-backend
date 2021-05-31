package migration

import "qkcode/model"

func AddTable() {
	InitMigration(map[string]interface{}{
		"user":              &model.User{},
		"problem":           &model.Problem{},
		"code":              &model.Code{},
		"role":              &model.Role{},
		"token":             &model.ApiToken{},
		"record":            &model.Record{},
		"verify_code":       &model.VerifyCode{},
		"judge_result":      &model.JudgeResult{},
		"language":          &model.Language{},
		"testcase":          &model.Testcase{},
		"contest_x_problem": &model.ContestXProblem{},
		"contest_x_user":    &model.ContestXUser{},
		"contest":           &model.Contest{},
		"solution":          &model.Solution{},
	})
}
