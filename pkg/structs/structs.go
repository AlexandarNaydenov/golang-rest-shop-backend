package structs

type Order struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Products []struct {
		ID       string `json:"id"`
		Quantity int    `json:"quantity"`
	} `json:"products"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type Product struct {
	ID       int64
	Title    string
	Category string
	Quantity int
	Price    float64
}
