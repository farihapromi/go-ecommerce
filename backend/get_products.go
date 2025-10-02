package main

import "net/http"

func getProducts(w http.ResponseWriter, r *http.Request) {

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
	sendData(w, productList, 200)

}
