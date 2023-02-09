package main

import (
	"fmt"
	"os"
)

// The variable os.Args là một lát cắt của chuỗi
// slice như là một chuỗi có kích thước động s của các phần tử mạng nơi các phần tử có thể truy cập riêng lẻ
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(len(os.Args))
}
