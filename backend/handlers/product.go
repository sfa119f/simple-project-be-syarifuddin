package handlers

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"simple-project-be/backend/dictionary"
	"simple-project-be/backend/service"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	pageStr := query.Get("page")
	if (pageStr == "") { pageStr = "1" }
	pageInt64, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	sizeStr := query.Get("size")
	if (sizeStr == "") { sizeStr = "5" }
	sizeInt64, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	order_by := query.Get("orderby")
	if (order_by == "") { order_by = "id" }
	order := query.Get("desc")
	desc := false
	if (order == "true") { desc = true }

	res, err := service.GetProducts(pageInt64, sizeInt64, order_by, desc)
	detail := map[string]interface{}{"page": pageInt64, "size": sizeInt64, "order_by": order_by, "desc": desc}
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Detail: detail, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: res, Detail: detail, Error: dictionary.NoError})
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idstring := params["id"]
	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	res, err := service.GetProduct(idInt64)
	
	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: nil, Error: dictionary.UndisclosedError})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{Data: res, Error: dictionary.NoError})
	}
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	arr_product := []dictionary.Product{}
	json.NewDecoder(r.Body).Decode(&arr_product)

	res, err := service.InsertProduct(arr_product)
	if err != nil {
		fmt.Println("err insert product:", err)
	}

	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil,
			Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: res,
			Error: dictionary.NoError,
		})
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var arr_id []int64
	json.NewDecoder(r.Body).Decode(&arr_id)

	res, err := service.DeleteProduct(arr_id)
	if err != nil {
		fmt.Println("err delete product:", err)
	}

	if err != nil {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: nil,
			Error: dictionary.UndisclosedError,
		})
	} else {
		json.NewEncoder(w).Encode(dictionary.APIResponse{
			Data: res,
			Error: dictionary.NoError,
		})
	}
}
