package models

type Product struct {
	Name     string  `json:"name"`
	Barcode  string  `json:"barcode"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
