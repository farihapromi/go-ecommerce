package database

var productList []Product

type Product struct {
	/*Struct tags are special metadata you can attach to struct fields.
	They are written in backticks (`...`) after the field definition.*/
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func Store(p Product) {
	//prodcut list er objext
	productList = append(productList, p)
}
func List() []Product {
	return productList
}
func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}

	}
	return nil
}
func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}
func Delete(productID int) {
	var tempList []Product

	for idx, p := range productList {
		if p.ID != productID {
			tempList[idx] = p
		}
	}
	productList = tempList

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
	// Append all product.Products
	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)

	// product.productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)

	// fmt.Println(productList)
}
