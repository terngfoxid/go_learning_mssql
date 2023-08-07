package Models

type Product struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Detail string `json:"detail"`
}

func (b *Product) TableName() string {
	return "product"
}
