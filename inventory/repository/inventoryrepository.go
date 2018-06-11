package repository

import (
	"github.com/crescenthikari/invis/inventory/model"
)

type InventoryRepository interface {
	GetProducts() (products []model.Product, err error)
	GetProductBySKU(sku string) (product model.Product, err error)
	AddProduct(product model.Product) error
	UpdateProductQuantity(sku string, qty int64) error

	GetPurchases() (purchases []model.Purchase, err error)
	AddPurchase(purchase model.Purchase) error
	AddPurchaseRequest(sku string, name string, orderQty int64,
		receivedQty int64, price float64, totalPrice float64,
		receiptNo string, note string, date string) (err error)
	UpdatePurchaseRequest(sku string, receiptNo string, receivedQty int64, note string) (err error)
	FindPurchaseBySkuAndReceiptNo(sku string, receiptNo string) (purchase model.Purchase, err error)

	GetSales() (sales []model.Sale, err error)
}
