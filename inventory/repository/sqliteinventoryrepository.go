package repository

import (
	"database/sql"
	"github.com/crescenthikari/invis/inventory/model"
	"errors"
	"fmt"
	"log"
	"github.com/satori/go.uuid"
)

const (
	queryCreateProductTable        = "CREATE TABLE IF NOT EXISTS `product` (`sku` TEXT NOT NULL UNIQUE, `name` TEXT,`quantity` INTEGER DEFAULT 0, PRIMARY KEY(`sku`))"
	queryCreatePurchaseTable       = "CREATE TABLE IF NOT EXISTS `purchase` ( `id` TEXT NOT NULL UNIQUE, `sku` TEXT NOT NULL, `receiptNo` TEXT NOT NULL, `date` TEXT, `orderQty` INTEGER, `receivedQty` INTEGER, `price` REAL, `note` TEXT, PRIMARY KEY(`id`))"
	querySelectProducts            = "SELECT * FROM product"
	insertProductStatement         = "INSERT INTO product(`sku`, `name`, `quantity`) VALUES(?,?,?)"
	udpateProductQuantityStatement = "UPDATE product SET quantity=? WHERE sku=?"
	querySelectPurchase            = "SELECT * FROM purchase"
	insertPurchaseStatement        = "INSERT INTO purchase(`id`, `sku`, `receiptNo`, `date`, `orderQty`, `receivedQty`, `price`, `note`) VALUES(?,?,?,?,?,?,?,?)"
	updatePurchaseStatement        = "UPDATE purchase SET receivedQty=?, note=? WHERE sku=? AND receiptNo=?"
)

type sqliteInventoryRepository struct {
	db *sql.DB
}

