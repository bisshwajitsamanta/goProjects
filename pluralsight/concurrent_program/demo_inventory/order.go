package main

import "fmt"

// order - Order Struct shows the properties of the order classes.
type order struct {
	ProductCode int
	Quantity    int
	Status      orderStatus
}

// orderStatus - Custom type as int
type orderStatus int

// const - iota keyword makes it easy to define a sequence of related constants without manually assigning each value.
const (
	none orderStatus = iota
	newly
	received
	reserved
	filled
)

var orders []order

// String - String Method is accepting order type with receiver as "o" . This means String method can be called on any variable of type order
func (o order) String() string {
	return fmt.Sprintf("Product Code: %v, Quantity: %v, Status: %v\n", o.ProductCode, o.Quantity, oderStatusToText(o.Status))
}

// orderStatusToText - This function takes integer status from const values assigned as iota and translate to a string
func oderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "None"
	case newly:
		return "Newly"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown status"
	}
}
