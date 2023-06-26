package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

//员工离职

var terminations []models.Termination
var termination models.Termination

func FindTerminations(name string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	terminationList := make([]interface{}, 0, len(terminations))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("name=?", name).Offset(offset).Limit(page_size).Find(&terminations)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, terminationList
	}

	total := len(terminations)

	result.Scan(terminationList)

	return total, terminationList
}

func CreateTermination(data models.Termination) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create Department", err)
		return false
	}
	return true
}

func UpdateTermination(id string, data forms.TerminationForm) bool {
	if err := global.DB.Model(&termination).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update Department", err)
		return false
	}
	return true
}
