package model

type Order struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Status      string `json:"status"`
	TotalAmount int    `json:"totalAount"`
	Currency    string `json:"currency"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type OrderItems struct {
	Id          int    `json:"id"`
	OrderId     int    `json:"orderId"`
	ProductId   int    `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	UnitPrice   int    `json:"unitPrice"`
	Currecy     string `json:"currency"`
	TotalPrice  int    `json:"totalPrice"`
}

type StatusHistory struct {
	Id        int    `json:"id"`
	OrderId   int    `json:"orderId"`
	OldStatus string `json:"oldStatus"`
	NewStatus string `json:"newStatus"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"createdAt"`
}
