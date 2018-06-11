package model

type Product struct {
	SKU      string `json:"sku" schema:"sku" csv:"SKU"`
	Name     string `json:"name" schema:"name" csv:"Nama Item"`
	Quantity int64  `json:"qty" schema:"quantity" csv:"Jumlah Sekarang"`
}
