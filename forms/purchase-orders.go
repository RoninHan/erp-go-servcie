package forms

import "time"

type PurchaseOrderForm struct {
	OrderNum    string
	SupplierId  string
	OrderDate   *time.Time
	TotalAmount float64
	Status      string
}
