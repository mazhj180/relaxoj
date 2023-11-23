package models

import (
	"gorm.io/gorm"
	"mazhj.com/relaxoj/internal/repository"
)

type ProblemBasic struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36)" json:"identity"`
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`
	Content    string `gorm:"column:content;type:text" json:"content"`
	MaxRuntime int    `gorm:"column:max_runtime;type:int" json:"max_runtime"`
	MaxMemory  int    `gorm:"column:max_memory;type:int" json:"max-memory"`
}

func (ProblemBasic) TableName() string {
	return "problem_basic"
}

func SelectProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := repository.Db.Model(new(ProblemBasic)).Distinct()
	return tx
}
