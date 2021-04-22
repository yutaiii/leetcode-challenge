package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2, 3}
	i2 := i
	i3 := i2

	i2[2] = 100
	i3[2] = 200
	fmt.Println(i)  // [1 2 200]
	fmt.Println(i2) // [1 2 200]
	fmt.Println(i3) // [1 2 200]
	// 全て参照渡しのような挙動になっている
}
