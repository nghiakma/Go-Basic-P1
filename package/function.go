package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Mutliple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total: ", totalPrice)
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f ", area, perimeter)
	fmt.Println()

	area, _ = rectProps(12.1, 13.4)
	fmt.Println("Area: ", area)
}

//black identified(định danh trống )
