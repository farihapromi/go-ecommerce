package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w, r)

	// if r.Method != "POST" {
	// 	http.Error(w, "Plz give me GET request", 400)
	// 	return
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(201)
	// 	return
	// }

	// }
	// r.body=>title,description,imageUrl,price=>Product er ekta instance=>ProdcutList=>append
	/*
		1.Take body informarion form r.body
		2.Create an instance using Product struct with the body information
		3.append the instsance into productList
	*/
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", 400)
		return
	}
	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, 201)
	// w.WriteHeader(201)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)

}
