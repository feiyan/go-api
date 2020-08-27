package main

import (
	"fmt"
	"gogin/app/route"
)

func main() {
	// 初始化引擎
	r := route.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}