package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var products []models.Products
var product models.Products

func FindProducts(productName string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	productList := make([]interface{}, 0, len(products))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("product_name=?", productName).Offset(offset).Limit(page_size).Find(&products)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, productList
	}

	total := len(products)

	result.Offset(offset).Limit(page_size).Find(&purchaseOrders).Scan(productList)

	return total, productList
}

func CreateProduct(data models.Products) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create product", err)
		return false
	}
	return true
}

func UpdateProducts(id string, data forms.PurchaseOrderForm) bool {
	if err := global.DB.Model(&product).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update product", err)
		return false
	}
	return true
}
