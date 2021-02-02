package main

import (
	"fmt"
	"math/rand"
)

type CryptoStatus string
type CryptoSide string
type CryptoType string

type cryptoType struct {
	Symbol              string
	orderId             int
	orderListId         int
	clientOrderId       string
	price               string
	origQty             string
	executedQty         string
	cummulativeQuoteQty string
	status              CryptoStatus
	timeInForce         string
	type_               CryptoType `json:"type"`
	side                CryptoSide
}

const N int = 10

func initArray(N int) []int {
	var result []int = make([]int, N, N)

	for i := range result {
		result[i] = rand.Int() % N
	}

	return result
}

func main() {
	var A []int = initArray(10)

	fmt.Println(A)
}
