package models

import "time"

/*
采购订单详情（Purchase Order Details）表格：
detail_id（订单详情ID，主键）
order_id（订单ID，外键关联到采购订单表）
product_id（产品ID，外键关联到产品表）
quantity（数量）
unit_price（单价）
subtotal（小计金额）
*/

type PurchaseOrderDetails struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	OrderId   string     `json:"order_id"`
	ProductId string     `json:"product_id"`
	Quantity  int        `json:"quantity"`
	Price     float64    `json:"price"`
	SubTotal  float64    `json:"sub_total"`
	CreatedAt *time.Time `json:"created_at" gorm:"type:datetime"`
	CreatedBy string     `json:"created_by"`
}

func (PurchaseOrderDetails) TableName() string {
	return "purchaseOrderDetails"
}
