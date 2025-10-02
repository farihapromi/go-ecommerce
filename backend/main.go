package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server running on :8080")

	// Use cors-enabled globalRouter
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is sweet and juicy",
		Price:       100,
		ImgUrl:      "https://images.othoba.com/images/thumbs/0040450_orange-1-kg.jpeg",
	}

	prd2 := product.Product{
		ID:          2,
		Title:       "Mango",
		Description: "Mango is juicy and delicious",
		Price:       300,
		ImgUrl:      "https://dookan.com/cdn/shop/files/Dookan_Kesar_Mangoes_82bf0570-50bf-4f04-97b4-b1b415ebc733.png?v=1724714757&width=823",
	}

	prd3 := product.Product{
		ID:          3,
		Title:       "Apple",
		Description: "Red apple, crispy and sweet",
		Price:       200,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_Apple.jpg",
	}

	prd4 := product.Product{
		ID:          4,
		Title:       "Banana",
		Description: "Fresh yellow bananas, rich in potassium",
		Price:       50,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd5 := product.Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Fresh pineapple, tropical and juicy",
		Price:       250,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/cb/Pineapple_and_cross_section.jpg",
	}
	// Append all product.Products
	product.ProductList = append(product.ProductList, prd1)
	product.ProductList = append(product.ProductList, prd2)
	product.ProductList = append(product.ProductList, prd3)
	product.ProductList = append(product.ProductList, prd4)
	product.ProductList = append(product.ProductList, prd5)

	// product.ProductList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)

	// fmt.Println(productList)
}
