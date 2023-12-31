package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `gorm:"column:problem_id;type:varchar(36)" json:"problem_id"`
	CategoryId    string         `gorm:"column:category_id;type:varchar(36)" json:"category_id"`
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`
}

func (ProblemCategory) TableName() string {
	return "problem_category"
}
