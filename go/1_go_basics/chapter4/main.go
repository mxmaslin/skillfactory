package main

import (
	"fmt"
	electronic "github.com/mxmaslin/skillfactory/go/1_go_basics/chapter4/electronic"
)

func main() {
	iphoneX := electronic.NewApplePhone("X", 999)
	xiaomi13 := electronic.NewAndroidPhone("Xiaomi", "13", 666)

	fmt.Println(iphoneX)
	fmt.Println(xiaomi13)
}
