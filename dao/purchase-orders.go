package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var purchaseOrders []models.PurchaseOrders
var purchaseOrder models.PurchaseOrders

func FindPurchaseOrders(orderNo string, page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	purchaseOrderList := make([]interface{}, 0, len(purchaseOrders))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Where("order_num=?", orderNo).Offset(offset).Limit(page_size).Find(&purchaseOrders)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, purchaseOrderList
	}

	total := len(suppliers)

	result.Offset(offset).Limit(page_size).Find(&purchaseOrders).Scan(purchaseOrderList)

	return total, purchaseOrderList
}

func CreatePurchaseOrder(data models.PurchaseOrders) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create purchase order", err)
		return false
	}
	return true
}

func UpdatePurchaseOrder(id string, data forms.PurchaseOrderForm) bool {
	if err := global.DB.Model(&purchaseOrder).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update purchase order", err)
		return false
	}
	return true
}
