package service

import (
	"fmt"
	"simple-project-be/backend/database"
	"simple-project-be/backend/dictionary"
)

func GetProducts(page int64, size int64, order_by string, desc bool) ([]dictionary.Product, error) {
	db := database.GetDB()

	order := "asc"
	if (desc) { order = "desc" }
	query := fmt.Sprintln("select * from products order by", order_by, order, "offset ((", page, "- 1 ) *", size, ") rows fetch next", size , "rows only;")

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []dictionary.Product{}
	for rows.Next() {
		var product dictionary.Product
		if err := rows.Scan(&product.Id, &product.Nama, &product.Jenis, &product.Jumlah, &product.Harga); err != nil {
			return products, err
		}
		products = append(products, dictionary.Product{Id: product.Id, Nama: product.Nama, Jenis: product.Jenis, Jumlah: product.Jumlah, Harga: product.Harga})
	}
	
	return products, nil
} 
