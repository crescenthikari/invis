package model

import (
	"github.com/satori/go.uuid"
	"time"
	"math/big"
)

const (
	BrokenItem  = "B"
	MissingItem = "M"
	OrderedItem = "O"

	BrokenItemNote  = "Barang Rusak"
	MissingItemNote = "Barang Hilang"
	OrderedItemNote = "Pesanan %s"
)

type Sale struct {
	Id              uuid.UUID
	Date            time.Time
	OrderId         string
	Sku             string
	OrderedQuantity int64
	SellingPrice    big.Float
	Status          string
}
