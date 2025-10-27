package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("Id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid ID", 400)
		return
	}
	database.Delete(pId)
	util.SendData(w, "product  Deleted succesfully ", 200)

}
