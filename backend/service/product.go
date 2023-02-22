package service

import (
	"fmt"
	"database/sql"
	"errors"
	"strings"

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

func GetProduct(id int64) (dictionary.Product, error) {
	db := database.GetDB()
	query := `select * from products where id = $1`

	var res dictionary.Product
	if err := db.QueryRow(query, id).Scan(&res.Id, &res.Nama, &res.Jenis, &res.Jumlah, &res.Harga); err != nil {
		if (err == sql.ErrNoRows) {
			return res, errors.New("user nil")
		}
		return res, errors.New("error")
	}
	return res, nil
}

func InsertProduct(arr_product []dictionary.Product) ([]int64, error) {
	db := database.GetDB()

	query :=`insert into products (nama, jenis, jumlah, harga) values `
	for idx, el := range arr_product {
		query = fmt.Sprint(query, "('", el.Nama, "', '", el.Jenis, "', ", el.Jumlah, ", ", el.Harga, ")")
		if (idx != len(arr_product) - 1) { query = fmt.Sprint(query, ", ") }
	}
	query = fmt.Sprint(query, " returning id")

	rows, err := db.Query(query)
	if err != nil { return nil, err }
	defer rows.Close()

	var arr_id []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil { return arr_id, err }
		arr_id = append(arr_id, id)
	}

	return arr_id, err
}

func DeleteProduct(arr_id []int64) ([]int64, error) {
	db := database.GetDB()

	string_id := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr_id)), ", "), "[]")
	query := fmt.Sprint("delete from products where id in (", string_id, ") returning id")

	rows, err := db.Query(query)
	if err != nil { return nil, err }
	defer rows.Close()

	var del_id []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil { return del_id, err }
		del_id = append(del_id, id)
	}

	return del_id, err
}
