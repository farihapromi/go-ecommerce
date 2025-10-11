package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	// handleCors(w)
	// handlePreflightReq(w, r)
	//for uisng HandlerFunc we dont have to check if it is not get
	// if r.Method != "GET" {
	// 	http.Error(w, "Plz give me GET request", 400)
	// 	return
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	// }
	// encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
	util.SendData(w, database.List(), 200)

}
