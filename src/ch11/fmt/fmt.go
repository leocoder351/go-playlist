package main

import (
	"fmt"
)

func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "Hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成了一个 rune 切片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红" // string
	c2 := '红' // rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"       // string
	c4 := 'H'       // int32
	c5 := byte('H') // byte(unit8)
	fmt.Printf("c3:%T c4:%T c5:%T\n", c3, c4, c5)
	fmt.Printf("%d\n", c5)
}
