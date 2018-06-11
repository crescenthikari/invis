package model

import (
	"github.com/satori/go.uuid"
)

type Purchase struct {
	Id               uuid.UUID `json:"id"`
	ReceiptNo        string    `json:"receipt_no"`
	Date             string    `json:"time"`
	Sku              string    `json:"sku"`
	OrderQuantity    int64     `json:"order_qty"`
	ReceivedQuantity int64     `json:"receive_qty"`
	Price            float64   `json:"price"`
	Note             string    `json:"note"`
}
