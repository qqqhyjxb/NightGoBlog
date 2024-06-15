package model

/*
分类模型
*/

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
