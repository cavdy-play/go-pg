package db

import (
	"time"
	"log"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

/*
	Table Name: product_items
*/
type ProductItem struct {
	tableName struct{}  `sql: "product_items_collection"`
	ID        int       `sql:"id,pk"`
	Name      string    `sql:"name,unique"`
	Desc      string    `sql:"desc"`
	Image     string    `sql:"image"`
	Price     float64   `sql:"price,type:real"`
	Features  string    `sql:"features"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

func CreateProdItemTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&ProductItem{}, opts)
	if createError != nil {
		log.Printf("Error while create product table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Prod table created")
	return nil
}
