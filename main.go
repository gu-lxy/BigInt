package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int64
	a = 1
	var b int64
	b = 3
	rs := a + b
	fmt.Println(rs)
	if a > b {
		fmt.Println("a > b")
	}else if a < b {
		fmt.Println("a < b")
	}else if a == b {
		fmt.Println("a = b")
	}
	bigA := big.NewInt(1)//定义一个大整数
	bigB := big.NewInt(3)
	//fmt.Println(bigA+bigB) //指针不能相加减
	bigC := bigA.Add(bigA,bigB)
	fmt.Println("大整数相加",bigC)//大整数相加

	//比较：compare
	//三个返回值：
	       // 小于 -1
	       // 大于 +1
	       // 等于 0
	r := bigA.Cmp(bigB)
	fmt.Println("大整数比较",r)

	bigN := big.NewInt(1)
	fmt.Println("大整数：",bigN)
	//移位运算
	rs2 := bigN.Lsh(bigN, 1)
	fmt.Println("移1位后的结果：",rs2)
	rs3 := bigN.Lsh(bigN, 254)
	fmt.Println("移254位后的结果：",rs3)
}
