package repo

type Product struct {
	/*Struct tags are special metadata you can attach to struct fields.
	They are written in backticks (`...`) after the field definition.*/
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create()
	Get()
	List()
	Delete()
	Update()
}
type productRepo struct {
	productList []Product
}

func NewProductRepo() productRepo {
	return productRepo{}

}
func (r productRepo) Create() {

}
func (r productRepo) Get() {

}
func (r productRepo) Delete() {

}
func (r productRepo) Update() {

}
