package models

import "time"

/*
供应商（Suppliers）表格：
	supplier_id（供应商ID，主键）
	supplier_name（供应商名称）
	contact_person（联系人）
	contact_number（联系电话）
	email（邮箱）
	address（地址）
*/

type Suppliers struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	SupplierName  string     `json:"supplier_name"`
	ContactPerson string     `json:"contact_person"`
	ContactNumber string     `json:"contact_number"`
	Email         string     `json:"email"`
	Address       string     `json:"address"`
	CreatedAt     *time.Time `json:"created_at" gorm:"type:datetime"`
	CreatedBy     string     `json:"created_by"`
}

func (Suppliers) TableName() string {
	return "suppliers"
}
