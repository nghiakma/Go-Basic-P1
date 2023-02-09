// biến là tên đc đặt cho một vị trí bộ nhớ để lưu trữ giá trị
package main

import (
	"fmt"
	"math"
)

func main() {
	var age int = 29

	fmt.Println("My age is ", age)

	//khai báo các biến khác nhau
	var (
		name   = "khai"
		age1   = 30
		height int
	)
	fmt.Println("my name is", name, ", age is", age1, "and height is", height)

	//cú pháp để khai báo ngắn gọn (short hand declaration)
	count, myname := 10, "hien"
	fmt.Println("count is", count, "name is", myname)

	a, b := 10.0, 20.0
	c := math.Min(a, b) //float 64
	fmt.Println("Minimum value is", c)

	//strongly typed: phải đảm bảo kiểu của một object không thay đổi đột ngột,không tường minh

}
