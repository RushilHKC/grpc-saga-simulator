package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
}

type OrderDetails struct {
	ID       int
	Price    int
	Currency string
}
