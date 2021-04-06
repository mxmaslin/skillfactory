package electronic


const (
	phoneTypeStationary = "station"
	phoneTypeSmartphone = "smartphone"
)



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



type phone struct {
	brand string
	model string
	phoneType string
}

type applePhone struct {
	phone
}

type androidPhone struct {
	phone
}

type radioPhone struct {
	phone
}

type stationPhone struct {
	phone
}



func (p phone) Brand() string {
	return p.brand
}

func (p phone) Model() string {
	return p.model
}

func (p phone) Type() string {
	return p.phoneType
}



func (p radioPhone) ButtonsCount() int {
	return 12
}

func (p applePhone) OS() string {
	return "iOS"
}

func (p androidPhone) OS() string {
	return "android"
}


func NewApplePhone(model string) *applePhone {
	p := new(applePhone)
	p.brand = "Apple"
	p.model = model
	p.phoneType = phoneTypeSmartphone
	return p
}

func NewAndroidPhone(brand, model string) *androidPhone {
	p := new(androidPhone)
	p.brand = brand
	p.model = model
	p.phoneType = phoneTypeSmartphone
	return p
}

func NewRadioPhone(brand, model string) *radioPhone {
	p := new(radioPhone)
	p.brand = brand
	p.model = model
	p.phoneType = phoneTypeStationary
	return p
}
