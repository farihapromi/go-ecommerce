package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am promi, I am a SW Engineer")
}

type Product struct {
	/*Struct tags are special metadata you can attach to struct fields.
	They are written in backticks (`...`) after the field definition.*/
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreflightReq(w, r)
	if r.Method != "GET" {
		http.Error(w, "Plz give me GET request", 400)
		return

	}
	// encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
	sendData(w, productList, 200)

}
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != "POST" {
		http.Error(w, "Plz give me GET request", 400)
		return

	}
	// r.body=>title,description,imageUrl,price=>Product er ekta instance=>ProdcutList=>append
	/*
		1.Take body informarion form r.body
		2.Create an instance using Product struct with the body information
		3.append the instsance into productList
	*/
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", 400)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
	// w.WriteHeader(201)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)

}
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //ALLOW  everyone.if we write 3000 isntead of * it wil suport 3000 port frontend
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Promi")

	w.Header().Set("Content-Type", "application/json")
}
func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

}
func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route
	mux.HandleFunc("/about", aboutHandler) //route
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is sweet and juicy",
		Price:       100,
		ImgUrl:      "https://images.othoba.com/images/thumbs/0040450_orange-1-kg.jpeg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "Mango is juicy and delicious",
		Price:       300,
		ImgUrl:      "https://dookan.com/cdn/shop/files/Dookan_Kesar_Mangoes_82bf0570-50bf-4f04-97b4-b1b415ebc733.png?v=1724714757&width=823",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "Red apple, crispy and sweet",
		Price:       200,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_Apple.jpg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Fresh yellow bananas, rich in potassium",
		Price:       50,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Fresh pineapple, tropical and juicy",
		Price:       250,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/cb/Pineapple_and_cross_section.jpg",
	}
	// Append all products
	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)

	// productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)

	// fmt.Println(productList)
}
