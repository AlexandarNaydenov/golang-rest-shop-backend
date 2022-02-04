package structs

type Order struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Products []Product `json:"products"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}

type Product struct {
	ID       string
	Name     string
	Category string
	Quantity int
	Price    float64
}

type OrderedProduct struct {
	ID              string
	ProductId       string
	ProductQuantity int
	OrderId         string
}
