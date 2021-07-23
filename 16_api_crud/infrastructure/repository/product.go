package repository

import (
	"16_api_crud/entities"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//ProductPG product postgreSQL repo
type ProductPG struct {
	db *sql.DB
}

//NewProductPG create new repository
func NewProductPG(db *sql.DB) *ProductPG {
	return &ProductPG{
		db: db,
	}
}

//List return list of all products
func (r *ProductPG) List() ([]*entities.Product, error) {
	var products []*entities.Product
	rows, err := r.db.Query("select * from Products")
	if err != nil {
		return nil, fmt.Errorf("cannot read product table, err: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p entities.Product

		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}

		products = append(products, &p)
	}

	return products, nil
}
