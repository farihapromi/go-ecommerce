package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProducts(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("Id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid ID", 400)
		return
	}
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", 400)
		return
	}
	newProduct.ID = pId
	database.Update(newProduct)
	util.SendData(w, "product succesfully updated", 201)

}
