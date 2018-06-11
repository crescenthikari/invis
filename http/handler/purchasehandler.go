package handler

import (
	"github.com/crescenthikari/invis/inventory/service"
	"net/http"
	"github.com/crescenthikari/invis/http/util"
	"errors"
	"github.com/gocarina/gocsv"
	"github.com/crescenthikari/invis/http/request"
	"log"
	"strconv"
)

type purchaseHandler struct {
	s *service.InventoryService
}

func NewPurchaseHandler(service *service.InventoryService) *purchaseHandler {
	return &purchaseHandler{service}
}

func (h *purchaseHandler) ListPurchasesHandler(w http.ResponseWriter, r *http.Request) {
	purchases, err := h.s.GetPurchases()
	if err != nil {
		ResponseErr(w, err)
		return
	}
	if util.HasContentType(r, "text/csv") {
		WriteCsv(w, "Catatan Barang Masuk", purchases)
	} else {
		ResponseOk(w, purchases)
	}
}

func (h *purchaseHandler) AddPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.PostFormValue("sku")
	name := r.PostFormValue("name")
	receiptNo := r.PostFormValue("receipt_number")
	date := r.PostFormValue("date")
	ordered := r.PostFormValue("order_quantity")
	orderQty, _ := strconv.ParseInt(ordered, 10, 64)
	received := r.PostFormValue("received_quantity")
	receivedQty, _ := strconv.ParseInt(received, 10, 64)
	price := r.PostFormValue("price")
	priceVal, _ := strconv.ParseFloat(price, 64)
	note := r.PostFormValue("note")
	purchaseRequest := request.PurchaseRequest{
		Sku:         sku,
		Name:        name,
		ReceiptNo:   receiptNo,
		Date:        date,
		Price:       priceVal,
		OrderQty:    orderQty,
		ReceivedQty: receivedQty,
		Note:        note,
	}

	err := h.s.AddPurchaseRequest(purchaseRequest)
	if err != nil {
		ResponseErr(w, err)
	} else {
		ResponseOk(w, "OK")
	}
}

func (h *purchaseHandler) UpdatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.PostFormValue("sku")
	receiptNo := r.PostFormValue("receipt_number")
	received := r.PostFormValue("received_quantity")
	receivedQty, _ := strconv.ParseInt(received, 10, 64)
	note := r.PostFormValue("note")

	err := h.s.UpdatePurchaseRequest(sku, receiptNo, receivedQty, note)
	if err != nil {
		ResponseErr(w, err)
	} else {
		ResponseOk(w, "OK")
	}
}

func (h *purchaseHandler) ImportPurchaseDataHandler(w http.ResponseWriter, r *http.Request) {
	if util.HasContentType(r, "text/csv") {
		var purchases []*request.PurchaseRequest
		var err error
		if err = gocsv.Unmarshal(r.Body, &purchases); err != nil {
			ResponseErr(w, err)
			return
		}
		for _, purchase := range purchases {
			log.Println(*purchase)
			err = h.s.AddPurchaseRequest(*purchase)
		}
		if err != nil {
			ResponseErr(w, err)
		} else {
			ResponseOk(w, "OK")
		}
	} else {
		ResponseErr(w, errors.New("unsupported content type"))
	}
}
