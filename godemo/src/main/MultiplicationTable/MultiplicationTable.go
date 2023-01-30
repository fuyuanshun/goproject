package main

import (
	"fmt"
)

/*
  - @author fys
    @date 2023/01/30
    @description 乘法表
*/
func main() {
	var i = 9
	for j := 1; j <= i; j++ {
		for k := 1; k <= j; k++ {
			fmt.Printf("%d x %d = %d\t", k, j, j*k)
		}
		fmt.Println()
	}
}
