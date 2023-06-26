package models

import "time"

/*
产品（Products）表格：

product_id（产品ID，主键）
product_name（产品名称）
description（描述）
unit（单位）
current_stock（当前库存量）
*/

type Products struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	ProductName  string     `json:"product_name"`
	Description  *time.Time `json:"description"`
	Unit         string     `json:"unit"`
	CurrentStock string     `json:"current_stock"`
	CreatedAt    *time.Time `json:"created_at" gorm:"type:datetime"`
	CreatedBy    string     `json:"created_by"`
}

func (Products) TableName() string {
	return "products"
}
