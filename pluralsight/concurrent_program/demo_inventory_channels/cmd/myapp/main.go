package main

import (
	"demo_inventory_channels/internal/orders"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

/*
	rawOrders → [receiveOrders] → receiveOrderCh → [validateOrders]
	→ validateOrderCh (valid) → invalidOrderCh (invalid)

*/

func main() {
	var wg sync.WaitGroup
	var (
		receiveOrderCh  = make(chan orders.Order)
		validateOrderCh = make(chan orders.Order)
		invalidOrderCh  = make(chan orders.InvalidOrder)
	)
	go receiveOrders(receiveOrderCh)
	go validateOrders(receiveOrderCh, validateOrderCh, invalidOrderCh)
	wg.Add(1)
	go func(validOrderCh <-chan orders.Order, invalidOrderCh <-chan orders.InvalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validateOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v. Issue: %v\n", order, order.Error)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validateOrderCh, invalidOrderCh)
	wg.Wait()
}

// validateOrders - This function is going to receive the orders and validate for negative orders
func validateOrders(in, out chan orders.Order, errCh chan orders.InvalidOrder) {
	for order := range in {
		if order.Quantity <= 0 {
			// error condition
			errCh <- orders.InvalidOrder{Order: order, Error: errors.New("invalid order quantity, quantity must be greater than zero")}
		} else {
			// success path
			out <- order
		}
	}
	close(out)
	close(errCh)
}

// receiveOrders - Take the values from rawOrders Json object and translate into Go Object
func receiveOrders(out chan<- orders.Order) {
	for _, rawOrder := range rawOrders {
		var newOrder orders.Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		out <- newOrder
	}
	close(out)
}

// rawOrders - Value grabbed from web services or database in a json format
var rawOrders = []string{
	`{"ProductCode":1111,"quantity":5,"status":1}`,
	`{"ProductCode":2222,"quantity":50,"status":2}`,
	`{"ProductCode":3333,"quantity":75,"status":3}`,
	`{"ProductCode":4444,"quantity":59,"status":4}`,
}
