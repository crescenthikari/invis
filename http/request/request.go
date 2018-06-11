package request

import (
	"time"
)

type PurchaseRequest struct {
	Date        string  `csv:"Waktu"`
	Sku         string  `csv:"SKU"`
	Name        string  `csv:"Nama Barang"`
	OrderQty    int64   `csv:"Jumlah Pemesanan"`
	ReceivedQty int64   `csv:"Jumlah Diterima"`
	Price       float64 `csv:"Harga Beli"`
	TotalPrice  float64 `csv:"Total"`
	ReceiptNo   string  `csv:"Nomer Kwitansi"`
	Note        string  `csv:"Catatan"`
}

type DateTime struct {
	time.Time
}

const (
	dateFormat = "2006/12/31 22:59"
)

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format(dateFormat), nil
}

// You could also use the standard Stringer interface
func (date *DateTime) String() (string) {
	return date.String() // Redundant, just for example
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse(dateFormat, csv)
	if err != nil {
		return err
	}
	return nil
}

func (date *DateTime) ToTime() time.Time {
	return date.Add(0)
}
