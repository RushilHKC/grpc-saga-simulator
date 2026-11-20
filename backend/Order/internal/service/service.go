package service

type OrderService struct {
	orderRepo OrderService.OrderRepository
}

func NewOrderService(OrderRepo OrderService.OrderRepository) OrderService.OrderService {
	return &OrderService{
		orderRepo: OrderRepo,
	}
}
