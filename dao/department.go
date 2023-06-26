package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var departments []models.Department
var department models.Department

func FindDepartments(name string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	departmentList := make([]interface{}, 0, len(departments))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("name=?", name).Offset(offset).Limit(page_size).Find(&departments)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, departmentList
	}

	total := len(departments)

	result.Scan(departmentList)

	return total, departmentList
}

func CreateDepartment(data models.Department) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create Department", err)
		return false
	}
	return true
}

func UpdateDepartment(id string, data forms.DepartmentForm) bool {
	if err := global.DB.Model(&department).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update Department", err)
		return false
	}
	return true
}
