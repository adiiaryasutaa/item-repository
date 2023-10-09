package util

import "strconv"

// ConvertToPrice Convert a string into item price type
func ConvertToPrice(s string) uint {
	price, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return uint(price)
}

// ConvertToQuantity Convert a string into item quantity type
func ConvertToQuantity(s string) uint {
	return ConvertToPrice(s)
}
