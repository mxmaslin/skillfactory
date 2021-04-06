package electronic


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
	brand, model, phone_type, os string
	price int
}

type androidPhone struct {
	brand, model, phone_type, os string
	price int
}

type radioPhone struct {
	brand, model, phone_type string
	price, frequency, num_buttons int
}


type stationPhone struct {
	brand, model string
	price, num_buttons int
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
	return p.model
}

func (p androidPhone) Model() string {
	return p.model
}

func (p radioPhone) Model() string {
	return p.model
}

func (p applePhone) Type() string {
	return "smartphone"
}

func (p androidPhone) Type() string {
	return "smartphone"
}

func (p radioPhone) Type() string {
	return "station"
}

func (p radioPhone) ButtonsCount() int {
	return 12
}

func (s applePhone) OS() string {
	return "iOS"
}

func (s androidPhone) OS() string {
	return "android"
}



func NewApplePhone(model string, price int) *applePhone {
	p := new(applePhone)
	p.brand = "Apple"
	p.model = model
	p.price = price
	p.phone_type = p.Type()
	p.os = p.OS()
	return p
}

func NewAndroidPhone(brand, model string, price int) *androidPhone {
	p := new(androidPhone)
	p.brand = brand
	p.model = model
	p.price = price
	p.phone_type = p.Type()
	p.os = p.OS()
	return p
}

func NewRadioPhone(brand, model string, price, frequency int) *radioPhone {
	p := new(radioPhone)
	p.brand = brand
	p.model = model
	p.price = price
	p.frequency = frequency
	p.phone_type = p.Type()
	p.num_buttons = p.ButtonsCount()
	return p
}
