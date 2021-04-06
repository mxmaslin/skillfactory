package main

import (
	"fmt"
	electronic "github.com/mxmaslin/skillfactory/go/1_go_basics/chapter4/electronic"
)


func printCharacteristics(i electronic.Phone) {
	characteristics := fmt.Sprintf("Бренд: %v, модель: %v, тип: %v", i.Brand(), i.Model(), i.Type())
	switch i.(type){
	case electronic.Smartphone:
		characteristics += fmt.Sprintf(", OC: %v", i.(electronic.Smartphone).OS())
	case electronic.StationPhone:
		characteristics += fmt.Sprintf(", количество кнопок: %v", i.(electronic.StationPhone).ButtonsCount())
	}
	fmt.Println(characteristics)
}

func main() {
	theApplePhone := electronic.NewApplePhone("XXX")
	theAndroidPhone := electronic.NewAndroidPhone("Xiaomi", "13")
	theRadioPhone := electronic.NewRadioPhone("Sanyo", "666")

	printCharacteristics(theApplePhone)
	printCharacteristics(theAndroidPhone)
	printCharacteristics(theRadioPhone)

}
