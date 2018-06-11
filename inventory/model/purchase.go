package model

import (
	"github.com/satori/go.uuid"
)

type Purchase struct {
	Id               uuid.UUID `json:"id" csv:"-"`
	ReceiptNo        string    `json:"receipt_no" csv:"Nomer Kwitansi"`
	Date             string    `json:"time" csv:"Waktu"`
	Sku              string    `json:"sku" csv:"SKU"`
	OrderQuantity    int64     `json:"order_qty" csv:"Jumlah Pemesanan"`
	ReceivedQuantity int64     `json:"receive_qty" csv:"Jumlah Diterima"`
	Price            float64   `json:"price" csv:"Harga Beli"`
	Note             string    `json:"note" csv:"Catatan"`
}
