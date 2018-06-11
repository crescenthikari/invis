package handler

import (
	"net/http"
	"github.com/crescenthikari/invis/inventory/service"
	"github.com/crescenthikari/invis/http/util"
)

type productHandler struct {
	s *service.InventoryService
}

func NewProductHanlder(service *service.InventoryService) *productHandler {
	return &productHandler{s: service}
}

func (p *productHandler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := p.s.GetProducts()
	if err != nil {
		ResponseErr(w, err)
	}
	if util.HasContentType(r, "text/csv") {
		WriteCsv(w, "Catatan Jumlah Barang", products)
	} else {
		ResponseOk(w, products)
	}
}
