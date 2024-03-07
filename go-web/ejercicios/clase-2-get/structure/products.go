package structure

func NewProductsSlice(products []Product) *StorageProductsSlice {
	return &StorageProductsSlice{products}
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type StorageProductsSlice struct {
	Products []Product
}
