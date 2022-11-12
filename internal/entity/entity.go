package entity

type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Price       uint64 `json:"price"`
	Quantity    uint   `json:"quantity"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
}