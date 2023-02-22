package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"

	"simple-project-be/backend/dictionary"
	"simple-project-be/backend/service"
	"simple-project-be/backend/utils"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params := mux.Vars(r)
	arr_col := []string{"id", "nama", "jenis", "jumlah", "harga"}
	arr_jenis := []string{"sayuran", "buah", "bumbu", "karbo", "protein"}
	err_syntax := ""

	jenis := params["jenis"]
	if (!utils.ArrContainsStr(arr_jenis, jenis) && jenis != "") {
		err_syntax = "invalid syntax"
	}

	hargaMinStr := query.Get("hargaMin")
	if (hargaMinStr == "") { hargaMinStr = "0" }
	hargaMinInt64, err := strconv.ParseInt(hargaMinStr, 10, 64)
	if err != nil {
		err_syntax = err.Error()
	}

	hargaMaxStr := query.Get("hargaMax")
	if (hargaMaxStr == "") { hargaMaxStr = "0" }
	hargaMaxInt64, err := strconv.ParseInt(hargaMaxStr, 10, 64)
	if err != nil {
		err_syntax = err.Error()
	}

	pageStr := query.Get("page")
	if (pageStr == "") { pageStr = "1" }
	pageInt64, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		err_syntax = err.Error()
	}

	sizeStr := query.Get("size")
	if (sizeStr == "") { sizeStr = "5" }
	sizeInt64, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		err_syntax = err.Error()
	}

	order_by := query.Get("orderby")
	if (!utils.ArrContainsStr(arr_col, order_by) && order_by != "") {
		err_syntax = "invalid syntax"
	}
	if (order_by == "") { order_by = "id" }

	order := query.Get("desc")
	if (order != "true" && order != "false" && order != "") {
		err_syntax = "invalid syntax"
	}
	desc := false
	if (order == "true") { desc = true }

	detail := map[string]interface{}{
		"page": pageInt64, "size": sizeInt64, "order_by": order_by, "desc": desc,
	}
	
	if (err_syntax != "") {
		fmt.Println("err parameter:", err_syntax)
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, Error: dictionary.InvalidParamError,
		})
	} else {
		res, err := service.GetProducts(
			jenis, hargaMinInt64, hargaMaxInt64, pageInt64, sizeInt64, order_by, desc,
		)
		if err != nil {
			fmt.Println("err get list products:", err)
			json.NewEncoder(w).Encode(dictionary.APIResponse{
				Data: nil, Detail: detail, Error: dictionary.UndisclosedError,
			})
		} else {
			json.NewEncoder(w).Encode(dictionary.APIResponse{
				Data: res, Detail: detail, Error: dictionary.NoError,
			})
		}
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err_syntax := ""

	idstring := params["id"]
	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		err_syntax = err.Error()
	}

	if (err_syntax != "") {
		fmt.Println("err parameter:", err_syntax)
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, Error: dictionary.InvalidParamError,
		})
	} else {
		res, err := service.GetProduct(idInt64)
		if err != nil {
			fmt.Println("err get product:", err)
			json.NewEncoder(w).Encode(dictionary.APIResponse{
				Data: nil, Error: dictionary.UndisclosedError,
			})
		} else {
			json.NewEncoder(w).Encode(dictionary.APIResponse{
				Data: res, Error: dictionary.NoError,
			})
		}
	}
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	arr_product := []dictionary.Product{}
	json.NewDecoder(r.Body).Decode(&arr_product)

	res, err := service.InsertProduct(arr_product)
	if err != nil {
		fmt.Println("err insert product:", err)
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: res, Error: dictionary.NoError,
		})
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product dictionary.Product
	json.NewDecoder(r.Body).Decode(&product)

	err := service.UpdateProduct(product)
	if err != nil {
		fmt.Println("err update product:", err)
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: product, Error: dictionary.NoError,
		})
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var arr_id []int64
	json.NewDecoder(r.Body).Decode(&arr_id)

	res, err := service.DeleteProduct(arr_id)
	if err != nil {
		fmt.Println("err delete product:", err)
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil, Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: res, Error: dictionary.NoError,
		})
	}
}
