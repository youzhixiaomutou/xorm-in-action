package main

import "fmt"

// 查询和统计数据
func query() {
	u := &xUser{}
	got, err := engine.ID(1).Get(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get=%t\n", got)
	if got {
		fmt.Printf("xUser=%+v\n", u)
	}
}
