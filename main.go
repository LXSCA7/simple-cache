package main

import (
	"fmt"
)

type Product struct {
	Name  string  `json:"product_name"`
	Price float32 `json:"product_price"`
	Stock int     `json:"product_stock"`
	Code  string  `json:"product_code"`
	Sells int     `json:"product_sells"`
}

func main() {
	products := getTopSellers()
	for _, p := range products {
		fmt.Println(p)
	}
}

// simulating db get:
func getTopSellers() []Product {
	// time.Sleep(5 * time.Second)
	return []Product{
		{
			Name:  "mechanical keyboard",
			Price: 199.99,
			Stock: 500,
			Code:  "000006",
			Sells: 156,
		},
		{
			Name:  "mousepad",
			Price: 49.99,
			Stock: 100,
			Code:  "000012",
			Sells: 316,
		},
		{
			Name:  "headset",
			Price: 149.99,
			Stock: 236,
			Code:  "000027",
			Sells: 164,
		},
		{
			Name:  "monitor 144hz",
			Price: 999.99,
			Stock: 14,
			Code:  "000015",
			Sells: 596,
		},
		{
			Name:  "webcam",
			Price: 79.99,
			Stock: 120,
			Code:  "000033",
			Sells: 85,
		},
		{
			Name:  "gaming chair",
			Price: 299.99,
			Stock: 75,
			Code:  "000041",
			Sells: 210,
		},
		{
			Name:  "external hard drive 1TB",
			Price: 89.99,
			Stock: 300,
			Code:  "000008",
			Sells: 95,
		},
		{
			Name:  "wireless mouse",
			Price: 59.99,
			Stock: 400,
			Code:  "000019",
			Sells: 250,
		},
		{
			Name:  "microphone",
			Price: 119.99,
			Stock: 90,
			Code:  "000022",
			Sells: 70,
		},
		{
			Name:  "stream deck",
			Price: 169.99,
			Stock: 60,
			Code:  "000030",
			Sells: 45,
		},
	}
}
