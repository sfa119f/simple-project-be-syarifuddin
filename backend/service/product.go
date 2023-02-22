package service

import (
	"fmt"

	"simple-project-be/backend/database"
	"simple-project-be/backend/dictionary"
	"simple-project-be/backend/utils"
)

func GetProducts (
	jenis string, harga_min int64, harga_max int64, page int64, size int64, order_by string, desc bool,
) ([]dictionary.Product, error) {
	db := database.GetDB()

	str_select := ""
	if (jenis != "") { str_select = fmt.Sprint("jenis = '", jenis, "'and ") }
	if (harga_min != 0) { str_select = fmt.Sprintln(str_select, "harga >= ", harga_min, "and") }
	if (harga_max != 0) { str_select = fmt.Sprintln(str_select, "harga <= ", harga_max, "and") }
	if (str_select != "") {
		str_select = fmt.Sprint("where ", str_select[:len(str_select) - 4])
	}

	offset := (page - 1) * size
	order := "asc"
	if (desc) { order = "desc" }
	query := fmt.Sprintln(
		"select * from products", str_select,
		"order by", order_by, order, 
		"limit", size, "offset (", offset, ")",
	)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []dictionary.Product{}
	for rows.Next() {
		var product dictionary.Product
		if err := 
			rows.Scan(&product.Id, &product.Nama, &product.Jenis, &product.Jumlah, &product.Harga); 
		err != nil {
			return products, err
		}
		products = append(
			products, 
			dictionary.Product{
				Id: product.Id, Nama: product.Nama, Jenis: product.Jenis, 
				Jumlah: product.Jumlah, Harga: product.Harga,
		})
	}
	
	return products, nil
} 

func GetProduct(id int64) (dictionary.Product, error) {
	db := database.GetDB()
	query := `select * from products where id = $1`

	var res dictionary.Product
	if err := 
		db.QueryRow(query, id).Scan(&res.Id, &res.Nama, &res.Jenis, &res.Jumlah, &res.Harga); 
	err != nil {
		return res, err
	}
	return res, nil
}

func InsertProduct(arr_product []dictionary.Product) ([]int64, error) {
	db := database.GetDB()

	query :=`insert into products (nama, jenis, jumlah, harga) values `
	for idx, el := range arr_product {
		query = fmt.Sprint(
			query, "('", el.Nama, "', '", el.Jenis, "', ", el.Jumlah, ", ", el.Harga, ")",
		)
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

func UpdateProduct(product dictionary.Product) error {
	db := database.GetDB()

	query := 
		`update products set nama = $2, jenis = $3, jumlah = $4, harga = $5 
		where id = $1 returning id`

	var id int64
	if err := 
		db.QueryRow(query, product.Id, product.Nama, product.Jenis, product.Jumlah, product.Harga).Scan(&id);
	err != nil {
		return err
	}
	return nil
}

func DeleteProduct(arr_id []int64) ([]int64, error) {
	db := database.GetDB()

	string_id := utils.ArrIntToStr(arr_id)
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
