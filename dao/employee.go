package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var employees []models.Employee
var employee models.Employee

func FindEmployees(name string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	employeeList := make([]interface{}, 0, len(employees))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("user_name=?", name).Offset(offset).Limit(page_size).Find(&employees)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, employeeList
	}

	total := len(employees)

	result.Scan(employeeList)

	return total, employeeList
}

func CreateEmployee(data models.Employee) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create employee", err)
		return false
	}
	return true
}

func UpdateEmployee(id string, data forms.EmployeeForm) bool {
	if err := global.DB.Model(&employee).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update employee", err)
		return false
	}
	return true
}
