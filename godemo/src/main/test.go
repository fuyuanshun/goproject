package main

import (
	"fmt"
)

func main() {
	i := 7
	var j int = 10
	fmt.Println("hello, world ", i)
	fmt.Println(j)
	fmt.Println(findAllData())
	var a int = 10000
	fmt.Printf("hello %T", a)

}

func findAllData() (int, int) {
	var (
		i = 1
		j = 2
	)
	return i, j
}
