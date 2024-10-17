package orders

import "fmt"

// Order - Order Struct shows the properties of the Order classes.
type Order struct {
	ProductCode int
	Quantity    int
	Status      OrderStatus
}

// OrderStatus - Custom type as int
type OrderStatus int

// const - iota keyword makes it easy to define a sequence of related constants without manually assigning each value.
const (
	none OrderStatus = iota
	newly
	received
	reserved
	filled
)

var Orders []Order

// String - String Method is accepting Order type with receiver as "o" . This means String method can be called on any variable of type Order
func (o Order) String() string {
	return fmt.Sprintf("Product Code: %v, Quantity: %v, Status: %v\n", o.ProductCode, o.Quantity, oderStatusToText(o.Status))
}

// orderStatusToText - This function takes integer status from const values assigned as iota and translate to a string
func oderStatusToText(o OrderStatus) string {
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
