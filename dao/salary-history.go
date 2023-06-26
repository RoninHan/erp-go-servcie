package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

//薪资历史

var salaryHistorys []models.SalaryHistory
var salaryHistory models.SalaryHistory

func FindSalaryHistorys(id string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	salaryHistoryList := make([]interface{}, 0, len(salaryHistorys))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("employee_id=?", id).Offset(offset).Limit(page_size).Find(&salaryHistorys)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, salaryHistoryList
	}

	total := len(salaryHistorys)

	result.Scan(salaryHistoryList)

	return total, salaryHistoryList
}

func CreateSalaryHistory(data models.SalaryHistory) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create Department", err)
		return false
	}
	return true
}

func UpdateSalaryHistory(id string, data forms.TerminationForm) bool {
	if err := global.DB.Model(&termination).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update Department", err)
		return false
	}
	return true
}
