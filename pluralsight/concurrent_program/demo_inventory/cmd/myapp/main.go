package main

import (
	"demo_inventory/internal/orders"
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOrders(&wg)
	wg.Wait()
	fmt.Println(orders.Orders)
}

// receiveOrders - Take the values from rawOrders Json object and translate into Go Object
func receiveOrders(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, rawOrder := range rawOrders {
		var newOrder orders.Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		orders.Orders = append(orders.Orders, newOrder)
	}
}

// rawOrders - Value grabbed from web services or database in a json format
var rawOrders = []string{
	`{"ProductCode":1111,"quantity":5,"status":1}`,
	`{"ProductCode":2222,"quantity":50,"status":2}`,
	`{"ProductCode":3333,"quantity":75,"status":3}`,
	`{"ProductCode":4444,"quantity":59,"status":4}`,
}
