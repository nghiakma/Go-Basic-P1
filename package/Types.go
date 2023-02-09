//Các kiểu cơ bản trong go
//bool, Number Types(int, unit(không âm), float, complex, byte, rune), string

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	fmt.Println("a: ", a, "b: ", b)
	c := a && b
	fmt.Println("c: ", c)
	d := a || b
	fmt.Println("d: ", d)

	//Singed integers, sizeof: kiem tra kich thuoc va kieu, 32bit(4byte), 64bit(8byte)
	var a1 int = 89
	b1 := 95
	fmt.Println("value of a is", a1, "and b is", b1)
	fmt.Printf("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1))

	//complex64: số phức có phần thực và phần ảo
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum", cadd)

	//chuyển đổi kiểu
	i := 55   //int
	j := 67.8 //float64
	sum := i + int(j)
	fmt.Println(sum)

}
