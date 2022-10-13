package product

type Product struct {
	id          int64  `json:"id"`
	name        string `json:"name"`
	description string `json:"description"`
	price       string `json:"price"`
	isVariant   bool   `json:"isVarian"`
	stock       int64  `json:"stock"`
	discount    int64  `json:"discount"`
}
