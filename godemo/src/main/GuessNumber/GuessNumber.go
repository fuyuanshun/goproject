package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  - @author fys
    @date 2023/01/13
    @description 猜数字小游戏
*/
func main() {
	//随机数种子
	rand.Seed(time.Now().Unix())
	//生成0~100的随机数
	var randomNum = rand.Intn(101)

	fmt.Println("随机生成了一个100以内（包含100）的数字，猜猜它是多少:")
	//统计猜数字的次数
	var count int8 = 0
	for {
		//统计次数
		count++

		var num int
		fmt.Scanln(&num)
		if num > randomNum {
			fmt.Println("比你想象中的小")
		} else if num < randomNum {
			fmt.Println("比你想象中的大")
		} else {
			fmt.Printf("恭喜你猜对了，共猜了%v次 \n", count)
			break
		}
	}
	if count < 3 {
		fmt.Printf("运气不错呦~")
	}
}
