package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36)" json:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity" json:"user_identity"`
	Path            string `gorm:"column:path" json:"path"` //代码存放路径
}

func (SubmitBasic) TableName() string {
	return "submit_basic"
}
