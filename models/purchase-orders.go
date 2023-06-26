package models

import "time"

/*
采购订单（Purchase Orders）表格：

order_id（订单ID，主键）
supplier_id（供应商ID，外键关联到供应商表）
order_date（订单日期）
total_amount（订单总金额）
status（订单状态）
*/

type PurchaseOrders struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	OrderNum    string     `json:"order_num"`
	SupplierId  string     `json:"supplier_id"`
	OrderDate   *time.Time `json:"order_date" gorm:"type:datetime"`
	TotalAmount float64    `json:"total_amount"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"type:datetime"`
	CreatedBy   string     `json:"created_by"`
}

func (PurchaseOrders) TableName() string {
	return "purchaseOrders"
}
