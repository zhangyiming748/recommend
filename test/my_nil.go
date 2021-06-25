package main

import (
	"fmt"
)

func main() {
	var a string
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有错误发生:%v", err)
		}
	}()
	fmt.Scanln(&a)
	fmt.Printf("接收到的字符是:%s", a)
}
