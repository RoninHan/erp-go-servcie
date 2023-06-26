package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var suppliers []models.Suppliers
var supplier models.Suppliers

func FindSupplies(name string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	supplierList := make([]interface{}, 0, len(suppliers))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("supplier_name=?", name).Offset(offset).Limit(page_size).Find(&suppliers)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, supplierList
	}

	total := len(suppliers)

	result.Offset(offset).Limit(page_size).Find(&suppliers).Scan(supplierList)

	return total, supplierList
}

func CreateSupplier(data models.Suppliers) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create supplier", err)
		return false
	}
	return true
}

func UpdateSupplier(id string, data forms.SupplierForm) bool {
	if err := global.DB.Model(&supplier).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update supplier", err)
		return false
	}
	return true
}
