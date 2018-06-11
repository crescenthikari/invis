package service

import (
	"github.com/crescenthikari/invis/inventory/repository"
	"github.com/crescenthikari/invis/inventory/model"
	"github.com/satori/go.uuid"
	"github.com/crescenthikari/invis/http/request"
	"errors"
)

type InventoryService struct {
	r repository.InventoryRepository
}

func NewInventoryService(repository repository.InventoryRepository) *InventoryService {
	return &InventoryService{repository}
}

func (s *InventoryService) GetProducts() (products []model.Product, err error) {
	products, err = s.r.GetProducts()
	if products == nil {
		products = []model.Product{}
	}
	return
}

func (s *InventoryService) GetPurchases() (purchases []model.Purchase, err error) {
	purchases, err = s.r.GetPurchases()
	if purchases == nil {
		purchases = []model.Purchase{}
	}
	return
}

func (s *InventoryService) AddPurchase(purchase model.Purchase) (err error) {
	purchase.Id = uuid.NewV4()
	err = s.r.AddPurchase(purchase)
	return
}

func (s *InventoryService) AddPurchaseRequest(p request.PurchaseRequest) (err error) {
	if err = validatePurchaseRequest(p); err == nil {
		err = s.r.AddPurchaseRequest(p.Sku, p.Name, p.OrderQty, p.ReceivedQty, p.Price, p.TotalPrice, p.ReceiptNo, p.Note, p.Date)
	}
	return
}

func (s *InventoryService) UpdatePurchaseRequest(sku string, receiptNo string, receivedQty int64, note string) (err error) {
	err = s.r.UpdatePurchaseRequest(sku, receiptNo, receivedQty, note)
	return
}

func validatePurchaseRequest(p request.PurchaseRequest) error {
	if p.Sku == "" {
		return errors.New("SKU kosong")
	}
	if p.Name == "" {
		return errors.New("nama kosong")
	}
	if p.ReceiptNo == "" {
		return errors.New("nomor kwitansi kosong")
	}
	if p.OrderQty <= 0 {
		return errors.New("jumlah pesanan tidak valid")
	}
	if p.ReceivedQty < 0 {
		return errors.New("jumlah diterima tidak valid")
	}
	if p.Price < 0 {
		return errors.New("harga tidak valid")
	}

	return nil
}
