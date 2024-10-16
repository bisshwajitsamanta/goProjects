package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	receiveOrders()
	fmt.Println(orders)
}

// receiveOrders - Take the values from rawOrders Json object and translate into Go Object
func receiveOrders() {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		orders = append(orders, newOrder)
	}
}

// rawOrders - Value grabbed from web services or database in a json format
var rawOrders = []string{
	`{"ProductCode":1111,"quantity":5,"status":1}`,
	`{"ProductCode":2222,"quantity":50,"status":2}`,
	`{"ProductCode":3333,"quantity":75,"status":3}`,
	`{"ProductCode":4444,"quantity":59,"status":4}`,
}
