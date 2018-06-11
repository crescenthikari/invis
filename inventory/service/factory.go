package service

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/crescenthikari/invis/inventory/repository"
)

func NewSqliteInventoryService() *InventoryService {
	db, err := sql.Open("sqlite3", "./invis.db")
	if err != nil {
		panic(err)
	}
	inventoryRepo := repository.NewSqliteInventoryRepository(db)
	inventoryRepo.InitRepo()
	service := NewInventoryService(inventoryRepo)
	return service
}
