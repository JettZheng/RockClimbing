package golang

import (
	"fmt"
)

func SliceCode() {
	var a = []int{1, 2, 3, 4, 5}
	var c []int = make([]int, 0)
	b := a[0:2]
	copy(c, b)
	b[0] = 9

	fmt.Println(c[0])
}
