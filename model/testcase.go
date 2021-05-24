package model

type Testcase struct {
	ID            int    `gorm:"primary_key"`
	RecordId      int    `gorm:"not null"`
	TestdataIndex int    `gorm:"not null"`
	JudgeResultId int    `gorm:"not null"`
	TimeCost      int    `gorm:"not null"`
	MemoryCost    int    `gorm:"not null"`
	Diff          string `gorm:"type:text"`
}

func (testcase *Testcase) GetData() map[string]interface{} {
	return map[string]interface{}{
		"testdataIndex": testcase.TestdataIndex,
		"timeCost":      testcase.TimeCost,
		"memoryCost":    testcase.MemoryCost,
		"diff":          testcase.Diff,
	}
}
