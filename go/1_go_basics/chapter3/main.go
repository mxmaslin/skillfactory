package main

import (
	"fmt"
	calculate "github.com/mxmaslin/skillfactory/go/1_go_basics/chapter3/calc"
)

func main() {
	c := calculate.SetCalculator(1.0, 2.0, "+")
	fmt.Println("%v", c)
}