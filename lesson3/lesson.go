package main

import "fmt"

const USDPrice = 34000.0

type BinanceOrder struct {
	Symbol   string
	Amount   float64
	Price    float64
	Currency string
}

func (bo BinanceOrder) GetAmount() float64 {
	return bo.Amount
}

func (bo BinanceOrder) GetPrice() float64 {
	return bo.Price
}

type PoloniexOrder struct {
	Rate         float64
	CurrencyPair string
	Type         string
	Amount       float64
}

func (po PoloniexOrder) GetAmount() float64 {
	return po.Amount
}

func (po PoloniexOrder) GetPrice() float64 {
	return po.Rate
}

type OrderServiceContract interface {
	GetFullSum() float64
	GetUSDTotal() float64
}

type OrderContract interface {
	GetAmount() float64
	GetPrice() float64
}

type OrderService struct {
	order OrderContract
}

func (o *OrderService) SetOrder(order OrderContract) {
	o.order = order
}

func (o OrderService) GetAmount() float64 {
	return o.order.GetAmount()
}

func (o OrderService) GetUSDTotal() float64 {
	amount := o.order.GetAmount()

	return amount * 1 / USDPrice
}

func main() {
	binanceOrder := BinanceOrder{Price: 1, Amount: 1}
	poloniexOrder := PoloniexOrder{Rate: 1, Amount: 1}

	service := OrderService{}
	service.SetOrder(binanceOrder)

	fmt.Println("Stats from binance", service.GetAmount(), service.GetUSDTotal())

	service.SetOrder(poloniexOrder)
	fmt.Println("Stats from poloniex", service.GetAmount(), service.GetUSDTotal())
}
