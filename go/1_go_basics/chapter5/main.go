package main

import "fmt"

type Man struct {
	Name string
	LastName string
	Age int
	Gender string
	Crimes int
}

var people = make(map[string]Man)

func buildSuspects(people map[string]Man) []string {
	suspects := make([]string, len(people))
	for name, _ := range people {
		suspects = append(suspects, name)
	}
	return suspects
}

func main() {

	people["Igor"] = Man{"Igor", "Kukov", 19, "M", 1}
	people["Mahmud"] = Man{"Mahmud", "Abdusatorov", 20, "M", 666}
	people["Egor"] = Man{"Egor", "Krukov", 29, "M", 2}
	people["Roman"] = Man{"Roman", "Alyoshin", 37, "M", 1}
	people["Elena"] = Man{"Elena", "Stein", 41, "F", 0}

	suspects := buildSuspects(people)

	direCriminalName := suspects[0]
	greatestCrimesAmount := people[direCriminalName].Crimes

	for name, personData := range people {
		personCrimesAmount := personData.Crimes
		if personCrimesAmount > greatestCrimesAmount {
			direCriminalName = name
			greatestCrimesAmount = personCrimesAmount
		}
	}

	fmt.Println(people[direCriminalName])
}