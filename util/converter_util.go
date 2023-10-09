package util

import "strconv"

func ConvertToPrice(s string) uint {
	price, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return uint(price)
}

func ConvertToQuantity(s string) uint {
	return ConvertToPrice(s)
}
