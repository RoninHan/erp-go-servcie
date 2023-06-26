package dao

import (
	"fmt"
	"sample-go-service/forms"
	"sample-go-service/global"
	"sample-go-service/models"
)

var PurchaseOrderDetails []models.PurchaseOrderDetails
var PurchaseOrderDetail models.PurchaseOrderDetails

func FindPurchaseOrdersDetail(orderID string) (int, []interface{}) {
	purchaseOrderDetailList := make([]interface{}, 0, len(PurchaseOrderDetails))
	// 查询所有的user
	result := global.DB.Where("order_id=?", orderID).Find(&PurchaseOrderDetails)

	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, purchaseOrderDetailList
	}

	total := len(PurchaseOrderDetails)

	result.Find(&PurchaseOrderDetails).Scan(purchaseOrderDetailList)

	return total, purchaseOrderDetailList
}

func CreatePurchaseOrderDetail(data models.PurchaseOrders) bool {
	if err := global.DB.Create(&data).Error; err != nil {
		fmt.Println("Fail to create purchase order detail", err)
		return false
	}
	return true
}

func UpdatePurchaseOrderDetail(id string, data forms.PurchaseOrderForm) bool {
	if err := global.DB.Model(&PurchaseOrderDetail).Where("id", id).Updates(data).Error; err != nil {
		fmt.Println("Fail to update purchase order detail", err)
		return false
	}
	return true
}
