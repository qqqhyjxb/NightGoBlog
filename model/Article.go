package model

/*
文章模型
*/

import (
	"NightGoBlog/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //声明外键关联关系
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null " json:"title"`
	Cid     int    `gorm:"type:int;not null " json:"cid"`
	Desc    string `gorm:"type:varchar(200) " json:"desc"`
	Content string `gorm:"type:longtext " json:"content"`
	Img     string `gorm:"type:varchar(100) " json:"img"`
}

// CreateArt CreateCate 新增文章
func CreateArt(data *Article) int {

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE

}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article //article 的切片
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	if len(cateArtList) == 0 {
		// 没有找到匹配的记录
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCSE
}

// GetArtInfo  查询单个文章
func GetArtInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCSE

}

// GetArt 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCSE
}

// EditArt  编辑文章
func EditArt(id int, data *Article) int {
	var Art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Art).Where("id = ?", id).Updates(maps)
	if err.Error != nil {
		fmt.Printf("Update error: %v\n", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var Art Article
	// 如果设置了gorm.MODEL
	// 软删除，数据库中并不会真的删除该数据，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录
	err := db.Where("id = ?", id).Delete(&Art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