func NewSqliteInventoryRepository(db *sql.DB) *sqliteInventoryRepository {
	return &sqliteInventoryRepository{db: db}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (r *sqliteInventoryRepository) InitRepo() (err error) {
	_, err = r.db.Exec(queryCreateProductTable)
	_, err = r.db.Exec(queryCreatePurchaseTable)
	return
}

func (r *sqliteInventoryRepository) GetProducts() (products []model.Product, err error) {
	rows, err := r.db.Query(querySelectProducts)
	checkErr(err)

	var sku string
	var name string
	var qty int64

	products = []model.Product{}
	for rows.Next() {
		err = rows.Scan(&sku, &name, &qty)
		checkErr(err)

		products = append(products, model.Product{SKU: sku, Name: name, Quantity: qty})
	}
	return
}

func (r *sqliteInventoryRepository) GetProductBySKU(sku string) (product model.Product, err error) {
	stmt, err := r.db.Prepare(querySelectProducts + " WHERE sku=?")
	checkErr(err)

	rows, err := stmt.Query(sku)
	checkErr(err)

	var skuId string
	var name string
	var qty int64

	if rows.Next() {
		err = rows.Scan(&skuId, &name, &qty)
		checkErr(err)

		product = model.Product{SKU: sku, Name: name, Quantity: qty}
	} else {
		err = errors.New(fmt.Sprintf("Product with SKU %s not found", sku))
	}
	rows.Close()
	return
}

func (r *sqliteInventoryRepository) AddProduct(product model.Product) error {
	stmt, err := r.db.Prepare(insertProductStatement)
	checkErr(err)

	_, err = stmt.Exec(product.SKU, product.Name, product.Quantity)

	return err
}

func (r *sqliteInventoryRepository) UpdateProductQuantity(sku string, qty int64) error {
	stmt, err := r.db.Prepare(udpateProductQuantityStatement)
	checkErr(err)

	_, err = stmt.Exec(qty, sku)

	return err
}

func (r *sqliteInventoryRepository) GetPurchases() (purchases []model.Purchase, err error) {
	rows, err := r.db.Query(querySelectPurchase)
	checkErr(err)

	var id uuid.UUID
	var skuId string
	var receiptNumber string
	var date string
	var orderQty int64
	var receivedQty int64
	var price float64
	var note string

	for rows.Next() {
		err = rows.Scan(&id, &skuId, &receiptNumber, &date, &orderQty, &receivedQty, &price, &note)
		checkErr(err)

		purchase := model.Purchase{
			Id:               id,
			ReceiptNo:        receiptNumber,
			Date:             date,
			Sku:              skuId,
			OrderQuantity:    orderQty,
			ReceivedQuantity: receivedQty,
			Price:            price,
			Note:             note,
		}
		purchases = append(purchases, purchase)
	}
	return
}

func (r *sqliteInventoryRepository) FindPurchaseBySkuAndReceiptNo(sku string, receiptNo string) (purchase model.Purchase, err error) {
	stmt, err := r.db.Prepare(querySelectPurchase + " WHERE sku=? AND receiptNo=?")
	checkErr(err)

	rows, err := stmt.Query(sku, receiptNo)
	checkErr(err)

	if rows.Next() {
		var id uuid.UUID
		var skuId string
		var receiptNumber string
		var date string
		var orderQty int64
		var receivedQty int64
		var price float64
		var note string

		err = rows.Scan(&id, &skuId, &receiptNumber, &date, &orderQty, &receivedQty, &price, &note)
		checkErr(err)

		purchase = model.Purchase{
			Id:               id,
			ReceiptNo:        receiptNumber,
			Date:             date,
			Sku:              sku,
			OrderQuantity:    orderQty,
			ReceivedQuantity: receivedQty,
			Price:            price,
			Note:             note,
		}
	} else {
		err = errors.New(fmt.Sprintf("Product with SKU %s and Receipt Number %s not found", sku, receiptNo))
	}
	rows.Close()
	return
}

func (r *sqliteInventoryRepository) AddPurchase(p model.Purchase) (err error) {
	stmt, err := r.db.Prepare(insertPurchaseStatement)
	checkErr(err)

	_, err = stmt.Exec(p.Id, p.Sku, p.ReceiptNo, p.Date, p.OrderQuantity, p.ReceivedQuantity, p.Price, p.Note)
	return
}

func (r *sqliteInventoryRepository) AddPurchaseRequest(sku string, name string, orderQty int64,
	receivedQty int64, price float64, totalPrice float64,
	receiptNo string, note string, date string) (err error) {
	purchaseStmt, err := r.db.Prepare(insertPurchaseStatement)
	checkErr(err)

	updateQtyStmt, err := r.db.Prepare(udpateProductQuantityStatement)
	checkErr(err)

	product, err := r.GetProductBySKU(sku)
	productAvailable := err == nil

	if !productAvailable {
		err = r.AddProduct(model.Product{SKU: sku, Name: name})
		checkErr(err)
	}

	_, err = r.FindPurchaseBySkuAndReceiptNo(sku, receiptNo)
	purchaseAvailable := err == nil

	trx, err := r.db.Begin()
	checkErr(err)

	if !purchaseAvailable {
		log.Println("purchase not available")
		p := model.Purchase{
			Id:               uuid.NewV4(),
			ReceiptNo:        receiptNo,
			Date:             date,
			Sku:              sku,
			OrderQuantity:    orderQty,
			ReceivedQuantity: receivedQty,
			Price:            price,
			Note:             note,
		}

		_, err = trx.Stmt(purchaseStmt).Exec(p.Id, p.Sku, p.ReceiptNo, p.Date, p.OrderQuantity, p.ReceivedQuantity, p.Price, p.Note)
		if err != nil {
			log.Printf("error : %v", err)
			trx.Rollback()
			return
		}
		_, err = trx.Stmt(updateQtyStmt).Exec(product.Quantity+receivedQty, sku)
		if err != nil {
			log.Println(err)
			trx.Rollback()
			return
		}
	} else {
		trx.Rollback()
		return errors.New(fmt.Sprintf("Product with SKU %s and Receipt Number %s was found, please update it", sku, receiptNo))
	}
	trx.Commit()
	return nil
}

func (r *sqliteInventoryRepository) UpdatePurchaseRequest(sku string, receiptNo string, receivedQty int64, note string) (err error) {
	updatePurchaseStmt, err := r.db.Prepare(updatePurchaseStatement)
	checkErr(err)

	updateQtyStmt, err := r.db.Prepare(udpateProductQuantityStatement)
	checkErr(err)

	product, err := r.GetProductBySKU(sku)
	productAvailable := err == nil

	if !productAvailable {
		return
	}

	purchaseData, err := r.FindPurchaseBySkuAndReceiptNo(sku, receiptNo)
	purchaseAvailable := err == nil

	trx, err := r.db.Begin()
	checkErr(err)

	if purchaseAvailable {
		_, err = trx.Stmt(updatePurchaseStmt).Exec(receivedQty, note, sku, receiptNo)
		if err != nil {
			log.Printf("error : %v", err)
			trx.Rollback()
			return
		}
		_, err = trx.Stmt(updateQtyStmt).Exec(product.Quantity+(receivedQty-purchaseData.ReceivedQuantity), sku)
		if err != nil {
			log.Println(err)
			trx.Rollback()
			return
		}
	} else {
		trx.Rollback()
		return errors.New(fmt.Sprintf("Product with SKU %s and Receipt Number %s not found", sku, receiptNo))
	}
	trx.Commit()
	return nil
}

func (r *sqliteInventoryRepository) GetSales() (sales []model.Sale, err error) {
	panic("implement me")
}
