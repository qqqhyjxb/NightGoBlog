package model

import (
	"NightGoBlog/utils/errmsg"
)

/*
分类模型
*/

type Category struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE

}

// CreateCate 新增分类
func CreateCate(data *Category) int {

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE

}

// todo 查询分类下的所有文章

// GetCate GetUsers GetUser 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {

	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil {
		return nil
	}
	return cate
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	// 如果设置了gorm.MODEL
	// 软删除，数据库中并不会真的删除该数据，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
