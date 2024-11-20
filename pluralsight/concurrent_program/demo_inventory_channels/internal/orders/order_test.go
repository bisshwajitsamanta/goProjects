package orders

import "testing"

func TestOrder_String(t *testing.T) {
	order := Order{ProductCode: 1111, Quantity: 5, Status: Newly}
	expected := "Product Code: 1111, Quantity: 5, Status: Newly\n"
	if order.String() != expected {
		t.Errorf("Expected %v, go %v", expected, order.String())
	}
}

func TestOrderStatusToText(t *testing.T) {
	tests := []struct {
		status   OrderStatus
		expected string
	}{
		{Newly, "Newly"},
		{none, "None"},
		{received, "Received"},
		{reserved, "Reserved"},
		{filled, "Filled"},
		{OrderStatus(10), "Unknown status"},
	}
	for _, test := range tests {
		result := oderStatusToText(test.status)
		if result != test.expected {
			t.Errorf("Expected %v, go %v", test.expected, result)
		}
	}
}
