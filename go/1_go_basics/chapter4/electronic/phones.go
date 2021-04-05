package electronic

import "fmt"

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}




type applePhone struct {
	model string
	price int
}

type androidPhone struct {
	brand, model string
	price int
}

type radioPhone struct {
	brand, model string
	price, frequency int
}


type stationPhone struct {
	brand, model string
	price int
}


func (p applePhone) Brand() string {
	return "apple"
}

func (p androidPhone) Brand() string {
	return p.brand
}

func (p radioPhone) Brand() string {
	return p.brand
}

func (p applePhone) Model() string {
	return fmt.Sprintf("I am iPhone model %v", p.Model)
}

func (p androidPhone) Model() string {
	return fmt.Sprintf("I am android phone  %v model %v", p.Brand, p.Model)
}

func (p radioPhone) Model() string {
	return fmt.Sprintf("I am radio phone %v model %v", p.Brand, p.Model)
}

func Type(i interface{}) string {
	switch i.(type) {
		case applePhone, androidPhone:
			return "smartphone"
		case radioPhone:
			return "station"
		default:
			return "?"
	}
}

func (p radioPhone) ButtonsCount() int {
	return 12
}

func (sp stationPhone) ButtonsCount() int {
	return 12
}

func OS(i interface{}) string {
	switch i.(type) {
		case applePhone:
			return "iOS"
		case androidPhone:
			return "android"
		default:
			return "?"
	}
}



func NewApplePhone(model string, price int) *applePhone {
	p := new(applePhone)
	p.model = model
	p.price = price
	return p
}

func NewAndroidPhone(brand, model string, price int) *androidPhone {
	p := new(androidPhone)
	p.brand = brand
	p.model = model
	p.price = price
	return p
}


func NewRadioPhone(brand, model string, price, frequency int) *radioPhone {
	p := new(radioPhone)
	p.brand = brand
	p.model = model
	p.price = price
	p.frequency = frequency
	return p
}
