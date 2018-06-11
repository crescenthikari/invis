package router

import (
	"github.com/crescenthikari/invis/http/middleware"
	"github.com/gorilla/mux"
	"github.com/crescenthikari/invis/http/handler"
	"github.com/crescenthikari/invis/inventory/service"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	// Add your routes as needed
	inventoryService := service.NewSqliteInventoryService()
	r.HandleFunc("/", handler.RootHandler)

	productHandler := handler.NewProductHanlder(inventoryService)
	r.HandleFunc("/products", productHandler.ListProductHandler).Methods("GET")

	purchaseHandler := handler.NewPurchaseHandler(inventoryService)
	r.HandleFunc("/purchases", purchaseHandler.ListPurchasesHandler).Methods("GET")
	r.HandleFunc("/purchases", purchaseHandler.AddPurchaseHandler).Methods("POST")
	r.HandleFunc("/purchases", purchaseHandler.UpdatePurchaseHandler).Methods("PATCH")
	r.HandleFunc("/purchases/import", purchaseHandler.ImportPurchaseDataHandler).Methods("POST")

	r.Use(middleware.LoggingMiddleware)

	return r
}
