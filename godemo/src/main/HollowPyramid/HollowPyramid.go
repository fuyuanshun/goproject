package main

import "fmt"

/*
  - @author fys
    @date 2023/01/12
    @description 空心金字塔
*/
func main() {
	var row = 10
	for i := 1; i <= row; i++ {
		//每层添加空格
		for k := 1; k <= row-i; k++ {
			fmt.Print(" ")
		}
		//最后一行交替输出*和空格
		var b = true
		//打印*
		for j := 1; j <= 2*i-1; j++ {
			//最后一行特殊处理
			if i == row {
				if b {
					fmt.Print("*")
					b = false
				} else {
					fmt.Print(" ")
					b = true
				}
				//空心处理
			} else {
				if j == 1 || j == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
