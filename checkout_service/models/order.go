package models

type Order struct {
	Products []Product `json:"products"`
}

type Product struct {
	Name       string `json:"name"`
	PriceCents int    `json:"price_cents"`
}
