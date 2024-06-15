package model

/*
文章模型
*/

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //声明关联关系，不声明有可能会报错（反正我报错了……）
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null " json:"title"`
	Cid     int    `gorm:"type:int;not null " json:"cid"`
	Desc    string `gorm:"type:varchar(200) " json:"desc"`
	Content string `gorm:"type:longtext " json:"content"`
	Img     string `gorm:"type:varchar(100) " json:"img"`
}
