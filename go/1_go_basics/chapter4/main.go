package main

import (
	"fmt"
	electronic "github.com/mxmaslin/skillfactory/go/1_go_basics/chapter4/electronic"
)


func printCharacteristics(i electronic.Phone) {
	info := fmt.Sprintf("Бренд: %v, модель: %v, тип: %v", i.Brand(), i.Model(), i.Type())
	addition := ". Моя ОС %v"
	if i.Type() == "smartphone" {
		if i.Brand() == "apple" {
			addition = fmt.Sprintf(addition, "iOS")
		} else {
			addition = fmt.Sprintf(addition, "android")
		}
	} else {
		addition = ". У меня 12 кнопок"
	}
	info = info + addition
	fmt.Println(info)
}

func main() {
	theApplePhone := electronic.NewApplePhone("X", 999)
	theAndroidPhone := electronic.NewAndroidPhone("Xiaomi", "13", 666)
	theRadioPhone := electronic.NewRadioPhone("Sanyo", "OldOne", 19, 127)

	printCharacteristics(theApplePhone)
	printCharacteristics(theAndroidPhone)
	printCharacteristics(theRadioPhone)

}
